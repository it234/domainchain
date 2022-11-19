package web

import (
	"github.com/gin-gonic/gin"
)

func Router(app *gin.Engine) {
	app.GET("/get_addr_list", GetAddrList)
	app.GET("/get_new_addr", CreateAddr)
	app.GET("/get_addr_info", GetAddrInfo)
	app.GET("/get_tx_by_hash", GetTxByHash)
	app.GET("/issue", Issue)
	app.GET("/transfer", Transfer)
}
