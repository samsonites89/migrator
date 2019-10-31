package njson

import (
	"encoding/json"
	log "github.com/sirupsen/logrus"
)

// Basic JSON structs for the migrator

type Response struct {
	Status  int             `json:"status,string"`
	Message string          `json:"message"`
	Result  json.RawMessage `json:"result"`
}

type TxResult struct {
	BlockNumber uint64 `json:"blockNumber,string"`
	TxHash      string `json:"hash"`
	Nonce       uint64 `json:"nonce,string"`
	From        string `json:"from"`
	To          string `json:"to"`
	Gas         uint64 `json:"gas,string"`
	GasPrice    uint64 `json:"gasPrice,string"`
	Input       string `json:"input"`
}

//TODO : create structs for CustomerJourney_Start , _Link

// Parse JSON Response from Etherscan into tx structures.
func ParseJSONResponse(responseBody []byte) ([]TxResult, error) {

	var resp Response
	var transactions []TxResult
	err := json.Unmarshal(responseBody, &resp)
	if err != nil {
		log.Errorf("Error Unmarshalling Response Body : %v", err)
		return nil, err
	}

	if resp.Status == 1 {
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
