package router

import (
	"github.com/gin-gonic/gin"
)

func setWebViewRouter(r *gin.Engine) {
	const path = "public"

	group := r.Group("/web")

	// 載入前端頁面
	group.GET("/login", func(c *gin.Context) {
		c.File("public/view/login.html")
	})

	group.GET("/register", func(c *gin.Context) {
		c.File("public/view/register.html")
	})

	group.GET("/board", func(c *gin.Context) {
		c.File("public/view/board.html")
	})

	// 載入前端資源
	group.Static("/css", path+"/css")
	group.Static("/js", path+"/js")
}
