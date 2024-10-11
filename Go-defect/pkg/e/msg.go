package e

var MsgFlags = map[int]string{
	SUCCESS:        "成功响应",
	ERROR:          "服务器内部错误",
	INVALID_PARAMS: "请求参数错误",

	ERROR_AUTH_CHECK_TOKEN_FAIL:    "Token鉴权失败",
	ERROR_AUTH_CHECK_TOKEN_TIMEOUT: "Token已超时",
	ERROR_AUTH_TOKEN:               "Token生成失败",
	ERROR_AUTH:                     "登录失败，用户名或密码错误",

	ERROR_PARSE_GET_ROLE: "登录失败,Token解析获取角色信息失败",
	ERROR_AUTH_PARSE:    "Token解析失败",

	ERROR_AUTH_CHECK_ROLE_FAIL: "角色权限不足",

	ERROR_ADD_USER:    "添加新用户失败",
	ERROR_USER_EXIST:  "用户名已存在",
}

func GetMsg(code int) string {
	msg, ok := MsgFlags[code]
	if ok {
		return msg
	}

	return MsgFlags[ERROR]
}
