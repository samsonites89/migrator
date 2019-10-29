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
	GethEndPoint	  = "https://rinkeby.infura.io/v3/382ffa6c91e9432e8f4bdaa088bf73c4"

	CustomerJourneyAddress = common.BytesToAddress([]byte("0xe320e5dce85fc135d1f1d5dadd0c5bc411c0ab5b"))

	// EthereumFromDefault holds the default sync block to sync from
	StartingBlock = uint64(5000735)



	DBFile = "migrator.db"

	// DEFileDefault holds the default name for the replay agent database
	DefaultKey = Key{
		File:     "config/rinkeby.key",
		Password: "rinkeby",
	}
)

type Key struct {
	File string
	Password string
}

func home() string {
	h, err := os.UserHomeDir()
	if err != nil {
		log.WithError(err).Fatal("could not determine user home directory")
	}
	return h
}
