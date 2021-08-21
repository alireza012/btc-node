package chaincfg

import (
	"btc-node/chaincfg/chainhash"
	"btc-node/wire"
	"math/big"
	"time"
)

type Checkpoint struct {
	Height int32
	Hash   *chainhash.Hash
}

type DNSSeed struct {
	Host string

	HasFiltering bool
}

type ConsensusDeployment struct {
	BitNumber uint8

	StartTime uint64

	ExpireTime uint64
}

const (
	DeploymentTestDummy = iota

	DeploymentCSV

	DeploymentSegwit

	DeploymentTaproot

	DefinedDeployments
)

type Params struct {
	Name string

	Net wire.BitcoinNet

	DefaultPort string

	DNSSeeds []DNSSeed

	GenesisBlock *wire.MsgBlock

	GenesisHash *chainhash.Hash

	PowLimit *big.Int

	PowLimitBits uint32

	BIP0034Height int32
	BIP0065Height int32
	BIP0066Height int32

	CoinbaseMaturity uint16

	SubsidyReductionInterval int32

	TargetTimespan time.Duration

	TargetTimePerBlock time.Duration

	RetargetAdjustmentFactor int64

	ReduceMinDifficulty bool

	MinDiffReductionTime time.Duration

	GenerateSupported bool

	Checkpoints []Checkpoint

	RuleChangeActivationThreshold uint32
	MinerConfirmationWindow       uint32
	Deployments                   [DefinedDeployments]ConsensusDeployment

	RelayNonStdTxs bool

	Bech32HRPSegwit string

	PubKeyHashAddrID        byte
	ScriptHashAddrID        byte
	PrivateKeyID            byte
	WitnessPubKeyHashAddrID byte
	WitnessScriptHashAddrID byte

	HDPrivateKeyID [4]byte
	HDPublicKeyID  [4]byte

	HDCoinType uint32
}
