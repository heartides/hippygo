package utils

type Response struct {
	Code    int         `json:"code"`
	Msg     string      `json:"msg"`
	Data    interface{} `json:"data,omitempty"`
	ErrDesc string      `json:"err_desc,omitempty"`
}
