package constance

const (
	RankCount   = 10
	UserOnline  = 2
	UserOffline = 0
)

const (
	Success = 200

	ParamErr = 2000010
	ParamServerErr
	UserNameExist
	UserPassWordFormatErr
	UserPassWordErr

	ParamDBErr = 2000011
)

const (
	SuccessMsg = "ok"

	ParamErrMsg              = "参数错误"
	ParamServerErrMsg        = "服务出错"
	UserNameExistMsg         = "用户名已存在"
	UserPassWordFormatErrMsg = "密码格式错误, 密码格式要求6 - 12位数字，需要由字母 + 数字"
	UserPassWordErrMsg       = "用户密码错误"

	ParamDBErrMsg = "数据库错误"
)
