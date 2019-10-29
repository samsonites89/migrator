package database

import (
	"fmt"
	"io/ioutil"
	"samsam.son/migrator/abi"
	"samsam.son/migrator/config"
	"samsam.son/migrator/njson"
	"testing"
)

func TestOpen(t *testing.T) {
	config.DBFile = "test.db"

	db, err := Init()
	if err != nil {
		fmt.Println(err)
	}

	// register in DB
	defer db.Close()
}

func TestInsertTransaction(t *testing.T) {

	abi.InitAbI()
	//Primary Setup
	config.DBFile = "test.db"

	db, err := Init()
	if err != nil {
		fmt.Printf("Error thrown while Opening Database : %v", err)
	}

	// If there is no error, it can be assumed that the data is ok
	// Read from file
	jsonData, err := ioutil.ReadFile("test_data.json")
	if err != nil {
		fmt.Printf("error reading test json file : %v", err)
		t.Fail()
	}
	tx, err := njson.ParseJSONResponse(jsonData)

	defer db.Close()
	i := 0
	for i < len(tx) {
		InsertTransaction(db, tx[i])
		i++
		fmt.Printf("iteration : %d", i)
	}
}

func TestGetCustomerJourneyTransaction(t *testing.T) {
	abi.InitAbI()

	//Primary Setup
	config.DBFile = "test.db"

	db, err := Init()
	if err != nil {
		fmt.Printf("Error thrown while Opening Database : %v", err)
	}

	var getTx njson.TxResult
	getTx, _ = GetCustomerJourneyTransaction(db, "0x1973cfaa91cc72ef4e6ae4f63d6c3453b08ff3146cd390381d2d9298fe415aa1")

	fmt.Println(getTx)

	defer db.Close()

}
