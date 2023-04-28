package commentBin

import (
	"board/model"
	myDB "board/repository/myDB"
)

// 取得所有留言
func GetAllComment() (comments []model.CommentModel, apiErr error) {
	comments, apiErr = myDB.GetAllComment()
	if apiErr != nil {
		return
	}

	return
}

// 新增留言
func AddComment(userID int64, nickName, content string) (apiErr error) {
	apiErr = myDB.AddComment(userID, nickName, content)
	return
}

// 修改留言
func UpdateComment(commentID int64, content string) (apiErr error) {
	apiErr = myDB.UpdateComment(commentID, content)
	return
}

// 刪除留言
func DeleteComment(commentID int64) (apiErr error) {
	apiErr = myDB.DeleteComment(commentID)
	return
}
