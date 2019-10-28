package main

import (
	"fmt"
	"samsam.son/migrator/config"
	myHttp "samsam.son/migrator/http"
	"samsam.son/migrator/njson"
)

func main() {

	// Transction Collector
	var transactions []njson.TxResult
	transactions, _ = myHttp.GetTxList(int(config.StartingBlock), 30000)

	for i := 0; i < len(transactions); i++ {
		fmt.Println(transactions[i].Input)
	}

	fmt.Println(len(transactions))
}
