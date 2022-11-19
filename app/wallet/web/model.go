package web

type GetAddrListResp struct {
	Code    int        `json:"code"`
	Message string     `json:"message"`
	List    []AddrInfo `json:"rows"`
}

type AddrInfo struct {
	Addr    string `json:"addr"`
	Balance uint64 `json:"balance"`
}

type GetAddrInfoResp struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	AddrInfo
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
