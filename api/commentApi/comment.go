package commentApi

import (
	"board/business/commentBin"
	"board/common"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

// 取留言回應格式
type GetAllComentResponse struct {
	Err  error     `json:"err"`
	Data []Comment `json:"data"`
}

// 登入顯示所有留言
type Comment struct {
	ID        int64  `json:"id"`
	UserID    int64  `json:"userID"`
	NickName  string `json:"nickname"`
	Content   string `json:"content"`
	CreatedAt string
}

// 取得所有留言
func GetAllComment(c *gin.Context) {
	comments, apiErr := commentBin.GetAllComment()
	if apiErr != nil {
		log.Println(apiErr)
		c.JSON(http.StatusOK, &GetAllComentResponse{
			Err: apiErr,
		})
		return
	}

	data := []Comment{}
	for _, v := range comments {
		comment := Comment{
			ID:        v.ID,
			UserID:    v.UserID,
			NickName:  v.NickName,
			Content:   v.Content,
			CreatedAt: v.CreatedAt.Format("2006-01-02 15:04:05"),
		}
		data = append(data, comment)
	}

	c.JSON(http.StatusOK, &GetAllComentResponse{
		Err:  nil,
		Data: data,
	})
}

// 新增留言請求格式
type AddCommentRequest struct {
	UserID   int64  `json:"userID"`
	NickName string `json:"nickName"`
	Content  string `json:"content"`
}

// 新增留言回應格式
type AddCommentResponse struct {
	Err  error  `json:"err"`
	Data string `json:"data"`
}

// 新增留言
func AddComment(c *gin.Context) {
	req := AddCommentRequest{}
	err := c.ShouldBindJSON(&req)
	if err != nil {
		log.Println(err)
		common.ReturnParseJSON(c, err)
		return
	}

	apiErr := commentBin.AddComment(req.UserID, req.NickName, req.Content)
	if apiErr != nil {
		log.Println(apiErr)
		c.JSON(http.StatusOK, &AddCommentResponse{
			Err:  apiErr,
			Data: "Fail",
		})
		return
	}

	c.JSON(http.StatusOK, &AddCommentResponse{
		Err:  nil,
		Data: "Success",
	})
}

// 修改留言請求格式
type UpdateCommentRequest struct {
	CommentID int64  `json:"commentID"`
	Content   string `json:"content"`
}

// 修改留言回應格式
type UpdateCommentResponse struct {
	Err  error  `json:"err"`
	Data string `json:"data"`
}

// 修改留言
func UpdateComment(c *gin.Context) {
	req := UpdateCommentRequest{}
	err := c.ShouldBindJSON(&req)
	if err != nil {
		log.Println(err)
		common.ReturnParseJSON(c, err)
		return
	}

	apiErr := commentBin.UpdateComment(req.CommentID, req.Content)
	if apiErr != nil {
		log.Println(apiErr)
		c.JSON(http.StatusOK, &UpdateCommentResponse{
			Err:  apiErr,
			Data: "Fail",
		})
		return
	}

	c.JSON(http.StatusOK, &UpdateCommentResponse{
		Err:  nil,
		Data: "Success",
	})
}

// 刪除留言請求格式
type DeleteCommentRequest struct {
	CommentID int64 `json:"commentID"`
}

// 刪除留言回應格式
type DeleteCommentResponse struct {
	Err  error  `json:"err"`
	Data string `json:"data"`
}

// 刪除留言
func DeleteComment(c *gin.Context) {
	req := DeleteCommentRequest{}
	err := c.ShouldBindJSON(&req)
	if err != nil {
		log.Println(err)
		common.ReturnParseJSON(c, err)
		return
	}

	apiErr := commentBin.DeleteComment(req.CommentID)
	if apiErr != nil {
		log.Println(apiErr)
		c.JSON(http.StatusOK, &DeleteCommentResponse{
			Err:  apiErr,
			Data: "Fail",
		})
		return
	}

	c.JSON(http.StatusOK, &DeleteCommentResponse{
		Err:  nil,
		Data: "Success",
	})
}
