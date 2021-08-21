package addrmgr

import (
	"btc-node/wire"
	"time"
)

type knownAddress struct {
	na          *wire.NetAddress
	srcAddr     *wire.NetAddress
	attempts    int
	lastattempt time.Time
	lastsuccess time.Time
	tried       bool
	refs        int
}
