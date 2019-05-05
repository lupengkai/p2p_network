#layer 1
##storage
###miner
###wallet
+ get balance
+ network

##network 
*communication between nodes*

**kinds of nodes**
+ miner(wallet)
+ wallet

**relationship**
+ miner-wallet
+ miner-miner

###miner-client
####wallet 
+ get miner address
+ send transaction to miner
+ request blocks from miner
+ download blocks from miner

####miner
+ get other miners' address
+ receive transaction
+ overlay received transaction
+ publish a mined block to other miners
+ received blocks
+ overlay blocks
+ send blocks to wallet
+ relay wallets' request of blocks
+ send blocks to new-joined miners
+ downloads from old miners

###miner-miner
####miner


