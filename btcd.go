package main

import (
	"btc-node/limits"
	"fmt"
	"os"
	"runtime/debug"
	"vqaap-engine/server"
)

func btcdMain(serverChan chan<- )

func main() {
	debug.SetGCPercent(10)

	if err := limits.SetLimits(); err != nil {
		fmt.Fprintf(os.Stderr, "failed to set limits: %v\n", err)
		os.Exit(1)
	}
}
