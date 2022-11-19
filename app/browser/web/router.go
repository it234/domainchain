package web

import (
	"github.com/gin-gonic/gin"
)

func Router(app *gin.Engine) {
	app.GET("/get_block_by_hight", GetBlockByHight)
	app.GET("/get_tx_by_hash", GetTxByHash)
	app.GET("/get_new_block_list", GetNewBlockList)
	app.GET("/get_block_page_list", GetBlockPageList)
}
