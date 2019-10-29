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

	abi.Init()
	//Primary Setup
	config.DBFile = "test.db"

	db, err := Init()
	if err != nil {
		fmt.Printf("Error thrown while Opening Database : %v", err)
	}

	defer db.Close()

	// If there is no error, it can be assumed that the data is ok
	// Read from file
	jsonData, err := ioutil.ReadFile("test_data.json")
	if err != nil {
		fmt.Printf("error reading test json file : %v", err)
		t.Fail()
	}
	tx, err := njson.ParseJSONResponse(jsonData)

	fmt.Println(len(tx))
	InsertTransaction(db, tx[0])

}
