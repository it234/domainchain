package web

type GetBlockByHightResp struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Block
}

type Block struct {
	Hight  int64    `json:"hight"`
	Hash   string   `json:"hash"`
	TxHash []string `json:"tx_hash"`
}

type Tx struct {
	Code      int    `json:"code"`
	Message   string `json:"message"`
	Hash      string `json:"hash"`
	FromAddr  string `json:"from_addr"`
	ToAddr    string `json:"to_addr"`
	Amount    uint64 `json:"amount"`
	Sequence  uint64 `json:"sequence"`
	Data      string `json:"data"`
	Signature string `json:"signature"`
}

type GetNewBlockListResp struct {
	Code              int     `json:"code"`
	Message           string  `json:"message"`
	CurrentBlockHight int64   `json:"current_block_hight"`
	List              []Block `json:"list"`
}

type GetBlockPageListReq struct {
	PageSize  int `json:"rows" form:"rows" uri:"rows"`
	PageIndex int `json:"page" form:"page" uri:"page"`
}

type GetBlockPageListResp struct {
	Code      int     `json:"code"`
	Message   string  `json:"message"`
	TotalSize int64   `json:"total"`
	List      []Block `json:"rows"`
}
