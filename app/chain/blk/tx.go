package blk

import (
	"domainchain/proto/chain/blk"

	abcitypes "github.com/tendermint/tendermint/abci/types"
)

func TxHandle(m *blk.Tx) (rsp abcitypes.ResponseDeliverTx) {
	switch m.TxType {
	case blk.TxType_ISSUE:
		rsp = TxIssue(m)
	case blk.TxType_TRANSFER:
		rsp = TxTransfer(m)
	}
	return
}
