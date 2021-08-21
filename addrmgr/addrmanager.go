package addrmgr

import (
	"container/list"
	"math/rand"
	"net"
	"sync"
	"time"

	"btc-node/wire"
)

type localAddress struct {
	na    *wire.NetAddress
	score AddressPriority
}

type AddressPriority int

const (
	needAddressThreshold = 1000

	dumpAddressInterval = time.Minute * 10

	triedBucketSize = 256

	triedBucketCount = 64

	newBucketSize = 64

	newBucketCount = 1024

	triedBucketsPerGroup = 8

	newBucketsPerGroup = 64

	newBucketsPerAddress = 8

	numMissingDays = 30

	numRetries = 3

	maxFailures = 10

	minBadDays = 7

	getAddrMax = 2500

	getAddrPercent = 23

	serialisationVersion = 2
)

type AddrManager struct {
	mtx            sync.RWMutex
	peersFile      string
	lookupFunc     func(string) ([]net.IP, error)
	rand           *rand.Rand
	key            [32]byte
	addrIndex      map[string]*knownAddress
	addrNew        [newBucketCount]map[string]*knownAddress
	addrTried      [triedBucketCount]*list.List
	started        int32
	shutdown       int32
	wg             sync.WaitGroup
	quit           chan struct{}
	nTried         int
	nNew           int
	lamtx          sync.Mutex
	localAddresses map[string]*localAddress
	version        int
}
