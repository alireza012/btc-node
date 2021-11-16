package main

import (
	"btc-node/limits"
	"fmt"
	"os"
	"runtime"
	"runtime/debug"
)

var winServiceMain func() (bool, error)

func btcNodeMain(serverChan chan<- *server) error {

	return nil
}

func main() {
	debug.SetGCPercent(10)

	if err := limits.SetLimits(); err != nil {
		fmt.Fprintf(os.Stderr, "failed to set limits: %v\n", err)
		os.Exit(1)
	}

	if runtime.GOOS == "windows" {
		isService, err := winServiceMain()
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		if isService {
			os.Exit(0)
		}
	}

	if err := btcNodeMain(nil); err != nil {
		os.Exit(1)
	}
}
