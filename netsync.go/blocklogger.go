package netsync

import (
	"sync"
	"time"

	"github.com/btcsuite/btclog"
)

type blockProgressLogger struct {
	receivedLogBlocks int64
	receivedLogTx     int64
	lastBlockLogTime  time.Time

	subsystemLogger btclog.Logger
	progressAction  string
	sync.Mutex
}
