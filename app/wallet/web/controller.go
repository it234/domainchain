package web

import (
	"context"
	"encoding/hex"
	"encoding/json"
	"log"
	"net/http"
	"strconv"
	"time"

	"domainchain/app/wallet/store"
	"domainchain/proto/chain/blk"

	"github.com/gin-gonic/gin"
	"github.com/gogo/protobuf/proto"
	"github.com/tendermint/tendermint/crypto"
	kf "github.com/tendermint/tendermint/crypto/secp256k1"
	client "github.com/tendermint/tendermint/rpc/client/http"
)

func GetAddrList(ginc *gin.Context) {
	resp := &GetAddrListResp{}
	resp.Code = 1001
	resp.List = []AddrInfo{}

	store.DbBlk.Mutex.Lock()
	defer store.DbBlk.Mutex.Unlock()
	for k, v := range store.DbBlk.AddrData {
		addrModel := AddrInfo{}
		addrModel.Addr = k
		addrModel.Balance = v.Balance
		resp.List = append(resp.List, addrModel)
	}

	resp.Code = 1000
	ginc.JSON(http.StatusOK, resp)
}

func CreateAddr(ginc *gin.Context) {
	resp := &GetAddrListResp{}
	resp.Code = 1001
	resp.List = []AddrInfo{}

	store.DbBlk.Mutex.Lock()
	defer store.DbBlk.Mutex.Unlock()

	addrModel := store.AddrInfo{}
	priv := kf.GenPrivKey()
	addrModel.PrivateKey = hex.EncodeToString(priv.Bytes())
	addrModel.Addr = priv.PubKey().Address().String()
	addrModel.Balance = 0
	addrModel.TxList = make(map[string]blk.Tx)
	store.DbBlk.AddrData[addrModel.Addr] = addrModel
	addrData, err := json.Marshal(store.DbBlk.AddrData)
	if err != nil {
		log.Println(err)
		ginc.JSON(http.StatusOK, resp)
		return
	}
	store.DbBlk.DB.Set([]byte(store.KeyAddrData), addrData)

	resp.Code = 1000
	ginc.JSON(http.StatusOK, resp)
}

func GetAddrInfo(ginc *gin.Context) {
	resp := &GetAddrInfoResp{}
	resp.Code = 1001
	addr, _ := ginc.GetQuery("addr")
	addrModel := store.DbBlk.AddrData[addr]
	if addrModel.Addr == "" {
		log.Println("地址不存在")
		ginc.JSON(http.StatusOK, resp)
		return
	}
	var txList []string
	for k, _ := range addrModel.TxList {
		txList = append(txList, k)
	}
	resp.TxHash = txList
	resp.Code = 1000
	resp.Addr = addrModel.Addr
	resp.Balance = addrModel.Balance
	ginc.JSON(http.StatusOK, resp)
}

func GetTxByHash(ginc *gin.Context) {
	tx := &Tx{}
	tx.Code = 1001
	addrStr, _ := ginc.GetQuery("addr")
	txHashStr, _ := ginc.GetQuery("hash")
	tx.Hash = txHashStr
	addrModel := store.DbBlk.AddrData[addrStr]
	if addrModel.Addr == "" {
		log.Println("地址不存在")
		ginc.JSON(http.StatusOK, tx)
		return
	}
	txBase := addrModel.TxList[txHashStr]
	if len(txBase.Signature) == 0 {
		log.Println("交易不存在")
		ginc.JSON(http.StatusOK, tx)
		return
	}

	tx.Signature = hex.EncodeToString(txBase.Signature)
	tx.Data = string(txBase.Data)
	fromPub := kf.PubKey(txBase.FromPubKey)
	tx.FromAddr = fromPub.Address().String()
	if txBase.TxType == blk.TxType_ISSUE {
		txM := &blk.TxIssue{}
		err := proto.Unmarshal(txBase.Payload, txM)
		if err != nil {
			log.Println(err)
			ginc.JSON(http.StatusOK, tx)
			return
		}
		tx.Amount = txM.Amount
		tx.ToAddr = crypto.Address(txM.ToAddr).String()
		tx.Data = "发行"
	} else if txBase.TxType == blk.TxType_TRANSFER {
		txM := &blk.TxTransfer{}
		err := proto.Unmarshal(txBase.Payload, txM)
		if err != nil {
			log.Println(err)
			ginc.JSON(http.StatusOK, tx)
			return
		}
		tx.Amount = txM.Amount
		tx.Sequence = txM.Sequence
		tx.ToAddr = crypto.Address(txM.ToAddr).String()
	} else {
		log.Println("tx type err")
		ginc.JSON(http.StatusOK, tx)
		return
	}
	tx.Code = 1000
	ginc.JSON(http.StatusOK, tx)
}

func Issue(ginc *gin.Context) {
	resp := &Tx{}
	resp.Code = 1001
	addrStr, _ := ginc.GetQuery("addr")
	amountStr, _ := ginc.GetQuery("amount")
	amount, _ := strconv.Atoi(amountStr)

	priv001Byt, _ := hex.DecodeString("1888fc3d01efb146f41f3f7d14e742898c62d1a4f09cc20edc55046726a68d7c")
	priv001 := kf.PrivKey(priv001Byt)

	issuerAddrByt, err := hex.DecodeString(addrStr)
	if err != nil {
		log.Println(err)
		ginc.JSON(http.StatusOK, resp)
		return
	}
	issuerAddr := crypto.Address(issuerAddrByt).Bytes()

	txIssue := &blk.TxIssue{}
	txIssue.Amount = uint64(amount)
	txIssue.ToAddr = issuerAddr

	vb, _ := proto.Marshal(txIssue)

	tx := &blk.Tx{}
	tx.TxType = blk.TxType_ISSUE
	tx.FromPubKey = priv001.PubKey().Bytes()
	tx.Signature = []byte{}
	tx.Data = []byte(strconv.Itoa(int(time.Now().Unix())))
	tx.Payload = vb

	vb2, _ := proto.Marshal(tx)

	tx.Signature, err = priv001.Sign(vb2)
	if err != nil {
		log.Println(err)
		ginc.JSON(http.StatusOK, resp)
		return
	}

	cli, err := client.New("http://127.0.0.1:26657")
	if err != nil {
		log.Println(err)
		ginc.JSON(http.StatusOK, resp)
		return
	}

	vb3, _ := proto.Marshal(tx)
	ctx := context.Background()
	respCh, err := cli.BroadcastTxCommit(ctx, vb3)
	if err != nil {
		log.Println(err)
		ginc.JSON(http.StatusOK, resp)
		return
	}
	log.Println(respCh)

	resp.Code = 1000
	ginc.JSON(http.StatusOK, resp)
}

func Transfer(ginc *gin.Context) {
	resp := &Tx{}
	resp.Code = 1001
	fromAddr, _ := ginc.GetQuery("from_addr")
	toAddrStr, _ := ginc.GetQuery("to_addr")
	data, _ := ginc.GetQuery("data")
	amountStr, _ := ginc.GetQuery("amount")
	amount, _ := strconv.Atoi(amountStr)
	log.Println("转账-转出地址:", fromAddr, ",接收地址:", toAddrStr, ",数量:", amount)

	addrModel := store.DbBlk.AddrData[fromAddr]
	if addrModel.Addr == "" {
		log.Println("转出地址不存在")
		ginc.JSON(http.StatusOK, resp)
		return
	}
	priv001Byt, err := hex.DecodeString(addrModel.PrivateKey)
	if err != nil {
		log.Println(err)
		ginc.JSON(http.StatusOK, resp)
		return
	}
	priv001 := kf.PrivKey(priv001Byt)

	toAddrByt, err := hex.DecodeString(toAddrStr)
	if err != nil {
		log.Println(err)
		ginc.JSON(http.StatusOK, resp)
		return
	}
	toAddr := crypto.Address(toAddrByt).Bytes()

	txTran := &blk.TxTransfer{}
	txTran.Amount = uint64(amount)
	txTran.ToAddr = toAddr

	vb, _ := proto.Marshal(txTran)

	tx := &blk.Tx{}
	tx.TxType = blk.TxType_TRANSFER
	tx.FromPubKey = priv001.PubKey().Bytes()
	tx.Signature = []byte{}
	tx.Data = []byte(strconv.Itoa(int(time.Now().Unix())) + "-" + data)
	tx.Payload = vb

	vb2, _ := proto.Marshal(tx)

	tx.Signature, err = priv001.Sign(vb2)
	if err != nil {
		log.Println(err)
		ginc.JSON(http.StatusOK, resp)
		return
	}

	vb3, _ := proto.Marshal(tx)

	cli, err := client.New("http://127.0.0.1:26657")
	if err != nil {
		log.Println(err)
		ginc.JSON(http.StatusOK, resp)
		return
	}

	ctx := context.Background()
	respChain, err := cli.BroadcastTxCommit(ctx, vb3)
	if err != nil {
		log.Println(err)
		ginc.JSON(http.StatusOK, resp)
		return
	}
	log.Println(respChain)

	resp.Code = 1000
	ginc.JSON(http.StatusOK, resp)
}
