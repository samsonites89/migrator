package abi

import (
	"encoding/hex"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	log "github.com/sirupsen/logrus"
	"strings"
)

var (
	notaryABI    abi.ABI
	notaryABIStr = `[
	{
		"constant": false,
		"inputs": [
			{
				"internalType": "bytes32",
				"name": "hash",
				"type": "bytes32"
			},
			{
				"internalType": "enum Notary.Network",
				"name": "myNetwork",
				"type": "uint8"
			}
		],
		"name": "rejectOwnership",
		"outputs": [],
		"payable": false,
		"stateMutability": "nonpayable",
		"type": "function"
	},
	{
		"constant": false,
		"inputs": [
			{
				"internalType": "bytes32",
				"name": "hash",
				"type": "bytes32"
			},
			{
				"internalType": "enum Notary.Network",
				"name": "myNetwork",
				"type": "uint8"
			},
			{
				"internalType": "enum Notary.Network",
				"name": "newNetwork",
				"type": "uint8"
			},
			{
				"internalType": "address",
				"name": "newOwner",
				"type": "address"
			}
		],
		"name": "transferOwnership",
		"outputs": [],
		"payable": false,
		"stateMutability": "nonpayable",
		"type": "function"
	},
	{
		"constant": false,
		"inputs": [
			{
				"internalType": "bytes32",
				"name": "hash",
				"type": "bytes32"
			},
			{
				"internalType": "enum Notary.Network",
				"name": "myNetwork",
				"type": "uint8"
			}
		],
		"name": "acceptClaim",
		"outputs": [],
		"payable": false,
		"stateMutability": "nonpayable",
		"type": "function"
	},
	{
		"constant": false,
		"inputs": [
			{
				"internalType": "bytes32[]",
				"name": "hashes",
				"type": "bytes32[]"
			}
		],
		"name": "cancelTransferOwnership",
		"outputs": [],
		"payable": false,
		"stateMutability": "nonpayable",
		"type": "function"
	},
	{
		"constant": true,
		"inputs": [
			{
				"internalType": "bytes32",
				"name": "",
				"type": "bytes32"
			}
		],
		"name": "pendingTransfers",
		"outputs": [
			{
				"internalType": "enum Notary.Network",
				"name": "network",
				"type": "uint8"
			},
			{
				"internalType": "address",
				"name": "owner",
				"type": "address"
			},
			{
				"internalType": "enum Notary.ApplicationId",
				"name": "app",
				"type": "uint8"
			},
			{
				"internalType": "string",
				"name": "pathSpec",
				"type": "string"
			},
			{
				"internalType": "uint32",
				"name": "expiresAt",
				"type": "uint32"
			}
		],
		"payable": false,
		"stateMutability": "view",
		"type": "function"
	},
	{
		"constant": false,
		"inputs": [
			{
				"internalType": "enum Notary.Network",
				"name": "myNetwork",
				"type": "uint8"
			},
			{
				"internalType": "bytes32",
				"name": "hash",
				"type": "bytes32"
			},
			{
				"internalType": "enum Notary.ApplicationId",
				"name": "app",
				"type": "uint8"
			},
			{
				"internalType": "string",
				"name": "pathSpec",
				"type": "string"
			},
			{
				"internalType": "uint32",
				"name": "expiresAt",
				"type": "uint32"
			}
		],
		"name": "registerAsset",
		"outputs": [],
		"payable": false,
		"stateMutability": "nonpayable",
		"type": "function"
	},
	{
		"constant": false,
		"inputs": [
			{
				"internalType": "bytes32",
				"name": "hash",
				"type": "bytes32"
			},
			{
				"internalType": "enum Notary.Network",
				"name": "myNetwork",
				"type": "uint8"
			},
			{
				"internalType": "enum Notary.ApplicationId",
				"name": "app",
				"type": "uint8"
			},
			{
				"internalType": "string",
				"name": "pathSpec",
				"type": "string"
			},
			{
				"internalType": "uint32",
				"name": "expiresAt",
				"type": "uint32"
			}
		],
		"name": "claimAsset",
		"outputs": [],
		"payable": false,
		"stateMutability": "nonpayable",
		"type": "function"
	},
	{
		"constant": true,
		"inputs": [
			{
				"internalType": "bytes32",
				"name": "",
				"type": "bytes32"
			}
		],
		"name": "assets",
		"outputs": [
			{
				"internalType": "enum Notary.Network",
				"name": "network",
				"type": "uint8"
			},
			{
				"internalType": "address",
				"name": "owner",
				"type": "address"
			},
			{
				"internalType": "enum Notary.ApplicationId",
				"name": "app",
				"type": "uint8"
			},
			{
				"internalType": "string",
				"name": "pathSpec",
				"type": "string"
			},
			{
				"internalType": "uint32",
				"name": "expiresAt",
				"type": "uint32"
			}
		],
		"payable": false,
		"stateMutability": "view",
		"type": "function"
	},
	{
		"constant": false,
		"inputs": [
			{
				"internalType": "bytes32",
				"name": "hash",
				"type": "bytes32"
			},
			{
				"internalType": "enum Notary.Network",
				"name": "myNetwork",
				"type": "uint8"
			},
			{
				"internalType": "enum Notary.ApplicationId",
				"name": "app",
				"type": "uint8"
			},
			{
				"internalType": "string",
				"name": "pathSpec",
				"type": "string"
			},
			{
				"internalType": "uint32",
				"name": "expiresAt",
				"type": "uint32"
			}
		],
		"name": "acceptOwnership",
		"outputs": [],
		"payable": false,
		"stateMutability": "nonpayable",
		"type": "function"
	},
	{
		"constant": false,
		"inputs": [
			{
				"internalType": "enum Notary.Network",
				"name": "myNetwork",
				"type": "uint8"
			},
			{
				"internalType": "bytes32",
				"name": "currentHash",
				"type": "bytes32"
			},
			{
				"internalType": "bytes32",
				"name": "newHash",
				"type": "bytes32"
			},
			{
				"internalType": "enum Notary.ApplicationId",
				"name": "app",
				"type": "uint8"
			},
			{
				"internalType": "string",
				"name": "pathSpec",
				"type": "string"
			},
			{
				"internalType": "uint32",
				"name": "expiresAt",
				"type": "uint32"
			}
		],
		"name": "updateAsset",
		"outputs": [],
		"payable": false,
		"stateMutability": "nonpayable",
		"type": "function"
	},
	{
		"constant": false,
		"inputs": [
			{
				"internalType": "bytes32",
				"name": "hash",
				"type": "bytes32"
			},
			{
				"internalType": "enum Notary.Network",
				"name": "myNetwork",
				"type": "uint8"
			}
		],
		"name": "rejectClaim",
		"outputs": [],
		"payable": false,
		"stateMutability": "nonpayable",
		"type": "function"
	},
	{
		"constant": false,
		"inputs": [
			{
				"internalType": "bytes32",
				"name": "hash",
				"type": "bytes32"
			},
			{
				"internalType": "enum Notary.Network",
				"name": "myNetwork",
				"type": "uint8"
			}
		],
		"name": "deleteAsset",
		"outputs": [],
		"payable": false,
		"stateMutability": "nonpayable",
		"type": "function"
	},
	{
		"constant": true,
		"inputs": [
			{
				"internalType": "bytes32",
				"name": "",
				"type": "bytes32"
			}
		],
		"name": "pendingTransfersTimestamp",
		"outputs": [
			{
				"internalType": "uint256",
				"name": "",
				"type": "uint256"
			}
		],
		"payable": false,
		"stateMutability": "view",
		"type": "function"
	},
	{
		"anonymous": false,
		"inputs": [
			{
				"indexed": true,
				"internalType": "bytes32",
				"name": "hash",
				"type": "bytes32"
			},
			{
				"indexed": false,
				"internalType": "enum Notary.Network",
				"name": "network",
				"type": "uint8"
			},
			{
				"indexed": false,
				"internalType": "address",
				"name": "owner",
				"type": "address"
			}
		],
		"name": "AssetRegistered",
		"type": "event"
	},
	{
		"anonymous": false,
		"inputs": [
			{
				"indexed": true,
				"internalType": "bytes32",
				"name": "newHash",
				"type": "bytes32"
			},
			{
				"indexed": true,
				"internalType": "bytes32",
				"name": "previousHash",
				"type": "bytes32"
			},
			{
				"indexed": false,
				"internalType": "enum Notary.Network",
				"name": "network",
				"type": "uint8"
			}
		],
		"name": "AssetUpdated",
		"type": "event"
	},
	{
		"anonymous": false,
		"inputs": [
			{
				"indexed": false,
				"internalType": "bytes32",
				"name": "hash",
				"type": "bytes32"
			},
			{
				"indexed": true,
				"internalType": "address",
				"name": "addr",
				"type": "address"
			},
			{
				"indexed": true,
				"internalType": "enum Notary.Network",
				"name": "network",
				"type": "uint8"
			}
		],
		"name": "AssetNewClaim",
		"type": "event"
	},
	{
		"anonymous": false,
		"inputs": [
			{
				"indexed": true,
				"internalType": "bytes32",
				"name": "hash",
				"type": "bytes32"
			}
		],
		"name": "AssetClaimRejected",
		"type": "event"
	},
	{
		"anonymous": false,
		"inputs": [
			{
				"indexed": true,
				"internalType": "bytes32",
				"name": "hash",
				"type": "bytes32"
			},
			{
				"indexed": false,
				"internalType": "enum Notary.Network",
				"name": "oldNetwork",
				"type": "uint8"
			},
			{
				"indexed": false,
				"internalType": "address",
				"name": "oldOwner",
				"type": "address"
			},
			{
				"indexed": false,
				"internalType": "enum Notary.Network",
				"name": "newNetwork",
				"type": "uint8"
			},
			{
				"indexed": false,
				"internalType": "address",
				"name": "newOwner",
				"type": "address"
			}
		],
		"name": "AssetOwnershipTransferred",
		"type": "event"
	},
	{
		"anonymous": false,
		"inputs": [
			{
				"indexed": true,
				"internalType": "bytes32",
				"name": "hash",
				"type": "bytes32"
			}
		],
		"name": "AssetDeleted",
		"type": "event"
	}
]`

	customerJourneyABI    abi.ABI
	customerJourneyABIStr = `[
    {
      "constant": false,
      "inputs": [],
      "name": "kill",
      "outputs": [],
      "payable": false,
      "stateMutability": "nonpayable",
      "type": "function"
    },
    {
      "constant": true,
      "inputs": [],
      "name": "deployedOn",
      "outputs": [
        {
          "name": "",
          "type": "uint256"
        }
      ],
      "payable": false,
      "stateMutability": "view",
      "type": "function"
    },
    {
      "constant": true,
      "inputs": [],
      "name": "owner",
      "outputs": [
        {
          "name": "",
          "type": "address"
        }
      ],
      "payable": false,
      "stateMutability": "view",
      "type": "function"
    },
    {
      "constant": true,
      "inputs": [
        {
          "name": "",
          "type": "address"
        }
      ],
      "name": "allowed",
      "outputs": [
        {
          "name": "",
          "type": "bool"
        }
      ],
      "payable": false,
      "stateMutability": "view",
      "type": "function"
    },
    {
      "inputs": [
        {
          "name": "notaryAddress",
          "type": "address"
        }
      ],
      "payable": false,
      "stateMutability": "nonpayable",
      "type": "constructor"
    },
    {
      "anonymous": false,
      "inputs": [
        {
          "indexed": true,
          "name": "hash",
          "type": "bytes32"
        },
        {
          "indexed": false,
          "name": "typ",
          "type": "uint8"
        },
        {
          "indexed": false,
          "name": "key",
          "type": "string"
        },
        {
          "indexed": true,
          "name": "initiator",
          "type": "address"
        }
      ],
      "name": "JourneyStart",
      "type": "event"
    },
    {
      "anonymous": false,
      "inputs": [
        {
          "indexed": true,
          "name": "parent",
          "type": "bytes32"
        },
        {
          "indexed": true,
          "name": "hash",
          "type": "bytes32"
        },
        {
          "indexed": false,
          "name": "typ",
          "type": "uint8"
        },
        {
          "indexed": false,
          "name": "key",
          "type": "string"
        }
      ],
      "name": "AssetLinked",
      "type": "event"
    },
    {
      "constant": false,
      "inputs": [
        {
          "name": "user",
          "type": "address"
        }
      ],
      "name": "addUser",
      "outputs": [],
      "payable": false,
      "stateMutability": "nonpayable",
      "type": "function"
    },
    {
      "constant": false,
      "inputs": [
        {
          "name": "user",
          "type": "address"
        }
      ],
      "name": "delUser",
      "outputs": [],
      "payable": false,
      "stateMutability": "nonpayable",
      "type": "function"
    },
    {
      "constant": false,
      "inputs": [
        {
          "name": "hash",
          "type": "bytes32"
        },
        {
          "name": "typ",
          "type": "uint8"
        },
        {
          "name": "key",
          "type": "string"
        }
      ],
      "name": "Start",
      "outputs": [],
      "payable": false,
      "stateMutability": "nonpayable",
      "type": "function"
    },
    {
      "constant": false,
      "inputs": [
        {
          "name": "parent",
          "type": "bytes32"
        },
        {
          "name": "hash",
          "type": "bytes32"
        },
        {
          "name": "typ",
          "type": "uint8"
        },
        {
          "name": "key",
          "type": "string"
        }
      ],
      "name": "Link",
      "outputs": [],
      "payable": false,
      "stateMutability": "nonpayable",
      "type": "function"
    }
  ]`
)

type CustomerJourneyStart struct {
	Hash common.Hash
	Typ  uint8
	Key  string
}

type CustomerJourneyLink struct {
	Parent common.Hash
	Hash   common.Hash
	Typ    uint8
	Key    string
}

func initABI() {
	notaryABI, _ = abi.JSON(strings.NewReader(notaryABIStr))
	customerJourneyABI, _ = abi.JSON(strings.NewReader(customerJourneyABIStr))
}

func DetermineType(hexInput string) (int, error) {
	//log.Debugf("input HEX value : %s" , hexInput)
	initABI()
	var funcType int
	//var docType int
	// Method Data.
	byteData, err := hex.DecodeString(hexInput[2:10])

	if err != nil {
		log.Errorf("1: Error Parsing TX InputData Value : %v", err)
		return 0, err
	}

	m, err := customerJourneyABI.MethodById(byteData)
	if err != nil {
		log.Errorf("2: Error Parsing TX InputData Value : %v", err)
		return 0, err

	}

	switch m.Name {

	case "Start":
		log.Debug("Start")
		funcType = 1
	case "Link":
		log.Debug("Link")
		funcType = 2
	}
	// (probably) non-notary contract }
	return funcType, nil
}
