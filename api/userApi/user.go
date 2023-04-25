package userApi

import (
	"board/business/userBin"
	"board/common"
	"net/http"

	"github.com/gin-gonic/gin"
)

// LoginRequest 登入
type LoginRequest struct {
	Username string `json:"username" binding:"required"` // 使用者名稱
	Password string `json:"password" binding:"required"` // 使用者密碼
}

// LoginResponse 登入
type LoginResponse struct {
	Err  error  `json:"err"`
	Data string `json:"data"`
}

func Login(c *gin.Context) {
	req := LoginRequest{}
	err := c.ShouldBindJSON(&req)
	if err != nil {
		common.ReturnParseJSON(c, err)
	}

	err = userBin.Login(req.Username, req.Password)
	if err != nil {
		c.JSON(http.StatusOK, &LoginResponse{
			Err:  err,
			Data: "Fail",
		})
	}

	c.JSON(http.StatusOK, &LoginResponse{
		Err:  nil,
		Data: "Success",
	})
}
