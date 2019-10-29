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
	fmt.Println(client)

	//end, err := currentBlock(ctx, client)
	//if err != nil {
	//	log.Printf("could not determine blockchain head (err=%s)", err)
	//}
	//log.Debugf("The current Block is at : %v" , end )

}

func TestSendTransactionToGeth(t *testing.T) {

	//SendTransactionToGeth(ctx,)


}