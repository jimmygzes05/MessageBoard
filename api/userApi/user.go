package userApi

import (
	"board/business/userBin"
	"board/common"
	"log"
	"net/http"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

// 登入請求格式
type LoginRequest struct {
	UserName string `json:"userName" binding:"required"` // 使用者名稱
	Password string `json:"password" binding:"required"` // 使用者密碼
}

// 登入回應格式
type LoginResponse struct {
	ErrMsg string `json:"errorMsg"`
	Data   string `json:"data"`
}

// 登入
func Login(c *gin.Context) {
	req := LoginRequest{}
	err := c.ShouldBindJSON(&req)
	if err != nil {
		log.Println(err)
		common.ReturnParseJSON(c, err)
		return
	}

	user, err := userBin.Login(req.UserName, req.Password)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusOK, &LoginResponse{
			ErrMsg: err.Error(),
			Data:   "Fail",
		})
		return
	}

	session := sessions.Default(c)

	session.Set("userName", user.UserName)
	session.Set("userID", user.UserID)
	session.Set("nickName", user.NickName)

	session.Save()

	c.JSON(http.StatusOK, &LoginResponse{
		ErrMsg: "",
		Data:   "Success",
	})
}

// 註冊請求格式
type RegisterRequest struct {
	NickName string `json:"nickName" binding:"required"` // 暱稱
	UserName string `json:"userName" binding:"required"` // 使用者名稱
	Password string `json:"password" binding:"required"` // 使用者密碼
}

// 註冊回應格式
type RegisterResponse struct {
	Err  error  `json:"err"`
	Data string `json:"data"`
}

// 註冊
func Register(c *gin.Context) {
	req := RegisterRequest{}
	err := c.ShouldBindJSON(&req)
	if err != nil {
		log.Println(err)
		common.ReturnParseJSON(c, err)
		return
	}

	apiErr := userBin.Register(req.NickName, req.UserName, req.Password)
	if apiErr != nil {
		log.Println(apiErr)
		c.JSON(http.StatusOK, &RegisterResponse{
			Err:  apiErr,
			Data: "Success",
		})
		return
	}

	c.JSON(http.StatusOK, &RegisterResponse{
		Err:  nil,
		Data: "Success",
	})
}
