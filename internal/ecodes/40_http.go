package ecodes

/*
HTTP相关错误 (40xxx)

错误码范围	HTTP状态码对应	示例代码	说明
40000-40099	400 Bad Request	40001	参数校验失败
40100-40199	401 Unauthorized	40101	未授权/Token失效
40300-40399	403 Forbidden	40301	权限不足
40400-40499	404 Not Found	40401	资源不存在
40500-40599	405 Method Not Allowed	40501	方法不允许
40800-40899	408 Request Timeout	40801	请求超时
40900-40999	409 Conflict	40901	资源冲突
42900-42999	429 Too Many Requests	42901	请求过于频繁
*/

const (
	ErrCodeInvalidParam     = 40001 // 参数格式错误、必填字段缺失、类型不匹配等基础校验失败
	ErrCodeBadRequest       = 40002 // 业务规则校验不通过，如金额不能为负数
	ErrCodeUnauthorized     = 40101
	ErrCodeForbidden        = 40301
	ErrCodeNotFound         = 40401
	ErrCodeMethodNotAllowed = 40501
	ErrCodeRequestTimeout   = 40801
	ErrCodeConflict         = 40901
	ErrCodeTooManyRequests  = 42901
)
