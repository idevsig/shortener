package ecodes

/*
业务模块错误码 (10xxx)

用户模块 (10xxx)
错误码范围	类别	示例代码	说明
10000-10099	用户通用错误	10001 用户不存在
10100-10199	注册相关错误	10101 用户名已存在
10200-10299	登录相关错误	10201 用户名或密码错误
10300-10399	权限相关错误	10301 无访问权限
						  10302	身份验证失败
						  10303	未授权
10400-10499	资料修改错误	10401 手机号已被使用

订单模块 (11xxx)
错误码范围	类别	示例代码	说明
11000-11099	订单通用错误	11001	订单不存在
11100-11199	创建订单错误	11101	库存不足
11200-11299	支付相关错误	11201	支付失败
11300-11399	订单状态错误	11301	订单已取消

商品模块 (12xxx)
错误码范围	类别	示例代码	说明
12000-12099	商品通用错误	12001	商品不存在
12100-12199	库存相关错误	12101	库存不足
12200-12299	分类相关错误	12201	分类不存在

支付模块 (13xxx)
错误码范围	类别	示例代码	说明
13000-13099	支付通用错误	13001	支付渠道不可用
13100-13199	支付处理错误	13101	支付金额不符
13200-13299	退款相关错误	13201	退款失败
*/

const (
	// 用户模块
	ErrCodeUserNotFound      = 10001
	ErrCodeUserExists        = 10101
	ErrCodeUserLoginFailed   = 10201
	ErrCodeUserPasswordError = 10202

	ErrCodeUserPermissionDenied = 10301
	ErrCodeUserAuthFailed       = 10302

	ErrCodeUserPhoneExists = 10401

	// 订单模块
	ErrCodeOrderNotFound       = 11001
	ErrCodeOrderStockNotEnough = 11101
	ErrCodeOrderPaymentFailed  = 11201
	ErrCodeOrderStatusError    = 11301

	// 商品模块
	ErrCodeProductNotFound         = 12001
	ErrCodeProductStockNotEnough   = 12101
	ErrCodeProductCategoryNotFound = 12201

	// 支付模块
	ErrCodePaymentChannelNotAvailable = 13001
	ErrCodePaymentAmountMismatch      = 13101
	ErrCodeRefundFailed               = 13201
)
