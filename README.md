# Migrator for DELIVER Pilot
A Migration tool to deliver the txs to the correct chain

## Introduction.

The `migrator` is a simple "tool" that can be used to index locally all the transactions that associated with a specific 
address given a **range** of blocks to search for. It uses the APIs from [Etherscan](https://rinkeby.etherscan.io/) to 
get the job done. 


NOTES:

Currently synced (DB at least) to block number **5350304**