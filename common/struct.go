package common

// APIResponse  api 回傳格式
type APIResponse struct {
	ErrorText string      `json:"error_text"`
	Result    interface{} `json:"result"`
}