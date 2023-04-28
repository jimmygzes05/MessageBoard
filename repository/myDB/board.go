package mydb

import (
	"board/model"
	"errors"
	"fmt"

	"gorm.io/gorm"
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
	if err := db.master.Model(model.CommentModel{}).Order("created_at desc").Find(&comments).Error; err != nil {
		apiErr = err
		return
	}

	return
}

// 新增留言
func AddComment(userID int64, nickName, content string) (apiErr error) {
	newComment := model.CommentModel{
		UserID:   userID,
		NickName: nickName,
		Content:  content,
	}

	err := db.master.Model(model.CommentModel{}).Create(&newComment).Error
	if err != nil {
		apiErr = fmt.Errorf("新增留言失敗，原因:%v", err)
		return
	}

	return
}

// 修改留言
func UpdateComment(commentID int64, content string) (apiErr error) {
	err := db.master.Model(model.CommentModel{}).Where("id = ?", commentID).Update("content", content).Error
	if err != nil {
		apiErr = fmt.Errorf("修改留言失敗，原因:%v", err)
		return
	}

	return
}

func DeleteComment(commentID int64) (apiErr error) {
	comment := model.CommentModel{
		ID: commentID,
	}
	err := db.master.Model(model.CommentModel{}).Delete(&comment).Error
	if err != nil {
		apiErr = fmt.Errorf("刪除留言失敗，原因:%v", err)
		return
	}

	return
}

// 取得使用者資訊
func GetUserByUserName(userName string) (user model.UserModel, apiErr error) {
	err := db.master.Model(model.UserModel{}).Where("userName = ?", userName).Find(&user).Error
	if err != nil {
		apiErr = fmt.Errorf("取使用者資料失敗，原因:%v", err)
		return
	}

	return
}

// 使用者註冊
func AddUser(nickName, userName, password string) (apiErr error) {

	// 先確認使用者帳號是否為唯一值
	var count int64
	err := db.master.Model(model.UserModel{}).Where("userName = ?", userName).Count(&count).Error
	if err != nil {
		apiErr = fmt.Errorf("取使用者資料失敗，原因:%v", err)
		return
	}

	if count > 0 {
		apiErr = errors.New("已有此帳號")
		return
	}

	newUser := model.UserModel{
		NickName: nickName,
		UserName: userName,
		Password: password,
	}

	err = db.master.Model(model.UserModel{}).Create(&newUser).Error
	if err != nil {
		apiErr = fmt.Errorf("新增使用者資料失敗，原因:%v", err)
		return
	}

	return
}
