package http

import (
	"samsam.son/migrator/config"
	"samsam.son/migrator/njson"
	"testing"
)

func TestGetTxList(t *testing.T) {

	var txs []njson.TxResult
	GetTxList(&txs, int(config.StartingBlock), 1000)

	if len(txs) != 10 {
		t.Fail()
	}
}
