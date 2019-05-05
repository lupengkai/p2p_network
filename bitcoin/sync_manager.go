package bitcoin
/*The SyncManager communicates with connected peers to perform an initial block download,
keep the chain and unconfirmed transaction pool in sync, and announce new blocks connected to the chain.
Currently the sync manager selects a single sync peer that it downloads all blocks from
until it is up to date with the longest chain the sync peer is aware of.*/

//从其他peer下载block