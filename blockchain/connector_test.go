package blockchain

import (
	"context"
	"fmt"
	"github.com/ethereum/go-ethereum/ethclient"
	"samsam.son/migrator/config"
	"samsam.son/migrator/database"
	"samsam.son/migrator/njson"
	"testing"
)


var (
	ctx, stop  = context.WithCancel(context.Background())
)

func TestCurrentBlock(t *testing.T) {

	// 2. connect to Ethereum node
	client, err := ethclient.Dial(config.GethEndPoint)
	if err != nil {
		fmt.Printf("Error: %v \n" , err)
	}

	end, err := currentBlock(ctx, client)
	if err != nil {
		fmt.Printf("could not determine blockchain head (err=%s)", err)
	}
	fmt.Printf("The current Block is at : %v" , *end )

}

func TestGetBalance(t *testing.T) {


	height := uint64(5350447)

	client, err := ethclient.Dial(config.GethEndPoint)
	if err != nil {
		fmt.Printf("Error: %v \n" , err)
	}

	balance, err := currentBalance(ctx, client, height)
	if err != nil {
		fmt.Printf("could not determine blockchain head (err=%s)", err)
	}
	fmt.Println(*balance)

}

func TestSendTransactionToGeth(t *testing.T) {

	var tx njson.TxResult
	config.DefaultKey.File = "../config/new_rinkeby.key"
	config.DBFile = "../migrator.db"
	db ,err := database.Open()
	if err != nil {
		fmt.Println(err)
	}
	tx , err = database.GetCustomerJourneyTransaction(db,"0x1973cfaa91cc72ef4e6ae4f63d6c3453b08ff3146cd390381d2d9298fe415aa1")
	if err !=nil {
		fmt.Println(err)
	}
	fmt.Println(tx.To)
	defer db.Close()

	err = SendTransactionToGeth(ctx,db)
	if err != nil {
		fmt.Println(err)
	}

	signedTx , err := RecreateSignedTransaction("0x1973cfaa91cc72ef4e6ae4f63d6c3453b08ff3146cd390381d2d9298fe415aa1")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(signedTx)
	client, err := ethclient.Dial(config.GethEndPoint)
	if err != nil {
		fmt.Printf("Error: %v \n" , err)
	}

	err = ResendTransactions(ctx,client,signedTx)
	if err !=nil {
		fmt.Println(err)
	}
}