package blockchain

import (
	"math/big"
	"sync"

	"github.com/btcsuite/btcd/chaincfg"
	"github.com/btcsuite/btcd/chaincfg/chainhash"
	"github.com/btcsuite/btcd/database"
)

type blockStatus byte

type blockNode struct {
	parent *blockNode

	hash chainhash.Hash

	workSum *big.Int

	height int32

	version    int32
	bits       uint32
	nonce      uint32
	timestamp  int64
	merkleRoot chainhash.Hash

	status blockStatus
}

type blockIndex struct {
	db          database.DB
	chainParams *chaincfg.Params

	sync.RWMutex
	index map[chainhash.Hash]*blockNode
	dirty map[*blockNode]struct{}
}
