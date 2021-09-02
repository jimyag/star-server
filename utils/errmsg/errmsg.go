package errmsg

const (
	SUCCESS = 0
	ERROR   = 500
	// UserAlreadyExist 用户
	UserAlreadyExist = 101
	UserNotExist     = 102
	// TokenError token相关
	TokenError       = 201
	TokenTimeOut     = 202
	TokenTypeError   = 203
	TokenCreateError = 204
	// ParameterConstraintError 参数
	ParameterConstraintError = 301
	// SecretKeyError 部门
	SecretKeyError = 401
	SectorNotExist = 402
	// StudentNotExist 学生
	StudentNotExist = 501
	StudentExist    = 502

	// SectorKeyExist 邀请码
	SectorKeyExist    = 601
	SectorKeyNotExist = 602
	SectorKeyError    = 603
)

var codeMsg = map[int]string{
	SUCCESS:                  "成功",
	ERROR:                    "失败",
	UserAlreadyExist:         "用户已存在",
	UserNotExist:             "用户不存在",
	TokenError:               "Token错误",
	TokenTimeOut:             "Token过期",
	TokenTypeError:           "Token格式错误",
	TokenCreateError:         "Token生成错误",
	ParameterConstraintError: "参数格式错误",
	SecretKeyError:           "邀请码错误",
	SectorNotExist:           "部门不存在",
	StudentNotExist:          "该学生不存在",
	StudentExist:             "该学生已存在",
	SectorKeyExist:           "邀请码已存在",
	SectorKeyNotExist:        "邀请码不存在",
	SectorKeyError:           "邀请码有误",
}

func GetErrMsg(code int) string {
	return codeMsg[code]
}
