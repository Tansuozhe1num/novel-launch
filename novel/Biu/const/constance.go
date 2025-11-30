package constance

const (
	RankCount = 10
)

const (
	Success = 200

	ParamErr = 2000010
	ParamServerErr
	UserNameExist
	UserPassWordFormatErr
)

const (
	SuccessMsg = "ok"

	ParamErrMsg              = "参数错误"
	ParamServerErrMsg        = "服务出错"
	UserNameExistMsg         = "用户名已存在"
	UserPassWordFormatErrMsg = "密码格式错误, 密码格式要求6 - 12位数字，需要由字母 + 数字"
)
