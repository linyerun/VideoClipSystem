package resp_msg

const (
	SUCCESS       = 200 // 成功
	SystemERROR   = 500 // 系统内部错误
	InvalidParams = 400 // 无效参数

	ErrorAuthCheckTokenFail    = 601 //token 错误
	ErrorAuthCheckTokenTimeout = 602 //token 过期
	ErrorAuthCodeTimeout       = 603 //验证码过期
	FileLoadError              = 604 // 文件加载失败

	ErrorEmailFormat = 701 // 邮箱格式错误
	ErrorPasswordLen = 702 // 密码过长
	ErrorURLLen      = 703 //URL过长

	ErrorLogin    = 801 //登录失败
	ErrorRegister = 802 //注册失败
	BeRegistered  = 803 //邮箱已经被注册
)
