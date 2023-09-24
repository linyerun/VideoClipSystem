package resp_msg

var MsgMap = map[int]string{
	SUCCESS:       "ok",
	SystemERROR:   "系统错误",
	InvalidParams: "请求参数错误",

	ErrorAuthCheckTokenFail:    "Token鉴权失败",
	ErrorAuthCheckTokenTimeout: "Token已超时",
	ErrorAuthCodeTimeout:       "验证码过期",
	FileLoadError:              "文件加载失败",

	ErrorEmailFormat: "邮箱格式错误",
	ErrorPasswordLen: "密码长度超出50",
	ErrorURLLen:      "文件URL长度超过500",

	ErrorLogin:    "你在账号或密码错误",
	ErrorRegister: "注册失败",
	BeRegistered:  "邮箱已经被注册",
}

func GetMsgByCode(code int) string {
	return MsgMap[code]
}
