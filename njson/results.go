package njson

import (
	"encoding/json"
	log "github.com/sirupsen/logrus"
)

// Basic JSON structs for the migrator

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

// Parse JSON Response from Etherscan
func ParseJSONResponse(responseBody []byte) ([]TxResult, error) {

	var resp Response
	var transactions []TxResult
	err := json.Unmarshal(responseBody, &resp)
	if err != nil {
		log.Errorf("Error Unmarshalling Response Body : %v", err)
		return nil, err
	}

	if resp.Status == "1" {
		if err := json.Unmarshal(resp.Result, &transactions); err != nil {
			log.Errorf("Error Unmarshalling Response Body : %v", err)
			return nil, err
		}

	} else { // resp.Status is probably 0 and there is an error,
		log.Errorf("Error Response from Request : %s", resp.Message)
		return nil, err
	}

	return transactions, nil
}
