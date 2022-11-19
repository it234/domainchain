package blk

import (
	"encoding/hex"
	"log"

	"domainchain/app/chain/store"
	"domainchain/proto/chain/blk"
	store2 "domainchain/proto/chain/store"

	"github.com/gogo/protobuf/proto"
	abcitypes "github.com/tendermint/tendermint/abci/types"
)

const (
	CanIssuePubkeyHex = "03fda011a6e88c28de4cc1ce6360ae11c7087a2ca83f853b07f6a33af25239dacf"
)

func TxIssue(m *blk.Tx) (rsp abcitypes.ResponseDeliverTx) {
	fromPubKeyHex := hex.EncodeToString(m.FromPubKey)
	if fromPubKeyHex != CanIssuePubkeyHex {
		rsp.Code = 1
		return
	}
	txs := &blk.TxIssue{}
	err := proto.Unmarshal(m.Payload, txs)
	if err != nil {
		rsp.Code = 2
		log.Println(err)
		return
	}
	if txs.Amount == 0 {
		rsp.Code = 3
		return
	}
	store.DbBlk.Mutex.Lock()
	defer store.DbBlk.Mutex.Unlock()
	k := &store2.Key{
		Prefix: []byte("addr_amount"),
		Key:    txs.ToAddr,
	}
	kb, err := proto.Marshal(k)
	if err != nil {
		rsp.Code = 4
		log.Println(err)
		return
	}
	v, err := store.DbBlk.DB.Get(kb)
	if err != nil {
		rsp.Code = 5
		log.Println(err)
		return
	}
	addrAmount := &store2.AddrAmount{}
	err = proto.Unmarshal(v, addrAmount)
	if err != nil {
		rsp.Code = 6
		log.Println(err)
		return
	}
	addrAmount.Amount += txs.Amount
	vb, err := proto.Marshal(addrAmount)
	if err != nil {
		rsp.Code = 7
		log.Println(err)
		return
	}
	err = store.DbBlk.DB.Set(kb, vb)
	if err != nil {
		rsp.Code = 8
		log.Println(err)
		return
	}
	return
}
