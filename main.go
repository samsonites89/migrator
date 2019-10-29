package main

import (
	"fmt"
	"samsam.son/migrator/config"
	"samsam.son/migrator/database"
	myHttp "samsam.son/migrator/http"
	"samsam.son/migrator/njson"
)

func main() {

	db , _ := database.Init()
	// Transction Collector
	var transactions []njson.TxResult
	transactions, _ = myHttp.GetTxList(int(config.StartingBlock), 30000)

	for i := 0; i < len(transactions); i++ {
		fmt.Println(transactions[i].Input)
	}

	defer db.Close()
	fmt.Println(len(transactions))
}
