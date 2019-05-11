package bitcoin
//保存已知的peer地址 不包括不可达地址和洋葱地址
//将peer分组
//随机选择peer分组进行通信，保证peer多样性，避免只与被控制的恶意节点通信
//定期删去恶意节点

维护与外部的谅解， 选择peer， 禁止，限制最大连接数量，tor寻找

maxFailedAttempts 最大连续重连次数 如果到了最大次数还是失败，则要等一段时间再重连

ErrDialNil 用于指示configuration里的Dial不能为空

maxRetryDuration 尝试建立一个持久连接所允许的最长重试时间  这很有必要 因为重试时采用的是一种补偿机制  每次重试增大时间间隔

defaultRetryDuration 尝试持久化的连接使用的默认时间

dafaultTargetOutbound 需要维护的outbound连接的默认数量

ConnState 代表了被请求的连接的状态 可以是pending，established，disconnected，或者failed。当要请求一个新的连接时，该连接要么是established要么是failed。一个established的连接如果断开了就是disconnected

ConnReq 是向一个网络地址请求的连接  如果是永久性的，则断开连接时将重试连接

updateState 更新连接请求的状态

ID 返回连接请求的id

State 被请求的连接的状态

String 返回连接请求的易读字符串

Config 存着关于connection manager 的配置选项

registerPending 用于注册一个去pend connection的attempt。 这样如果不在想要的话，调用者能够取消pend connection的attmpt

handleConnected 用于将成功的连接放入队列

handleDisconnected 用于移除connection

handleFailed 用于移除pending connection

ConnManger 提供一个mananger用于处理网络连接

handleFailedConn 处理由于断连或者其他故障而引起的连接失败 如果是长久 它按照配置里设置的重试间隔进行重试。否则，它建立新的连接请求。在maxFailedConnectionAttempts之后，按照配置里的重试间隔，新的连接将会被重试

connHandler 处理所有与连接相关的请求 以goroutine允许 connHandler抱枕 我们维护了一个活跃的outbound 的连接池， 我们可以保持对网络的连接状态。 依据连接请求的id进行处理和分配

NewConnReq 建立一个新的连接请求 并且去连接相应的地址

Connect 指派一个id并拨打连接到连接请求的地址

Disconnect 断开与给定id的connection的连接 如果是永久性的，则将以增加的退避持续时间重试连接。

Remove 根据给定id从已知的连接里移除相应连接。 注意：此方法还可用于取消尚未成功的延迟连接尝试。

listenHandler 接受给定listener上的传入连接。 它必须作为goroutine运行。

Start 启动connection manager 并且开始连接到网络

Wait 阻塞 直到connection manager 正常停止

Stop  正常关闭连接管理器。

New 返回一个新的connection manager。使用Start去开始连接到网络。





