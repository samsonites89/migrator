package config

import (
	log "github.com/sirupsen/logrus"
	"os"
)

var (
	// Endpoint for etherscan
	EtherscanEndPoint = "http://api-rinkeby.etherscan.io"
	ApiKey            = "MDPT8U7JZ2DB54PXV59TPVQQZ94QJSB5PN"
	Address           = "0x53D5450f23d3FdA15b8C1813458f1e51a783d88a"

	// EthereumFromDefault holds the default sync block to sync from
	StartingBlock = uint64(5000735)

	DBFile = "migrator.db"

	// DEFileDefault holds the default name for the replay agent database

)

func home() string {
	h, err := os.UserHomeDir()
	if err != nil {
		log.WithError(err).Fatal("could not determine user home directory")
	}
	return h
}
