package database

import (
	"fmt"
	"board/repository/myDB"
	conf "board/package/config"

	_ "github.com/go-sql-driver/mysql" // mysql driver
	"github.com/jinzhu/gorm"
)

var MyDB *gorm.DB

// InitDB 初始化DB
func InitDB()  {
	master := connectDB(buildConnStr())
	MyDB = master

	mydb.InitMyDB(MyDB)
}

func buildConnStr() string {
	config := conf.Conf
	return buildDBUrl(config.DB.Account, config.DB.Password, config.DB.Host, config.DB.Port, config.DB.DBName)
}

func buildDBUrl(account, pwd, host, port, dbName string) string {
	return account + ":" + pwd + "@tcp(" + host + ":" + port + ")/" + dbName + "?parseTime=true&loc=Asia%2FTaipei"
}

func connectDB(connStr string) *gorm.DB {
	db, err := gorm.Open("mysql", connStr)
	if err != nil {
		panic(
			fmt.Sprintf("[ connectDB Connect Database Error ] - Err : [ %s ]", err),
		)
	}

	return db
}