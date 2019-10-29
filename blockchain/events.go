package blockchain

import (
	"context"
	"database/sql"
	"log"
	"math/big"
	"os"
	"samsam.son/migrator/config"
	"sync/atomic"
	"time"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

const (
	BookingConfirmationNumber = 2
	IoTDocType                = 4
	GeoFencingAppID           = 4
	GeoFencingDocType         = 6
	DeliveryOrderDocType      = 9
)

var (
	nonce uint64
)

// NextNonce returns the next nonce for the Oracles account.
func NextNonce() uint64 {
	return atomic.AddUint64(&nonce, 1) - 1
}

func setupNonce(ctx context.Context, me common.Address, c *ethclient.Client) error {
	var err error
	nonce, err = c.NonceAt(ctx, me, nil)
	if err == nil {
		log.Printf("found nonce %d for address 0x%x", nonce, me)
	}
	return err
}

// HandleEvents listens for events in the customer journey contract
// and processes them if necessary.
func HandleEvents(ctx context.Context , db *sql.DB) error {
	setup(db)

	// load key
	r, err := os.Open(config.DefaultKey.File)
	if err != nil {
		return err
	}
	auth, err := bind.NewTransactor(r, config.DefaultKey.Password)
	if err != nil {
		r.Close()
		return err
	}
	defer r.Close()

	// connect to Ethereum node
	client, err := ethclient.Dial(config.GethEndPoint)
	if err != nil {
		return nil
	}

	if err := setupNonce(ctx, auth.From, client); err != nil {
		return err
	}

	// instantiate customer journey smart contract filterer
	f, err := NewCustomerJourneyFilterer(cfg.Blockchain.CustomerJourney.Address, client)
	if err != nil {
		client.Close()
		return err
	}

	sm, err := NewCustomerJourney(config.CustomerJourneyAddress, client)
	if err != nil {
		client.Close()
		return err
	}

	// handle events
	go func() {
		defer client.Close()
		log.Printf("poll for AssetLinked events on the customer journey smart contract...")

		for {
			select {
			case <-ctx.Done():
				return
			case <-time.NewTimer(5 * time.Second).C:
				start, err := lastProcessedAssetLinkedEventBlock(db, cfg, client)
				if err != nil {
					log.Printf("could not determine last processed event block (err=%s)", err)
					continue
				}
				start++ // start was the last processed block, next unprocessed block is start+1

				end, err := currentBlock(ctx, client)
				if err != nil {
					log.Printf("could not determine blockchain head (err=%s)", err)
					continue
				}

				if start >= *end {
					continue // no new blocks
				}

				if *end-start > 500000000 {
					// the ethereum node sometimes reports a hugh value for its last
					// block without an error. If the difference between start and end
					// is too big retry it later because we have seen that the node
					// will report the correct block later.
					log.Printf("end -start block number too big (start=%d, end=%d)", start, *end)
					continue
				}

				opts := bind.FilterOpts{
					Start: start,
					End:   end,
				}

				linkedAssets, err := f.FilterAssetLinked(&opts, nil, nil)
				if err != nil {
					log.Printf("error fetch 'AssetLinked' events (start=%d,end=%d,err=%s)", start, end, err)
					continue
				}

				var (
					nIoT   = 0
					nDO    = 0
					failed = 0
				)

				for linkedAssets.Next() {
					if err := linkedAssets.Error(); err != nil {
						log.Printf("error parsing LinkedAsset event (err=%s)", err)
						continue
					}

					ev := linkedAssets.Event

					switch ev.Typ {
					case IoTDocType:
						bookingConfNr := "bkngConfNr" // TODO: Possibly ask another oracle with ID?
						data, err := getIoTAsset(cfg, notary, ev.Hash)
						if err != nil {
							log.Printf("could not retrieve IoT data (err=%s)", err)
							failed++
							continue
						}
						/*
								TODO: Currently All IoT Events from notary is stored as GeoFenceAssets if they are within one.
								If there are IoT events from the same device, it will record multiple geofence asset
								I think there should be only 2 geofence assets per predefined fence per device:
							    1) When it enters a geofence -> register geofence asset (entered)
							    2) When it leaves a geofence -> register geofence asset (left)
						*/
						a, hash, err := iot.Process(db, cfg, bookingConfNr, data, ev.Hash, ev.Parent, ev.Key)
						if err != nil {
							log.Printf("could not process IoT AssetLinked event (err=%s)", err)
							failed++
							continue
						}
						if a != nil {
							if err := storeGeofenceAsset(*auth, notary, sm, a, hash, ev.Parent); err != nil {
								log.Printf("could not register geo fence asset within notary or customer journey (err=%s)", err)
								failed++
								continue
							}
						}
						nIoT++

					// ? TODO:  This is not part of geofencing? I'll have to double-check
					case DeliveryOrderDocType:
						containers, err := getContainers(cfg, notary, ev.Hash)
						if err != nil {
							log.Printf("could not retrieve container information (err=%s)", err)
							failed++
							continue
						}
						if err := order.Process(db, containers, ev.Hash, ev.Parent, ev.Key); err != nil {
							log.Printf("could not process LinkedAsset delivery order event (err=%s)", err)
							failed++
							continue
						}
						nDO++
					}
				}

				if err := setProcessedAssetLinkedEventBlock(db, *end); err != nil {
					log.Printf("could not update processed asset link event block number (err=%s)", err)
				} else {
					log.Printf("processed %d AssetLinked events in range blocknr[%d,%d] (nIoT=%d,nDO=%d,failed=%d)",
						nIoT+nDO+failed, start, *end, nIoT, nDO, failed)
				}
			}
		}
	}()

	return nil
}

func storeGeofenceAsset(auth bind.TransactOpts, n *DeliverNotary, s *CustomerJourney, a *iot.GeoFenceAsset, hash, parent [32]byte) error {
	// store in notary
	var (
		network       uint8 = 2
	)

	go func() {
		// wait till transaction is included in blockchain
		failed := 0
		timeout := time.NewTimer(10 * time.Minute)
		for failed < 5 {
			select {
			case <-time.NewTimer(10 * time.Second).C:
				ass, err := n.Assets(nil, hash)
				if err != nil {
					failed++
					continue
				}
				failed = 0

				if ass.Network == network {
					auth.Nonce = new(big.Int).SetUint64(NextNonce())
					tx, err := s.Link(&auth, parent, hash, GeoFencingDocType, a.Container)
					if err != nil {
						log.Printf("could not link geofence asset to journey (err=%s)", err)
						timeout.Stop()
						return
					}
					log.Printf("linked geofence asset in customer journey (tx=0x%x)", tx.Hash())
					timeout.Stop()
					return
				}
			case <-timeout.C:
				log.Printf("could not link IoT data to customer journey, 10m timeout")
				return
			}
		}
	}()

	return nil
}

func setup(db *sql.DB) error {
	sql := `create table if not exists asset_linked_event(block number)`
	_, err := db.Exec(sql)
	return err
}

func currentBlock(ctx context.Context, c *ethclient.Client) (*uint64, error) {
	h, err := c.HeaderByNumber(ctx, nil)
	if err != nil {
		return nil, err
	}

	num := h.Number.Uint64()
	return &num, nil
}
