package database

import (
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
	log "github.com/sirupsen/logrus"
	"samsam.son/migrator/abi"
	"samsam.son/migrator/config"
	"samsam.son/migrator/njson"
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
	q := `create table if not exists collected_transaction (block number, txid text, sender text, toAddr text, nonce number, gas number, gasPrice number, data text, funcType number, primary key (block, txid))`
	if _, err := db.Exec(q); err != nil {
		log.WithError(err).Fatal(`could not create "collected_CSStart" table`)
	}

}

func InsertTransaction(db *sql.DB, tx njson.TxResult) {

	//Parse Input to see what it is: Link, Start?
	funcType , _ := abi.DetermineType(tx.Input)
	log.Debugf("SC Function Type : &d" , funcType)
	_, err := db.Exec(`insert into collected_CSStart(block number, txidx number, sender text, toAddr text, nonce number, gas number, gasPrice number, data text,funcType number) values (?,?,?,?,?,?,?,?,?)`,
		tx.BlockNumber, tx.TxHash, tx.From, tx.To, tx.Nonce, tx.Gas, tx.GasPrice, tx.Input, funcType)
	if err != nil {
		log.WithError(err).Fatal(`could not insert into "CSLink" table`)
	}
}
