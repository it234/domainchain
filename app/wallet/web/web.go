package web

import (
	"net/http"
	"path/filepath"
	"time"

	"domainchain/pkg"

	"github.com/gin-gonic/gin"
)

func Web() (srv *http.Server) {
	gin.SetMode(gin.ReleaseMode)
	app := gin.New()
	srv = &http.Server{
		Addr:         ":9102",
		Handler:      app,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  15 * time.Second,
	}

	path := pkg.GetRootDir() + "wallet/web"
	indexC := func(c *gin.Context) {
		fpath := filepath.Join(path, "/static/html/p_addr_list.html")
		http.ServeFile(c.Writer, c.Request, fpath)
		c.Abort()
	}
	app.GET("/", indexC)
	app.StaticFS("/static", http.Dir(path+"/static"))
	Router(app)

	go srv.ListenAndServe()
	return
}
