package ecodes

/*
系统级错误码 (00xxx)

错误码范围	类别	示例代码	说明
00000	成功	00000	操作成功
00100-00199	系统通用错误	00101	系统内部错误
00200-00299	数据库错误	00201	数据库连接失败
00300-00399	缓存错误	00301	Redis连接失败
00400-00499	文件/IO操作错误	00401	文件读写失败
00500-00599	网络通信错误	00501	第三方API调用失败
*/

const (
	// 成功
	ErrCodeSuccess = 0

	// 系统通用错误
	ErrCodeSystemInternalError = 101

	// 数据库错误
	ErrCodeDatabaseError = 201

	// 缓存错误
	ErrCodeCacheError = 301
	// 缓存未启用
	ErrCodeCacheDisabled = 302

	// 文件/IO操作错误
	ErrCodeFileIOError = 401

	// 网络通信错误
	ErrCodeNetworkError = 501
)
