package browser

import (
	"net/http"

	"domainchain/app/browser/store"
	"domainchain/app/browser/web"
)

func Run() (srv *http.Server) {
	srv = web.Web()
	store.Store()
	return
}
