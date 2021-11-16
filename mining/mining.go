package mining

import (
	"btc-node/blockchain"
	"btc-node/chaincfg"
	"btc-node/chaincfg/chainhash"
	"btc-node/txscript"
	"btc-node/wire"
	"time"

	"github.com/btcsuite/btcutil"
)

type TxDesc struct {
	Tx *btcutil.Tx

	Added time.Time

	Height int32

	Fee int64

	FeePerKB int64
}

type TxSource interface {
	LastUpdated() time.Time

	MiningDescs() []*TxDesc

	HaveTransaction(hash *chainhash.Hash) bool
}

type BlockTemplate struct {
	Block *wire.MsgBlock

	Fees []int64

	SigOpCosts []int64

	Height int32

	ValidPayAddress bool

	WitnessCommitment []byte
}

type BlkTmplGenerator struct {
	policy      *Policy
	chainParams *chaincfg.Params
	txSource    TxSource
	chain       *blockchain.BlockChain
	timeSource  blockchain.MedianTimeSource
	sigCache    *txscript.SigCache
	hashCache   *txscript.HashCache
}
