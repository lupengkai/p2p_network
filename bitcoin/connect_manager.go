package bitcoin
/*Connection Manager handles all the general connection concerns such as maintaining a set number of outbound connections, sourcing peers, banning, limiting max connections, tor lookup, etc.

The package provides a generic connection manager which is able to accept connection requests from a source or a set of given addresses, dial them and notify the caller on connections. The main intended use is to initialize a pool of active connections and maintain them to remain connected to the P2P network.

In addition the connection manager provides the following utilities:

Notifications on connections or disconnections
Handle failures and retry new addresses from the source
Connect only to specified addresses
Permanent connections with increasing backoff retry timers
Disconnect or Remove an established connection

*/
//接收连接请求和发起连接

import (
	"errors"
	"fmt"
	"net"
	"sync"
	"sync/atomic"
)