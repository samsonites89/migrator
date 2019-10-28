package database

import (
	"fmt"
	"samsam.son/migrator/config"
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

func TestInsert(t *testing.T) {
	config.DBFile = "test.db"

	db, err := Open()
	if err != nil {
		fmt.Printf("Error thrown while Opening Database : %v", err)
	}

	defer db.Close()

	// If there is no error, it can be assumed that the data is ok

}
