package blockchain

import (
	"context"
	"database/sql"
	"encoding/hex"
	"fmt"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/accounts/keystore"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	log "github.com/sirupsen/logrus"
	"io/ioutil"
	"math/big"
	"os"
	"samsam.son/migrator/config"
	"samsam.son/migrator/database"
	"samsam.son/migrator/njson"
	"sync/atomic"
	"time"
)

const (
	BookingConfirmationNumber = 2
	IoTDocType                = 4
	GeoFencingAppID           = 4
	GeoFencingDocType         = 6
	DeliveryOrderDocType      = 9
)

const (
	maxPendingTransactionInPool = uint(200)
)



var (
	nonce uint64
	signingKey	*keystore.Key
)

// NextNonce returns the next nonce for the Oracles account.
func NextNonce() uint64 {
	return atomic.AddUint64(&nonce, 1) - 1
}

// This has to be used as to not create collision
func setupNonce(ctx context.Context, me common.Address, c *ethclient.Client) error {
	var err error
	nonce, err = c.NonceAt(ctx, me, nil)
	if err == nil {
		log.Printf("found nonce %d for address 0x%x", nonce, me)
	}
	return err
}


func SendTransactionToGeth(ctx context.Context , db *sql.DB) error {

	// TODO maybe ill include this info in the db just to make sure, but if it's a repeat, it will probably not go in anyways,
	//setup(db)

	// 1. load key
	r, err := os.Open(config.DefaultKey.File)
	if err != nil {
		return err
	}

	// TODO I don't Think I have to do this every single time..
	// 2. connect to Ethereum node
	client, err := ethclient.Dial(config.GethEndPoint)
	if err != nil {
		return err
	}

	// 3. Create Nonce Value for the transaction
	if err := setupNonce(ctx, config.MyAddress, client); err != nil {
		return err
	}

	json, err := ioutil.ReadAll(r)
	if err != nil {
		return err
	}

	signingKey, err = keystore.DecryptKey(json, config.DefaultKey.Password)
	if err != nil {
		return err
	}
	fmt.Println()

	return nil
}

func RecreateSignedTransaction(txID string) (*types.Transaction, error) {

	//gas and data are input values.
	var dbtx njson.TxResult

	db ,err := database.Open()
	if err != nil {
		fmt.Println(err)
	}
	dbtx , err = database.GetCustomerJourneyTransaction(db,txID)
	if err !=nil {
		fmt.Println(err)
	}

	// Cut Data off	byteData, err := hex.DecodeString(textData[2:10])


	data , err := hex.DecodeString(dbtx.Input[2:])

	// What is this happens at the same time? with the other indexing?
	tx := types.NewTransaction(NextNonce(), config.CustomerJourneyAddress,
		big.NewInt(0), dbtx.Gas, new(big.Int).SetUint64(dbtx.GasPrice), data )

	rinkebyChainID := new(big.Int).SetUint64(uint64(4))

	signer := types.NewEIP155Signer(rinkebyChainID)
	signedTx, _ := types.SignTx(tx, signer, signingKey.PrivateKey)
	fmt.Println(signedTx)

	return signedTx , nil
}

func ResendTransactions(ctx context.Context, client *ethclient.Client, tx *types.Transaction) error {
	ctx, cancel := context.WithTimeout(ctx, 10*time.Second)
	err := client.SendTransaction(ctx, tx)
	cancel()
	if err != nil {
		return err
	}

	return nil
}

func createCustomerJourneyStart(auth bind.TransactOpts, s *CustomerJourney, hash [32]byte, parent [32]byte) error {

	auth.Nonce = new(big.Int).SetUint64(NextNonce())
	tx, err := s.Start(&auth, parent, GeoFencingDocType, "Container")
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

func currentBalance(ctx context.Context, c *ethclient.Client, blockHeight uint64) (*uint64, error) {

	height := new(big.Int).SetUint64(blockHeight)

	balance, err := c.BalanceAt(ctx, config.MyAddress,height)
	if err != nil {
		return nil, err
	}
	fmt.Println(balance)
	num := balance.Uint64()

	return &num, nil
}

func targetTransactionPoolHasRoom(ctx context.Context, client *ethclient.Client) (uint, error) {
	ctx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()

	count, err := client.PendingTransactionCount(ctx)
	if err != nil {
		return 0, err
	}
	if count <= maxPendingTransactionInPool {
		return maxPendingTransactionInPool - count, nil
	}
	return 0, nil
}
