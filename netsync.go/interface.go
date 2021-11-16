package netsync

import (
	"btc-node/chaincfg/chainhash"
	"btc-node/mempool"
	"btc-node/peer"
	"btc-node/wire"

	"github.com/btcsuite/btcutil"
)

type PeerNotifier interface {
	AnnounceNewTransactions(newTxs []*mempool.TxDesc)

	UpdatePeerHeights(latestBlkHash *chainhash.Hash, latestHeight int32, updateSource *peer.Peer)

	RelayInventory(invVect *wire.InvVect, data interface{})

	TransactionConfirmed(tx *btcutil.Tx)
}
