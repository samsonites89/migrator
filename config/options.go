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
	QueryAddress = "0x53D5450f23d3FdA15b8C1813458f1e51a783d88a"
	StartingBlock = uint64(5000735) // Default Starting Block

	// Endpoint for Infura for TX indexing
	GethEndPoint = "https://rinkeby.infura.io/v3/382ffa6c91e9432e8f4bdaa088bf73c4"

	// The targets.
	CustomerJourneyAddress = common.HexToAddress("0xe320e5dce85fc135d1f1d5dadd0c5bc411c0ab5b")
	MyAddress              = common.HexToAddress("0x58d08109e2b445ab96114fe8b749fdd79b43b291")


	// DB Name
	DBFile = "migrator.db"

	// ! Key from @bas
	//DefaultKey = Key{
	//	File:     "config/new_rinkeby.key",
	//	Password: "sam",
	//}
	DefaultKey = Key{
		File:     "config/cello_rinkeby.key",
		Password: "merong30",
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
