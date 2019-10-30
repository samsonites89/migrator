package blockchain

import (
	"context"
	"fmt"
	"github.com/ethereum/go-ethereum/ethclient"
	"samsam.son/migrator/config"
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

	//SendTransactionToGeth(ctx,)


}