package store

import (
	"context"
	"encoding/hex"
	"encoding/json"
	"log"
	"strconv"
	"strings"
	"sync"
	"time"

	"domainchain/pkg"
	"domainchain/proto/chain/blk"

	"github.com/gogo/protobuf/proto"
	"github.com/robfig/cron/v3"
	"github.com/tendermint/tendermint/crypto"
	client "github.com/tendermint/tendermint/rpc/client/http"
	dbm "github.com/tendermint/tm-db"
)

var (
	DbBlk *DataModel
	Rmote = "http://localhost:26657"
)

const (
	KeyCurrentBlockHight = "KeyCurrentBlockHight"
	KeyBlockHightPrefix  = "KeyBlockHightPrefix_"
	KeyTxHashPrefix      = "KeyTxHashPrefix_"
)

type DataModel struct {
	DB                dbm.DB
	Mutex             *sync.Mutex
	CurrentBlockHight int64
}

type Block struct {
	Hight  int64
	Hash   string
	TxHash []string
}

func Store() {
	err := InitBlkDb()
	if err != nil {
		panic(err)
	}
	getBlockData()
	return
}

func InitBlkDb() (err error) {
	DbBlk = new(DataModel)
	DbBlk.Mutex = &sync.Mutex{}
	dbDir := pkg.GetRootDir() + "browser/db/block"
	DbBlk.DB, err = dbm.NewDB("block", dbm.GoLevelDBBackend, dbDir)
	return
}

func getBlockData() {
	v, err := DbBlk.DB.Get([]byte(KeyCurrentBlockHight))
	if err != nil {
		panic(err)
	}
	var currentBlockHight int64
	currentBlockHight, _ = strconv.ParseInt(string(v), 0, 0)
	DbBlk.CurrentBlockHight = currentBlockHight

	timezone, err := time.LoadLocation("UTC")
	if err != nil {
		timezone = time.Local
	}
	cc := cron.New(cron.WithSeconds(), cron.WithLocation(timezone))
	spec := "*/5 * * * * ?"
	_, err = cc.AddJob(spec, cron.NewChain(
		cron.SkipIfStillRunning(cron.DefaultLogger),
		cron.Recover(cron.DefaultLogger),
	).Then(&taskCheckBlock{}))
	cc.Start()
}

type taskCheckBlock struct {
}

func (this taskCheckBlock) Run() {
	cli, err := client.New(Rmote)
	if err != nil {
		log.Println(err)
		return
	}
	ctx := context.Background()
	resp, err := cli.Status(ctx)
	if err != nil {
		log.Println(err)
		return
	}
	if resp.SyncInfo.LatestBlockHeight <= DbBlk.CurrentBlockHight {
		return
	}
	for i := DbBlk.CurrentBlockHight + 1; i <= resp.SyncInfo.LatestBlockHeight; i++ {
		err = getBlockDataByHeight(i)
		if err != nil {
			log.Println(err)
			return
		}
	}
}

func getBlockDataByHeight(height int64) (err error) {
	cli, err := client.New(Rmote)
	if err != nil {
		log.Println(err)
		return
	}
	ctx := context.Background()
	resp, err := cli.Block(ctx, &height)
	if err != nil {
		log.Println(err)
		return
	}
	DbBlk.Mutex.Lock()
	defer DbBlk.Mutex.Unlock()
	dbb := DbBlk.DB.NewBatch()
	defer dbb.Close()
	blkModel := &Block{}
	blkModel.Hight = height
	blkModel.Hash = resp.BlockID.String()
	for _, txByte := range resp.Block.Data.Txs {
		txHash := strings.ToUpper(hex.EncodeToString(crypto.Sha256(txByte)))
		txM := &blk.Tx{}
		err = proto.Unmarshal(txByte, txM)
		if err != nil {
			log.Println(err)
			continue
		}
		blkModel.TxHash = append(blkModel.TxHash, txHash)
		err = dbb.Set([]byte(KeyTxHashPrefix+txHash), txByte)
		if err != nil {
			log.Println(err)
			return
		}
	}
	blkByt, err := json.Marshal(blkModel)
	if err != nil {
		log.Println(err)
		return
	}
	err = dbb.Set([]byte(KeyBlockHightPrefix+strconv.FormatInt(height, 10)), blkByt)
	if err != nil {
		log.Println(err)
		return
	}
	err = dbb.Set([]byte(KeyCurrentBlockHight), []byte(strconv.FormatInt(height, 10)))
	if err != nil {
		log.Println(err)
		return
	}
	err = dbb.Write()
	if err != nil {
		log.Println(err)
		return
	}
	DbBlk.CurrentBlockHight = height
	return
}
