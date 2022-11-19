package web

import (
	"encoding/hex"
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"domainchain/app/browser/store"
	"domainchain/proto/chain/blk"

	"github.com/gin-gonic/gin"
	"github.com/gogo/protobuf/proto"
	"github.com/tendermint/tendermint/crypto"
	kf "github.com/tendermint/tendermint/crypto/secp256k1"
)

func GetBlockByHight(ginc *gin.Context) {
	blkM := &GetBlockByHightResp{}
	blkM.Code = 1001
	hightStr, _ := ginc.GetQuery("hight")
	v, err := store.DbBlk.DB.Get([]byte(store.KeyBlockHightPrefix + hightStr))
	if err != nil {
		log.Println(err)
		ginc.JSON(http.StatusOK, blkM)
		return
	}
	dbBlk := &store.Block{}
	err = json.Unmarshal(v, dbBlk)
	if err != nil {
		log.Println(err)
		ginc.JSON(http.StatusOK, blkM)
		return
	}
	blkM.Code = 1000
	blkM.Hash = dbBlk.Hash
	blkM.Hight = dbBlk.Hight
	blkM.TxHash = dbBlk.TxHash
	ginc.JSON(http.StatusOK, blkM)
}

func GetTxByHash(ginc *gin.Context) {
	tx := &Tx{}
	tx.Code = 1001
	txHashStr, _ := ginc.GetQuery("hash")
	tx.Hash = txHashStr
	v, err := store.DbBlk.DB.Get([]byte(store.KeyTxHashPrefix + txHashStr))
	if err != nil {
		log.Println(err)
		ginc.JSON(http.StatusOK, tx)
		return
	}
	txBase := &blk.Tx{}
	err = proto.Unmarshal(v, txBase)
	if err != nil {
		log.Println(err)
		ginc.JSON(http.StatusOK, tx)
		return
	}
	tx.Signature = hex.EncodeToString(txBase.Signature)
	tx.Data = string(txBase.Data)
	fromPub := kf.PubKey(txBase.FromPubKey)
	tx.FromAddr = fromPub.Address().String()
	if txBase.TxType == blk.TxType_ISSUE {
		txM := &blk.TxIssue{}
		err = proto.Unmarshal(txBase.Payload, txM)
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
		err = proto.Unmarshal(txBase.Payload, txM)
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

func GetNewBlockList(ginc *gin.Context) {
	resp := &GetNewBlockListResp{}
	resp.Code = 1001
	resp.CurrentBlockHight = store.DbBlk.CurrentBlockHight

	for i := store.DbBlk.CurrentBlockHight; i > (store.DbBlk.CurrentBlockHight - 5); i-- {
		v, err := store.DbBlk.DB.Get([]byte(store.KeyBlockHightPrefix + strconv.FormatInt(i, 10)))
		if err != nil {
			log.Println(err)
			ginc.JSON(http.StatusOK, resp)
			return
		}
		dbBlk := &store.Block{}
		err = json.Unmarshal(v, dbBlk)
		if err != nil {
			log.Println(err)
			ginc.JSON(http.StatusOK, resp)
			return
		}
		blkM := Block{}
		blkM.Hash = dbBlk.Hash
		blkM.Hight = dbBlk.Hight
		blkM.TxHash = dbBlk.TxHash
		resp.List = append(resp.List, blkM)
	}
	resp.Code = 1000
	ginc.JSON(http.StatusOK, resp)
}

func GetBlockPageList(ginc *gin.Context) {
	resp := &GetBlockPageListResp{}
	resp.Code = 1001
	resp.List = []Block{}

	req := GetBlockPageListReq{}
	if err := ginc.ShouldBindQuery(&req); err != nil {
		log.Println(err)
		ginc.JSON(400, nil)
		return
	}
	var fromHight, endHight int64
	fromHight = int64(req.PageIndex-1) * int64(req.PageSize)
	fromHight++
	endHight = int64(req.PageIndex) * int64(req.PageSize)
	for i := fromHight; i <= endHight; i++ {
		if i > store.DbBlk.CurrentBlockHight {
			break
		}
		v, err := store.DbBlk.DB.Get([]byte(store.KeyBlockHightPrefix + strconv.FormatInt(i, 10)))
		if err != nil {
			log.Println(err)
			ginc.JSON(http.StatusOK, resp)
			return
		}
		dbBlk := &store.Block{}
		err = json.Unmarshal(v, dbBlk)
		if err != nil {
			log.Println(err)
			ginc.JSON(http.StatusOK, resp)
			return
		}
		blkM := Block{}
		blkM.Hash = dbBlk.Hash
		blkM.Hight = dbBlk.Hight
		blkM.TxHash = dbBlk.TxHash
		resp.List = append(resp.List, blkM)
	}
	resp.TotalSize = store.DbBlk.CurrentBlockHight
	resp.Code = 1000
	ginc.JSON(http.StatusOK, resp)
}
