# Migrator for DELIVER Pilot
A Migration tool to deliver the txs to the correct chain

## Introduction.

The `migrator` is a simple "tool" that can be used to index locally all the transactions that associated with a specific 
address given a **range** of blocks to search for. It uses the APIs from [Etherscan](https://rinkeby.etherscan.io/) to 
get the job done. 

The main features of the `migrator` is:
 - Poll data from a specific blockchain
 - Resign the data with another private key and send it to a specific address, with no changes made to the *data* of the original data.
 
The purpose of the `migrator` is to simple data migration purposes.

### building the module

Simply run:
 ```shell script
 go build 
```

### Polling data from the blockchain. 

`migrator` uses **Etherscan** API to poll blocks and save it to a local db, specifically the API below
```
# API request to get data from etherscan
http://api-rinkeby.etherscan.io/api?module=account&action=txlist&
address=${address}}
&startblock=${startblock}
&endblock=${endblock}
&sort=asc&apikey=${API_KEY}
```
The inputs for the API is as follows :

 1. address for querying the txs
 2. starting block
 3. ending block
 4. API_Key for *Etherscan*

The return data is then stored in the current format.

```go

type TxResult struct {
	BlockNumber uint64		`json:"blockNumber,string"`
	TxHash      string 		`json:"hash"`
	Nonce       uint64 		`json:"nonce,string"`
	From        string 		`json:"from"`
	To          string 		`json:"to"`
	Gas         uint64 		`json:"gas,string"`
	GasPrice    uint64 		`json:"gasPrice,string"`
	Input       string 		`json:"input"`
}

```


[WIP] Recreating the data within the blockchain
 - Currently withinng the *blockchain/connector_test.go* . Will implement a batch-like method to upload all the txs together


### NOTES:

***Currently synced (DB at least) to block number 5350304***