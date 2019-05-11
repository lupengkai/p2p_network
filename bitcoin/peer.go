package bitcoin
//provides a common base for creating and managing bitcoin network peers.
//connect_manager -> peer -> sync_manager  (同步是其中一个功能）
//
MaxProtocolVersion peer能够支持的最大协议版本号

DefaultTrickleInterval 尝试发送inv信息给一个peer的最小间隔时间

MinAcceptableProtocolVersion  相连的peer所支持的最小协议版本号

outputBufferSize 输出channel使用的elements数量

maxInvTrickleSize 向remote peer 传送信息时单个信息里所包含inv的最大数量

maxKnownInventory 在已知的inv的缓存里保存条目的最大数量

pingInterval 发送ping的间隔时间

negotiateTimeout  一个没有完成最初的版本协定的peer没有动静超过30s 判断为超时

idleTimeout  一个peer没有动静超过5分钟 判断为超时

stallTickInterval 相邻两次检查peer是否stalled的时间间隔 15s

stallResponseTimeout

nodeCount  自启动开始拥有过的peer连接数量，也用这个给peer编号

zeroHash 值全是0的hash

sentNonces 存储push版本信息时产生的nonce信息 用来探测自己的连接

allowSelfConns 测试用 测试与自己的连接

MessageListeners 定义回调函数

OnGetAddr 当peer接收到 getaddr 时调用

OnAddr 当peer接收到 addr时调用

OnPing  当peer接收到ping时调用

OnPong 当peer接收到pong时调用

OnAlert 当peer接收到alert时调用

OnMemPool 当peer接收到memoool消息时调用

OnTx 当peer接收到Tx时调用


OnBlock 当peer接收到block时调用

OnCFilter 当收到 cfilter消息时调用

OnCFHeaders 当收到cfheader时调用

OnCFCheckpt 当收到 cfcheckpt时调用

OnInv 当收到inv时调用

OnHeaders 当收到header时调用

OnNotFound 当收到notfound时调用

OnGetData 当收到 getdata时调用

OnGetBlocks 当收到getblocks时调用

OnGetHeaders 当收到getheaders时调用

OnGetCFilters 当收到getcfilters时调用

OnGetCFHeaders 当收到getcheaders时调用

OnGetCFCheckpt 当收到getccfcheckpt时调用

OnFeeFilter 当接收到feelfilter时调用

OnFilterAdd 当接收到filteradd时调用

OnFilterClear 当接收到filterclear时调用

OnFilterLoad 当接收到filterload时调用

OnMerkleBlock 当接收到merkleblock时调用

OnVersion 当收到version时调用 可能会返回拒绝的信息，然后断开连接

OnVerAck 当接收到verack时调用

OnReject 当接收到rejiedect时调用

OnSendHeaders 当收到sendheaders时调用

OnRead 当收到消息时调用 包含要读的的字节数量，消息本身，读的时候是否出错  一般调用者选择专门的消息类型，但是这种用法可以用于没有peer没有提供特定的回调函数时使用

OnWrite 当发送消息时调用 可以用来记个数 或者其他任务


Config 包含有用的配置选项

NewestBlock 指定一个回调函数，这个回调函数 在peer需要时能够提供最新区块的细节信息给peer。当区块链高度是0时返回nil。最好实现一下这个方法，保证检查最新区块时不会出错

HostToNetAddress 返回给定host的网址。可以返回nil 这时候host名字作为ip地址处理

Proxy 使用tor代理时使用 表示是通过tor代理进行连接的

UserAgentName 指定advertise user agent的名字

UserAgentVersion 指定advertise时 user agent的版本号

UserAgentComments 指定advertise时 user agent 的注释

ChainParams 标识出peer与哪些关于chain的参数有关，测试网络的环境下可能被忽略

Services 指定advertise时本地peer支持的哪种service

ProtocolVersion 指定使用的和advertise时支持的最大版本号

DisableRelayTx 指定是否要通知remote peer不要再发送transaction的inv信息

Listeners 包含一些回调函数，这些函数在收到peer message时被调用

TrickleInterval 发送inv给一个peer的计时

minUint32 返回两个unint32数较小的一个

newNetAddress 尝试从net.Addr接口里收取出IP地址，并用这个地址产生一个bitcoin NetAddress结构的对象

outMsg 包含一个要发送的信息 和一个channel，用以指示什么时候发出的这条信息

stallControlCmd 代表了stall的控制命令

sccSendMessage 表明一个消息已经发送给其他peer

sccReceiveMessage 表明从其他peer处接收到一个消息

sccHandlerStart 表明一个回调函数即将被调用

sccHandlerDone 表明一个回调函数已经完成

stallControlMsg 用来指示关于特定事件的stall handler，这样可以探测和处理stalled的其他peer

StatsSnap 某个时间点上peer stats的快照

HashFunc 返回一个块的hash和高度， 是get newest block details 的回调函数

AddrFunc 输入一个地址 返回一个有关的地址

HostToNetAddrFunc 输入 主机 端口 服务 返回 netaddress

一个peer的数据流分成3个goroutine
流入的数据被inHandler的goroutine读到，然后分发到相应的handler 包括block transaction inventory
流出的数据用两个goroutine处理，queueHandler和outHandler
queueHandler 将外部的实体存入queue message ，通过QueueMessage函数 可以不用管peer是不是正在向外发消息
outHandler应该是实际负责向外发消息的，peer把消息放入queue message 就不管了， outHandler从queue里取消息发出去

Peer提供全双工读写，自动处理初始化时的握手过程，查询使用时的统计数据和其他关于remote peer的信息 比如地址，user agent， 协议版本
还负责 output message queuing， inventory trickling 以及动态注册和注销用以处理信息的回调函数。

流出的信息 一般用QueueMessage 或者QueueInventory放入队列。
QueueMessage处理所有类型的信息，包含对接收到block和transaction之后的replay
QueueInventory只负责传递inventory， 他利用trickling机制将inventory在一起批处理。
另外，还有一些因为经常使用，为了方便而写的helper函数

Peer


String 以易读的字符串表示peer的地址和directionality


UpdateLastBlockHeight 更新peer最新知道的区块高度 并发安全（用来mutex）


UpdateLastAnnouncedBlock 更新peer知道的关于刚产生的block的hash信息 并发安全

AddKnownInventory 将传来的inventory放到 已知inventory的缓存里 并发安全

StatsSnapshot 返回当前时刻下peer的flag和统计数据的快照

ID 返回peer的id 并发安全

NA 返回peer的网络地址 并发安全

Addr 返回peer的地址 并发安全

Inbound 返回peer是否inbound

Services 返回remote peer的service flag 并发安全

UserAgent 返回remote peer的user agent 并发安全

LastAnnouncedBlock 返回remote peer的最后发布的block 并发安全

LastPingNonce 返回对remote peer发出最后一次ping的nonce 并发安全 如果有pending ping的话

LastPingTime  返回对remote peer发出最后一次ping的时间 并发安全

LastPingMicros 返回对remote peer发出最后一次ping的响应回复时间 并发安全

VersionKnown 是否本地已经知道peer的版本号 并发安全

VerAckReceived 返回peer是否收到了version 的回复信息 verack

ProtocolVersion 返回协商的peer协议版本

LastBlock 返回peer的最新区块 并发安全

LastSend 返回peer上次发送的时间


LastRecv 返回peer上次接收的时间

LocalAddr 返回连接的本地地址 并发安全

BytesSent 返回peer发出的所有byte数 并发安全

BytesReceived  返回所有收到的byte数 并发安全

TimeConnected 返回peer连接上的时间

TimeOffset 返回自peer在初始协商阶段报告的时刻与本地相差的秒数，负值代表remote peer的时间早于本地时间 并发安全

StartingHeight 返回在初始协商阶段peer报告的最新已知高度 并发安全

WantsHeaders 返回peer师傅偶像要头部信息，而不是区块的inventory信息 并发安全

IsWitnessEnabled 当peer已经表示它支持隔离的见证者时返回true 并发安全

PushAddrMsg 使用提供的一些地址向连接上的peer发送地址消息。这个行数在通过queuemessage手动发送消息时是有用的，因为可以自动限制地址的数量，防止超多最大值。它返回了实际发出去的那些地址  并发安全

PushGetBlocksMsg 根据提供的区块的locator和stop hash发送一条getblocks的消息 它会忽略紧接着的重复请求

PushGetHeadersMsg 根据提供的区块的locator和stop hash发送一条getheaders的消息 它会忽略紧接着的重复请求

PushRejectMsg 根据提供的命令发送拒绝信息，拒绝的代号，拒绝的理由，和hash  只有当命令是tx或者block时，hash有值  参数wait将会导致该函数阻塞直到拒绝信息真的已经发出去了 并发安全

handlePingMsg 当一个peer收到ping的消息时调用。对于最近的客户端，它将回复pong的消息 对于老版本的客户端，它不做任何事

handlePongMsg  当一个peer收到pong的消息时调用 它将按照最近客户端的需求更新ping的统计数据， 对老客户端和一个之前没有发出去的ping不会有影响

readMessage  从peer那儿读取下一个消息 并记录下来

writeMessage 发送一个消息给peer并记录

isAllowedReadError 返回 传过来的错误在没有断开与peer的连接时是否被允许 特别的回归测试需要被允许在没有与peer断开连接时发送格式错误的邮件

shouldHandleReadError 返回 传过来的错误是否应该被记录并且回复一个拒绝的消息，该错误应该是要来自inHandler里对remote peer的读操作

maybeAddDeadline  为命令的回复设置一个可能合适的截止时间

stallHandler 为peer处理stall问题的检测 这需要跟踪预期的回复和指派的截止时间，同时考虑回调所花费的时间。它必须以goroutine的形式运行

inHandler 为peer处理所有流入的消息。以goroutine方式运行

queueHandler 为peer把传出的数据放入队列 为各种输入源调用，这样不会阻塞消息的发送。

shouldLogWriteError 返回传入的错误是否应该被记录下来，该错误应该是来自outHandler里向remote peer的写操作

outHandler 为peer处理所有流出的消息 必须以goroutine的形式处理 它用一个个缓存的channel去序列化输出的消息，同时允许发送者可以异步的继续运行

pingHandler 阶段性地ping peer。 以goroutine运行

QueueMessage 将传入的消息放到要发送给peer的消息队列中去 并发安全

QueueMessageWithEncoding  将传入的消息放到要发送给peer的消息队列中去 这个函数等同于QueueMessage，但是它允许调用者指定编码的类型

QueueInventory 将传入的inventory放到inventory queue中去，该queue可能不会立即发送，而是分批流向peer。

Connected 返回当前是否和peer连接着 并发安全

Disconnect 关闭连接以断开连接。当peer已经处于断开状态或者正在断开时，调用此函数无效

readRemoteVersionMsg 等待下一个来自remote peer的消息到达。 如果下一个消息不是版本信息 或者版本时不可接受的，返回一个error

localVersionMsg 创造一个可以发送给remote peer的版本信息

writeLocalVersionMsg 将我们的版本信息发送给remote peer

negotiateInboundProtocol 等待接收一个来自peer的版本信息，然后发送我们的版本信息。如果没有按顺序来就报错。

negotiateOutboundProtocol 发送版本信息然后等待接收来自peer的版本信息。如果没有按顺序来就报错

start 开始处理输入和输出的消息

AssociateConnection 将给定的conn与peer联系在一起。当这个peer已经连接上时调用这个函数没有影响

WaitForDisconnect  一直等到与peer完全断开连接，并清理完占用的资源。如果本地或者remote peer其中一个断开，或者peer通过Disconnect强制断开时 会被调用

newPeerBase 返回一个基于inbound flag的新的基本peer。这个被NewInboundPeer和NewOutboundPeer函数用于执行两种peer都需要的基本设置

NewInboundPeer  返回一个新的inbound的peer。使用start开始处理流入和流出的消息

NewOutboundPeer 返回一个新的outbound peer。

init