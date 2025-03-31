package ecodes

/*
服务器错误码 (6xxxx)

错误码范围	HTTP状态码对应	示例代码	说明
60000-60099	500 Internal Server Error	60001	服务器内部错误
60100-60199	501 Not Implemented	60101	功能未实现
60200-60299	502 Bad Gateway	60201	网关错误
60300-60399	503 Service Unavailable	60301	服务不可用
60400-60499	504 Gateway Timeout	60401	网关超时
*/

const (
	ErrCodeServerInternalError  = 60001
	ErrCodeServerNotImplemented = 60101
	ErrCodeServerBadGateway     = 60201
	ErrCodeServerUnavailable    = 60301
	ErrCodeServerGatewayTimeout = 60401
)
