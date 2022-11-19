package chain

import (
	"flag"
	"fmt"
	"os"

	"domainchain/app/chain/blk"
	"domainchain/app/chain/store"

	abciserver "github.com/tendermint/tendermint/abci/server"
	"github.com/tendermint/tendermint/libs/log"
	"github.com/tendermint/tendermint/libs/service"
)

var socketAddr string

func init() {
	flag.StringVar(&socketAddr, "socket-addr", "127.0.0.1:9080", "Unix domain socket address")
}

func Run() (server service.Service) {
	flag.Parse()

	fmt.Println("socketAddr:", socketAddr)
	err := store.InitBlkDb()
	if err != nil {
		panic(err)
	}

	app := blk.NewDomainApplication()
	logger := log.MustNewDefaultLogger(log.LogFormatPlain, log.LogLevelInfo, false)
	server = abciserver.NewSocketServer(socketAddr, app)
	server.SetLogger(logger)
	if err := server.Start(); err != nil {
		fmt.Fprintf(os.Stderr, "error starting socket server: %v", err)
		os.Exit(1)
	}
	return
}
