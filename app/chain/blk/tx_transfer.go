package blk

import (
	"log"

	"domainchain/app/chain/store"
	"domainchain/proto/chain/blk"
	store2 "domainchain/proto/chain/store"

	"github.com/gogo/protobuf/proto"
	abcitypes "github.com/tendermint/tendermint/abci/types"
	kf "github.com/tendermint/tendermint/crypto/secp256k1"
)

func TxTransfer(m *blk.Tx) (rsp abcitypes.ResponseDeliverTx) {
	txs := &blk.TxTransfer{}
	err := proto.Unmarshal(m.Payload, txs)
	if err != nil {
		rsp.Code = 1
		log.Println(err)
		return
	}
	if txs.Amount == 0 {
		rsp.Code = 2
		return
	}
	fromAddr := kf.PubKey(m.FromPubKey).Address().Bytes()
	store.DbBlk.Mutex.Lock()
	defer store.DbBlk.Mutex.Unlock()

	kFrom := &store2.Key{
		Prefix: []byte("addr_amount"),
		Key:    fromAddr,
	}
	kbFrom, err := proto.Marshal(kFrom)
	if err != nil {
		rsp.Code = 3
		log.Println(err)
		return
	}
	vFrom, err := store.DbBlk.DB.Get(kbFrom)
	if err != nil {
		rsp.Code = 4
		log.Println(err)
		return
	}
	addrAmountFrom := &store2.AddrAmount{}
	err = proto.Unmarshal(vFrom, addrAmountFrom)
	if err != nil {
		rsp.Code = 5
		log.Println(err)
		return
	}
	if addrAmountFrom.Sequence != txs.Sequence {
		rsp.Code = 6
		return
	}
	addrAmountFrom.Sequence++
	if txs.Amount > addrAmountFrom.Amount {
		rsp.Code = 7
		return
	}
	addrAmountFrom.Amount -= txs.Amount
	dbb := store.DbBlk.DB.NewBatch()
	defer dbb.Close()
	vbFrom, err := proto.Marshal(addrAmountFrom)
	if err != nil {
		rsp.Code = 8
		log.Println(err)
		return
	}
	err = dbb.Set(kbFrom, vbFrom)
	if err != nil {
		rsp.Code = 9
		log.Println(err)
		return
	}
	kTo := &store2.Key{
		Prefix: []byte("addr_amount"),
		Key:    txs.ToAddr,
	}
	kbTo, err := proto.Marshal(kTo)
	if err != nil {
		rsp.Code = 10
		log.Println(err)
		return
	}
	vTo, err := store.DbBlk.DB.Get(kbTo)
	if err != nil {
		rsp.Code = 11
		log.Println(err)
		return
	}
	addrAmountTo := &store2.AddrAmount{}
	err = proto.Unmarshal(vTo, addrAmountTo)
	if err != nil {
		rsp.Code = 12
		log.Println(err)
		return
	}
	addrAmountTo.Amount += txs.Amount
	vbTo, err := proto.Marshal(addrAmountTo)
	if err != nil {
		rsp.Code = 13
		log.Println(err)
		return
	}
	err = dbb.Set(kbTo, vbTo)
	if err != nil {
		rsp.Code = 14
		log.Println(err)
		return
	}
	err = dbb.Write()
	if err != nil {
		rsp.Code = 15
		log.Println(err)
		return
	}
	return
}
