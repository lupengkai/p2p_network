wire 保存协议规定网络协议 消息格式（以及包含消息的封装解析的方法）
Handlers 根据共识算法处理到来的信息
Peers Manager 管理所有peer信息
TCP server： 监听本地端口等待其他节点连接 或者主动连接 发出请求 或者把请求装发到handlers




上层调用下层的服务

TCP->Peer->Consensus

Peer层包含addrmgr

缓存inventory(回复或者通知对方自己已经知道的Tx或者Block)
缓存nonce(nonce用来判断version消息是不是发给自己的)

网络和存储



Pee
TCP
