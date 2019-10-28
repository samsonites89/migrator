package njson

import "encoding/json"

type Response struct {
	Status  string          `json:"status"`
	Message string          `json:"message"`
	Result  json.RawMessage `json:"result"`
}

type TxResult struct {
	BlockNumber string `json:"blockNumber"`
	TxHash      string `json:"hash"`
	Nonce       string `json:"nonce"`
	From        string `json:"from"`
	To          string `json:"to"`
	Gas         string `json:"gas"`
	GasPrice    string `json:"gasPrice"`
	Input       string `json:"input"`
}
