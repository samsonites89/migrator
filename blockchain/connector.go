package blockchain

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	log "github.com/sirupsen/logrus"
	"math/big"
	"os"
	"samsam.son/migrator/config"
	"sync/atomic"
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
func SendTransactionToGeth(ctx context.Context , db *sql.DB) error {

	// TODO maybe ill include this info in the db just to make sure, but if it's a repeat, it will probably not go in anyways,
	//setup(db)

	// 1. load key
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

	// 2. connect to Ethereum node
	client, err := ethclient.Dial(config.GethEndPoint)
	if err != nil {
		return nil
	}

	// 3. Create Nonce Value for the transaction
	if err := setupNonce(ctx, auth.From, client); err != nil {
		return err
	}

	// 4. Instantiate a New Customer Journey Addr.
	sm, err := NewCustomerJourney(config.CustomerJourneyAddress, client)
	if err != nil {
		client.Close()
		return err
	}
	fmt.Println(sm)
	// Just a simple logic to see to figure out the blockchain height.
	//end, err := currentBlock(ctx, client)
	//if err != nil {
	//	log.Printf("could not determine blockchain head (err=%s)", err)
	//}
	//log.Debugf("The current Block is at : %v" , end )


	return nil
}

func createCustomerJourneyLink(auth bind.TransactOpts, s *CustomerJourney, hash [32]byte, parent [32]byte) error {

	auth.Nonce = new(big.Int).SetUint64(NextNonce())
	tx, err := s.Link(&auth, parent, hash, GeoFencingDocType, "Container")
	if err != nil {
		log.Printf("could not link geofence asset to journey (err=%s)", err)
	}
	log.Printf("linked geofence asset in customer journey (tx=0x%x)", tx.Hash())


	return nil
}



//func setup(db *sql.DB) error {
//	sql := `create table if not exists asset_linked_event(block number)`
//	_, err := db.Exec(sql)
//	return err
//}

// Get the current block information
func currentBlock(ctx context.Context, c *ethclient.Client) (*uint64, error) {
	h, err := c.HeaderByNumber(ctx, nil)
	if err != nil {
		return nil, err
	}

	num := h.Number.Uint64()
	return &num, nil
}
