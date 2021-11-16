package blockchain

import "sync"

type chainView struct {
	mtx   sync.Mutex
	nodes []*blockNode
}
