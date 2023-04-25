package router

import "github.com/gin-gonic/gin"

func setWebViewRouter(r *gin.Engine)  {
	const path = "public"

	group := r.Group("/web")

	// 載入前端頁面
	group.GET("/", webViewHandler)

	// 載入前端資源
	group.Static("/css", path+"/css")
	group.Static("/js", path+"/js")
}

func webViewHandler(c *gin.Context)  {
	c.File("public/login.html")
}