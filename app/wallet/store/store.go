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

	kf "github.com/tendermint/tendermint/crypto/secp256k1"
)

var (
	DbBlk *DataModel
	Rmote = "http://localhost:26657"
)

const (
	KeyCurrentBlockHight = "KeyCurrentBlockHight"
	KeyAddrData          = "KeyAddrData"
)

type DataModel struct {
	DB                dbm.DB
	Mutex             *sync.Mutex
	CurrentBlockHight int64
	AddrData          map[string]AddrInfo
}

type AddrInfo struct {
	Addr       string `json:"addr"`
	PrivateKey string `json:"private_key"`
	Balance    uint64 `json:"balance"`
	TxList     map[string]blk.Tx
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
	dbDir := pkg.GetRootDir() + "wallet/db/block"
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

	addrData := make(map[string]AddrInfo)
	v, err = DbBlk.DB.Get([]byte(KeyAddrData))
	if err != nil {
		panic(err)
	}
	if len(v) > 1 {
		err = json.Unmarshal(v, &addrData)
		if err != nil {
			panic(err)
		}
	}
	DbBlk.AddrData = addrData

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
	// 事务内处理
	dbb := DbBlk.DB.NewBatch()
	defer dbb.Close()

	isNeedSaveAddrData := false
	for _, txByte := range resp.Block.Data.Txs {
		txHash := strings.ToUpper(hex.EncodeToString(crypto.Sha256(txByte)))
		txBase := &blk.Tx{}
		err = proto.Unmarshal(txByte, txBase)
		if err != nil {
			log.Println(err)
			continue
		}
		fromAddr := kf.PubKey(txBase.FromPubKey).Address().String()
		toAddr := ""
		var amount uint64
		if txBase.TxType == blk.TxType_ISSUE {
			txM := &blk.TxIssue{}
			err = proto.Unmarshal(txBase.Payload, txM)
			if err != nil {
				log.Println(err)
				continue
			}
			toAddr = crypto.Address(txM.ToAddr).String()
			amount = txM.Amount
		} else if txBase.TxType == blk.TxType_TRANSFER {
			txM := &blk.TxTransfer{}
			err = proto.Unmarshal(txBase.Payload, txM)
			if err != nil {
				log.Println(err)
				continue
			}
			toAddr = crypto.Address(txM.ToAddr).String()
			amount = txM.Amount
		} else {
			continue
		}
		handleFromAddr := ""
		fromAddrInfoModel := AddrInfo{}
		handleToAddr := ""
		toAddrInfoModel := AddrInfo{}
		for k, v := range DbBlk.AddrData {
			if k == fromAddr {
				handleFromAddr = fromAddr
				fromAddrInfoModel = v
				fromAddrInfoModel.Balance = fromAddrInfoModel.Balance - amount
				fromAddrInfoModel.TxList[txHash] = *txBase
				DbBlk.AddrData[handleFromAddr] = fromAddrInfoModel
				isNeedSaveAddrData = true

				log.Println("发送地址:", fromAddr, ",发送金额:", amount)
			}
			if k == toAddr {
				handleToAddr = toAddr
				toAddrInfoModel = v
				toAddrInfoModel.Balance = toAddrInfoModel.Balance + amount
				toAddrInfoModel.TxList[txHash] = *txBase
				DbBlk.AddrData[handleToAddr] = toAddrInfoModel
				isNeedSaveAddrData = true

				log.Println("接收地址:", toAddr, ",接收金额:", amount)
			}
		}
	}

	if isNeedSaveAddrData {
		addrData, err1 := json.Marshal(DbBlk.AddrData)
		if err1 != nil {
			err = err1
			log.Println(err)
			return
		}
		err = DbBlk.DB.Set([]byte(KeyAddrData), addrData)
		if err != nil {
			log.Println(err)
			return
		}
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
