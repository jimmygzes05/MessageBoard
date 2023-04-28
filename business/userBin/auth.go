package userBin

import (
	"board/model"
	"board/package/helper"
	myDB "board/repository/myDB"
	"errors"
	"fmt"
)

// 登入
func Login(userName, password string) (user model.UserModel, apiErr error) {
	user, apiErr = myDB.GetUserByUserName(userName)
	if apiErr != nil {
		return
	}

	if !helper.VerifyPassword(password, user.Password) {
		apiErr = errors.New("密碼錯誤")
		return
	}

	return
}

// 註冊
func Register(nickName, userName, password string) (apiErr error)  {
	hashPass, err := helper.EncryptPassword(password)
	if err != nil {
		apiErr = fmt.Errorf("加密失敗,Err:%v", err)
		return
	}

	err = myDB.AddUser(nickName, userName, hashPass)
	if err != nil {
		apiErr = err
		return
	}

	return
}