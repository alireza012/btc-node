package mining

import "github.com/btcsuite/btcutil"

type Policy struct {
	BlockMinWeight uint32

	BlockMaxWeight uint32

	BlockMinSize uint32

	BlockMaxSize uint32

	BlockPrioritySize uint32

	TxMinFreeFee btcutil.Amount
}
