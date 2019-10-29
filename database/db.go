package database

import (
	"context"
	"database/sql"
	"fmt"
	_ "github.com/mattn/go-sqlite3"
	log "github.com/sirupsen/logrus"
	"samsam.son/migrator/abi"
	"samsam.son/migrator/config"
	"samsam.son/migrator/njson"
)

var (
	ctx, stop = context.WithCancel(context.Background())
)

// Initialize Database
func Init() (*sql.DB, error) {
	db, err := sql.Open("sqlite3", config.DBFile)
	if err != nil {
		return nil, err
	}
	createTables(db)

	return db, nil

}

// Open returns a database connection ready to use or an
// error in case the connection could not be opened.
func Open() (*sql.DB, error) {
	return sql.Open("sqlite3", config.DBFile)
}

// Create tables if table does not exist
func createTables(db *sql.DB) {

	//Create a Table for Transactions
	q := `create table if not exists collected_transaction (block number, txid text, sender text, toAddr text, nonce number, gas number, gasPrice number, data text, funcType number, primary key (txid))`
	if _, err := db.Exec(q); err != nil {
		log.WithError(err).Fatal(`could not create "collected_transaction" table`)
	}

}

func InsertTransaction(db *sql.DB, tx njson.TxResult) {

	//Parse Input to see what it is: Link, Start?
	funcType, _ := abi.DetermineType(tx.Input)
	log.Debugf("SC Function Type : %d", funcType)

	_, err := db.Exec(`insert into collected_transaction(block , txid , sender , toAddr , nonce , gas , gasPrice , data , funcType) values (?,?,?,?,?,?,?,?,?)`,
		tx.BlockNumber, tx.TxHash, tx.From, tx.To, tx.Nonce, tx.Gas, tx.GasPrice, tx.Input, funcType)
	if err != nil {
		log.WithError(err).Fatal(`could not insert into " collected_transaction" table`)
	} else {
		log.Debugf("%s inserted into database", tx.TxHash)
	}

}

// Fetch a Single Tx from the DB.
func GetCustomerJourneyTransaction(db *sql.DB, txid string) (njson.TxResult, error) {
	var (
		q  = `select * from collected_transaction where txid = ?`
		tx njson.TxResult

		block    uint64
		sender   string
		toAddr   string
		nonce    uint64
		gas      uint64
		gasPrice uint64
		data     string
		funcType int
		txidQ    string
	)

	rows := db.QueryRowContext(ctx, q, txid)
	err := rows.Scan(&block, &txidQ, &sender, &toAddr, &nonce, &gas, &gasPrice, &data, &funcType)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(rows)

	tx = njson.TxResult{
		BlockNumber: block,
		TxHash:      txidQ,
		To:          sender,
		From:        sender,
		Gas:         gas,
		GasPrice:    gasPrice,
		Input:       data,
		Nonce:       nonce,
	}

	print(tx.Input)

	return tx, nil
}
