package http

import (
	"fmt"
	"samsam.son/migrator/config"
	"samsam.son/migrator/njson"
	"testing"
)

func TestGetTxList(t *testing.T) {

	var txs []njson.TxResult
	txs, err := GetTxList(int(config.StartingBlock), 1000)
	if err != nil {
		t.Fail()
	}
	fmt.Println(len(txs))
	if len(txs) != 10 {
		t.Fail()
	}
}
