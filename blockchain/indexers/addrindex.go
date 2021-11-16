package indexers

import (
	"btc-node/chaincfg"
	"btc-node/chaincfg/chainhash"
	"btc-node/database"
	"sync"

	"github.com/btcsuite/btcutil"
)

const (
	addrKeySize = 1 + 20
)

type CfIndex struct {
	db          database.DB
	chainParams *chaincfg.Params
}

type TxIndex struct {
	db         database.DB
	curBlockID uint32
}

type AddrIndex struct {
	db          database.DB
	chainParams *chaincfg.Params

	uncomfirmedLocked sync.RWMutex
	txnsByAddr        map[[addrKeySize]byte]map[chainhash.Hash]*btcutil.Tx
	addrsByTx         map[chainhash.Hash]map[[addrKeySize]byte]struct{}
}
