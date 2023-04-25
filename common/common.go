package common

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

// ReturnParseJSON 回傳JSON解析錯誤
func ReturnParseJSON(c *gin.Context, err error) {

	c.JSON(http.StatusOK, APIResponse{
		ErrorText: fmt.Sprintf("json解析錯誤, err:%v", err),
	})
}
