package router

import (
	"board/api/commentApi"
	"board/api/userApi"
	config "board/package/config"
	"fmt"
	"net/http"
	"time"

	"github.com/facebookgo/grace/gracehttp"
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
)

func InitRouter() {
	r := gin.Default()

	store := cookie.NewStore([]byte("secret"))
	r.Use(sessions.Sessions("mysession", store))

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

	r.GET("/session", func(c *gin.Context) {
		session := sessions.Default(c)
		userName := session.Get("userName")
		nickName := session.Get("nickName")
		userID := session.Get("userID")
		c.JSON(http.StatusOK, gin.H{
			"userName": userName,
			"userID":   userID,
			"nickName": nickName,
		})
	})

	// 載入前端畫面
	setWebViewRouter(r)

	// 
	setUserRouter(r)

	setCommentRouter(r)
}

func setUserRouter(r *gin.Engine) {
	group := r.Group("api/user")

	group.POST("/login", userApi.Login)
	group.POST("/register", userApi.Register)
}

func setCommentRouter(r *gin.Engine) {
	group := r.Group("api/comment")

	group.GET("", commentApi.GetAllComment)
	group.POST("", commentApi.AddComment)
	group.PUT("", commentApi.UpdateComment)
	group.DELETE("", commentApi.DeleteComment)
}
