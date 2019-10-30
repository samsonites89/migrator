package main

import (
	"flag"
	log "github.com/sirupsen/logrus"
	"samsam.son/migrator/config"
	"samsam.son/migrator/database"
	comm "samsam.son/migrator/http"
	"samsam.son/migrator/njson"
)

func main() {

	// Initialize Database
	db, _ := database.Init()

	// Take in simple arg value for the number of Txs
	start := flag.Int("start", int(config.StartingBlock), "starting block info")
	collectRange := flag.Int("range", 100, "range for query")
	flag.Parse()

	log.Infof("START: %v. The tx collect range is currently set to %v ", *start, *collectRange)

	var transactions []njson.TxResult
	transactions, _ = comm.GetTxList(*start, *collectRange)

	log.Debugf("The number of tx is %d", len(transactions))

	for i := 0; i < len(transactions); i++ {
		database.InsertTransaction(db, transactions[i])
	}

	defer db.Close()

}
