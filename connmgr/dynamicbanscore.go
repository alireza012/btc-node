package connmgr

import "sync"

type DynamicBanScore struct {
	lastUnix   int64
	transient  float64
	persistent uint32
	mtx        sync.Mutex
}
