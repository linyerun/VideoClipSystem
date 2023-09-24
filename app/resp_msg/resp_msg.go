package resp_msg

type RespMsg struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
	Data any    `json:"data,omitempty"`
}

func NewRespMsg(code int, data any) *RespMsg {
	return &RespMsg{
		Code: code,
		Msg:  GetMsgByCode(code),
		Data: data,
	}
}
