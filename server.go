package main

import (
	"btc-node/addrmgr"
	"btc-node/chaincfg"
	"btc-node/connmgr"
	"btc-node/txscript"
)

type server struct {
	bytesReceived uint64
	bytesSent     uint64
	started       int32
	shutdown      int32
	shutdownSched int32
	startupTime   int64

	chainParams *chaincfg.Params
	addrManager *addrmgr.AddrManager
	connManager *connmgr.ConnManager
	sigCache    *txscript.SigCache
	hashCache   *txscript.HashCache
}
