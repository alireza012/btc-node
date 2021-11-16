package database

import (
	"github.com/btcsuite/btclog"
)

type Driver struct {
	DbType string

	Create func(args ...interface{}) (DB, error)

	Open func(args ...interface{}) (DB, error)

	UseLogger func(logger btclog.Logger)
}

var drivers = make(map[string]*Driver)

func SupportedDrivers() []string {
	supperotedDBs := make([]string, 0, len(drivers))
	for _, drv := range drivers {
		supperotedDBs = append(supperotedDBs, drv.DbType)
	}

	return supperotedDBs
}
