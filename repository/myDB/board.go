package mydb

import (
	"board/model"

	"github.com/jinzhu/gorm"
)

type repository struct {
	master *gorm.DB
}

var db *repository

// InitMyDB 初始化DB
func InitMyDB(master *gorm.DB) {
	db = &repository{
		master: master,
	}
}

// 取得所有留言
func GetAllComment() (comments []model.CommentModel, apiErr error) {
	if err := db.master.Model(model.CommentModel{}).Find(&comments).Error; err != nil {
		apiErr = err
		return
	}

	return
}

// GetUserByUsername 取得使用者資訊
func GetUserByUsername(username string) (user model.UserModel, apiErr error) {
	err := db.master.Model(model.UserModel{}).Where("username = ?", username).Find(&user).Error
	if err != nil {
		apiErr = err
		return
	}

	return
}
