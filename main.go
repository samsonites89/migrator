package main

import (
	log "github.com/sirupsen/logrus"
	"os"
	"samsam.son/migrator/config"
	"samsam.son/migrator/database"
	comm "samsam.son/migrator/http"
	"samsam.son/migrator/njson"
)

func main() {

	// Initialize Database
	db , _ := database.Init()

	// Take in simple arg value for the number of Txs

	var collectRange int
	start := int(config.StartingBlock)

	args := os.Args[1]
	collectRange = int(args[1])
	log.Debugf("The tx collect range is currently set to 100")
	// Transction Collector
	var transactions []njson.TxResult
	transactions, _ = comm.GetTxList(start, collectRange)

	log.Debugf("The number of tx is %d" , len(transactions))


	defer db.Close()
	for i := 0; i < len(transactions); i++ {
		database.InsertTransaction(db,transactions[i])
	}


}
