# Migrator for DELIVER Pilot
A Migration tool to deliver the txs to the correct chain

## Introduction.

The `migrator` is a simple "tool" that can be used to index locally all the transactions that associated with a specific 
address given a **range** of blocks to search for. It uses the APIs from [Etherscan](https://rinkeby.etherscan.io/) to 
get the job done. 

The main features of the `migrator` is:
 - Poll data from a specific blockchain
 - Resign the data with another private key and send it to a specific address, with no changes made to the *data* of the original data.
 
The intention of the `migrator` is to faciliate the migration the data for the pilot.

### building the module

Simply run to build the `migrator`:
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

The return data is then stored in the format below.

```go
package njson 

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

#### Using the migrator to poll data.

In the current iteration, the migrator can only poll data. *REINDEXING DATA is still a WIP*.
The input for polling is `start` and `range`.
- `start` : starting block_height for the poll
- `range` : the number of blocks the query tx for an *address*

[remark] The API can only poll up to 10,000 TX. This should be accounted for when deciding the range of the polls.

To poll data from the network currently (set to `rinkeby`), simply run command as follows:
```shell script
# Set start and range.
./migrator -start=5000735 -range=5000
``` 
The command will create a db (default: `migrator.db`) and index the txs there. If the tx with the same `txid` already exists, 
the tx will not be indexed.          

[WIP] Recreating the data within the blockchain
 - Currently there are only a set of test codes within the *blockchain/connector_test.go* for recreating the tx. Will implement a batch-like method to upload all the txs together


### NOTES:

***Currently synced (DB at least) to block number 5350304***