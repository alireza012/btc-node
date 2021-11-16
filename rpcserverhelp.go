package main

import "sync"

type helpCacher struct {
	sync.Mutex
	usage      string
	methodHelp map[string]string
}
