package userBin

import (
	"errors"
	"board/model"
	"board/package/helper"
	myDB "board/repository/myDB"
)

// 登入
func Login(username, password string) (apiErr error) {
	var user model.UserModel
	user, apiErr = myDB.GetUserByUsername(username)
	if apiErr != nil {
		return
	}

	if !helper.VerifyPassword(password, user.Password) {
		apiErr = errors.New("密碼錯誤")
		return
	}

	return
}