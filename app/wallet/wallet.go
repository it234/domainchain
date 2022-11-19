package wallet

import (
	"net/http"

	"domainchain/app/wallet/store"
	"domainchain/app/wallet/web"
)

func Run() (srv *http.Server) {
	srv = web.Web()
	store.Store()
	return
}
