package router

import (
	"fmt"
	"board/api/userApi"
	config "board/package/config"
	"net/http"
	"time"

	"github.com/facebookgo/grace/gracehttp"
	"github.com/gin-gonic/gin"
)

func InitRouter() {
	r := gin.Default()

	setRouter(r)

	server := &http.Server{
		Addr:              config.Conf.Server.Port,
		Handler:           r,
		ReadHeaderTimeout: 30 * time.Second,
		WriteTimeout:      30 * time.Second,
	}

	err := gracehttp.Serve(server)
	if err != nil {
		panic(fmt.Sprintf("[ InitAPI 設定API PORT失敗 ] - Err : [ %s ]", err))
	}
}

func setRouter(r *gin.Engine) {

	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, nil)
	})

	// 載入前端畫面
	setWebViewRouter(r)

	setUserRouter(r)
}

func setUserRouter(r *gin.Engine)  {
	group := r.Group("web/api/user")

	group.POST("/login", userApi.Login)
}