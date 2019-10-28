package http

import (
	"fmt"
	log "github.com/sirupsen/logrus"
	"io/ioutil"
	"net/http"
	"samsam.son/migrator/config"
	"samsam.son/migrator/njson"
	"time"
)

// TODO should throw error if there is something wrong like connection cannot be made etc.
func GetTxList(start int, size int) ([]njson.TxResult, error) {

	endpoint := config.EtherscanEndPoint
	startBlock := start
	endBlock := size + start
	GetAccountTxList := fmt.Sprintf("api?module=account&action=txlist&address=%s&startblock=%v&endblock=%v&sort=asc&apikey=%s",
		config.Address, startBlock, endBlock, config.ApiKey)

	fullGETUrl := fmt.Sprintf("%s/%s", endpoint, GetAccountTxList)

	var myClient = &http.Client{
		Transport: &http.Transport{
			Proxy:                 http.ProxyFromEnvironment,
			ResponseHeaderTimeout: time.Second * 30,
		},
	}

	response, err := myClient.Get(fullGETUrl)
	if err != nil {
		log.Errorf("Fetched Data is not in the correct format")
	}

	bodyBytes, err := ioutil.ReadAll(response.Body)
	txs, err := njson.ParseJSONResponse(bodyBytes)

	log.Debugf("Transactions from Response has been marshalled : transaction counts = %d", len(txs))

	if err != nil {
		log.Errorf("Fetched Data is not in the correct format")

	}

	defer response.Body.Close()

	return txs, nil
}
