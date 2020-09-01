package util

// 请求状态码
const (
	RECODE_OK 		= 1		// 请求成功 正常
	RECODE_FAIL 	= 0 	// 失败
	RECODE_UNLOGIN 	= -1	// 未登录
)

// 业务逻辑状态码
const (
	RESPMSG_OK = "1"
	RESPMSG_FAIL = "0"

	//  管理员
	RESPMSG_SUCCESSLOGIN = "SUCCESS_LOGIN"
	RESPMSG_FAILURELOGIN = "FAILURE_LOGIN"
	RESPMSG_SUCCESSSESSION = "SUCCESS_SESSION"
	RESPMSG_ERRORSESSION = "ERROR_SESSION"
	RESPMSG_SIGNOUT = "SUCCESS_SIGNOUT"
	RESPMSG_HASNOACCESS = "HAS_NO_ACCESS"
	RESPMSG_ERRORADMINCOUNT = "ERROR_ADMINCOUNT"
)

