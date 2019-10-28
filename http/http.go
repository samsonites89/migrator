package http

import (
	"encoding/json"
	_ "encoding/json"
	"fmt"
	log "github.com/sirupsen/logrus"
	"io/ioutil"
	"net/http"
	"samsam.son/migrator/config"
	"samsam.son/migrator/njson"
	"time"
)

// TODO should throw error if there is something wrong like connection cannot be made etc.
func GetTxList(transactions *[]njson.TxResult, start int, size int) {

	var resp njson.Response

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
	err = json.Unmarshal(bodyBytes, &resp)
	if err != nil {
		log.Errorf("Error Unmarshalling Response Body : %v", err)
	}

	if resp.Message == "OK" {
		err = json.Unmarshal(resp.Result, &transactions)
		if err != nil {
			log.Errorf("Error Unmarshalling Response Body : %v", err)
		}

	}

	defer response.Body.Close()

}
