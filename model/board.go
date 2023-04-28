package model

import "time"

type CommentModel struct {
	ID        int64  `gorm:"column:id;primary_key"`
	UserID    int64  `gorm:"column:userID;type:int(11);NOT NULL"`
	NickName  string `gorm:"column:nickname;type:varchar(128);NOT NULL"`
	Content   string `gorm:"column:content;type:text;NOT NULL"`
	CreatedAt time.Time
}

func (CommentModel) TableName() string {
	return "comment"
}

type UserModel struct {
	UserID   int64  `gorm:"column:userID;primary_key"`
	NickName string `gorm:"column:nickName;type:varchar(128);NOT NULL"`
	UserName string `gorm:"column:userName;type:varchar(64);NOT NULL"`
	Password string `gorm:"column:password;type:varchar(128);NOT NULL"`
	CreatedAt time.Time
}

func (UserModel) TableName() string {
	return "users"
}