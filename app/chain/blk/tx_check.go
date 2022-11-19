package blk

import (
	"log"

	"domainchain/proto/chain/blk"

	"github.com/gogo/protobuf/proto"
	kf "github.com/tendermint/tendermint/crypto/secp256k1"
)

func TxBaseCheck(buf []byte) (m *blk.Tx, code uint32) {
	m = &blk.Tx{}
	err := proto.Unmarshal(buf, m)
	if err != nil {
		log.Println(err)
		code = 1
		return
	}
	m2 := *m
	m2.Signature = []byte{}
	buf2, err := proto.Marshal(&m2)
	if err != nil {
		log.Println(err)
		code = 2
		return
	}
	pubKey := kf.PubKey(m.FromPubKey)
	valid := pubKey.VerifySignature(buf2, m.Signature)
	if !valid {
		log.Println(err)
		code = 3
		return
	}
	code = 0
	return
}
