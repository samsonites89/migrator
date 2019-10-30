package config

import (
	"github.com/ethereum/go-ethereum/common"
	log "github.com/sirupsen/logrus"
	"os"
)

var (
	// Endpoint for etherscan
	EtherscanEndPoint = "http://api-rinkeby.etherscan.io"
	ApiKey            = "MDPT8U7JZ2DB54PXV59TPVQQZ94QJSB5PN"
	Address           = "0x53D5450f23d3FdA15b8C1813458f1e51a783d88a"
	GethEndPoint      = "https://rinkeby.infura.io/v3/382ffa6c91e9432e8f4bdaa088bf73c4"

	CustomerJourneyAddress = common.HexToAddress("0xe320e5dce85fc135d1f1d5dadd0c5bc411c0ab5b")
	MyAddress              = common.HexToAddress("0x626ca86b2c9ef9c0b8c29e0a0af059da124b2c6c")
	// EthereumFromDefault holds the default sync block to sync from
	StartingBlock = uint64(5000735)

	DBFile = "migrator.db"

	// DEFileDefault holds the default name for the replay agent database
	DefaultKey = Key{
		File:     "config/new_rinkeby.key",
		Password: "sam",
	}
)

type Key struct {
	File     string
	Password string
}

func home() string {
	h, err := os.UserHomeDir()
	if err != nil {
		log.WithError(err).Fatal("could not determine user home directory")
	}
	return h
}
