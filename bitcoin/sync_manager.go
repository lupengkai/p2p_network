package bitcoin

//同步协议

//1.新加入节点上的区块链初始化
//2.同步别人挖出的区块
//2.同步未确认的交易
//2.发布新挖出的区块

resetHeaderState 记录同步到第几块
findNextHeaderCheckpoint 初始化同步的时候分阶段，一个阶段一个checkpoint，这个返回下一个checkpoint

startSync 选择最佳的peer去同步区块链 同时删去不合适的候选peer （涉及到peer选择算法）
isSyncCandidate 返回是不是考虑和一个peer同步
handleNewPeerMsg 处理新的peer, 这些peer有潜力成为候选的同步peer（可能就是分分组吧，猜的）
handleDonePeerMsg 处理发出“done”消息的peer，将他们移出同步候选节点，如果是选择的当前同步节点，则换一个

handleTxMsg 处理来自所有peer的交易消息
current 返回 是 我们已经同步完毕 否 还有块没有同步完成
handleBlockMsg 处理来自所有peer的区块消息
handleHeadersMsg 处理来自所有peer的区块头消息 当采用头部优先的同步模式时会用到

haveInventory 返回 是否 被传递的目录inv已经有了 检查目录里的区块是否在主链、侧链、孤儿池 检查目录里的交易是否在内存池、孤儿池

limitMap 将容量已经到达最大的map里随机删掉一个条目

blockHandler 是一个协程 是一个独立于peer handler以外的协程去处理block 和 inv消息, 这样不会一直在那里等待处理

handleBlockchainNotification 处理来自blockchain的通知。 请求找到孤儿区块的父母，传递已经接收的区块给其他相连的节点

NewPeer 通知sync manger 来了一个新peer

QueueTx 将传过来的交易信息和peer添加到区块处理队列。

QueueBlock 将传过来的区块信息和peer添加到区块处理队列。

QueueInv 将传过来的inv和peer添加到区块处理队列。

QueueHeaders 将传过来的头部信息添加到区块处理队列。

DonePeer 通知block manager 失去了对一个peer的连接

Start 启动核心的区块处理handler  用以处理区块和inv信息

Stop 停止sync manager 通过停止所有有异步的handle 并等待他们结束

SyncPeerID 返回 当前挑选出的同步peer的id 或者0

ProcessBlock 利用ProcessBlock处理区块链内部的一个实例

IsCurrent 返回 sync manager是否觉得他正在和相连的peer同步

Pause 暂停sync manager 直到返回的channel关闭  暂停的时候 所有的peer和block处理停止，不应该暂停太久

New 产生一个新的sync manager 调用Start方法开始异步处理bloc、tx和inv的更新