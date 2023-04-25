package main

import (
	config "board/package/config"
	database "board/package/database/mysql"
	"board/router"
)

func main() {

	// 讀取設定檔
	config.Load()

	// 初始化DB
	database.InitDB()

	// 初始化路由
	router.InitRouter()
}
