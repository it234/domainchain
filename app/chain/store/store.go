package store

import (
	"sync"

	"domainchain/pkg"

	dbm "github.com/tendermint/tm-db"
)

var (
	DbBlk *DataModel
)

type DataModel struct {
	DB    dbm.DB
	Mutex *sync.Mutex
}

func InitBlkDb() (err error) {
	DbBlk = new(DataModel)
	DbBlk.Mutex = &sync.Mutex{}
	dbDir := pkg.GetRootDir() + "confdata/db/block"
	DbBlk.DB, err = dbm.NewDB("block", dbm.GoLevelDBBackend, dbDir)
	return
}
