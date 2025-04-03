package ecodes

/*
// 分类结构
// 采用 AABBB 格式：
// AA: 模块/大类代码 (01-99)
// BBB: 具体错误代码 (000-999)
//
// 70xxx-89xxx: 为未来业务扩展预留
// 90xxx-99xxx: 为特殊场景预留
*/
var errCodeMessages = map[int]string{
	ErrCodeSuccess:             "成功",
	ErrCodeSystemInternalError: "系统内部错误",
	ErrCodeDatabaseError:       "数据库错误",
	ErrCodeCacheError:          "缓存错误",
	ErrCodeCacheDisabled:       "缓存未启用",
	ErrCodeFileIOError:         "文件/IO操作错误",
	ErrCodeNetworkError:        "网络通信错误",

	ErrCodeUserNotFound:         "用户不存在",
	ErrCodeUserExists:           "用户已存在",
	ErrCodeUserLoginFailed:      "用户登录失败",
	ErrCodeUserPermissionDenied: "用户权限不足",
	ErrCodeUserPhoneExists:      "手机号已存在",

	ErrCodeOrderNotFound:       "订单不存在",
	ErrCodeOrderStockNotEnough: "库存不足",
	ErrCodeOrderPaymentFailed:  "支付失败",
	ErrCodeOrderStatusError:    "订单状态错误",

	ErrCodeProductNotFound:         "商品不存在",
	ErrCodeProductStockNotEnough:   "库存不足",
	ErrCodeProductCategoryNotFound: "分类不存在",

	ErrCodePaymentChannelNotAvailable: "支付渠道不可用",
	ErrCodePaymentAmountMismatch:      "支付金额不符",
	ErrCodeRefundFailed:               "退款失败",

	ErrCodeInvalidParam:     "参数错误",
	ErrCodeBadRequest:       "请求失败",
	ErrCodeUnauthorized:     "未授权",
	ErrCodeForbidden:        "禁止访问",
	ErrCodeNotFound:         "数据不存在",
	ErrCodeMethodNotAllowed: "方法不允许",
	ErrCodeRequestTimeout:   "请求超时",
	ErrCodeConflict:         "数据已存在",
	ErrCodeTooManyRequests:  "请求过多",

	ErrCodeSMSFailed:            "短信发送失败",
	ErrCodeEmailFailed:          "邮件发送失败",
	ErrCodeFileUploadFailed:     "文件上传失败",
	ErrCodeGeocodingFailed:      "地理编码失败",
	ErrCodePaymentGatewayFailed: "支付网关连接失败",

	ErrCodeServerInternalError:  "服务器内部错误",
	ErrCodeServerNotImplemented: "功能未实现",
	ErrCodeServerBadGateway:     "网关错误",
	ErrCodeServerUnavailable:    "服务不可用",
	ErrCodeServerGatewayTimeout: "网关超时",
}

// GetErrCodeMessage 获取错误码对应的错误信息
func GetErrCodeMessage(code int) string {
	if message, ok := errCodeMessages[code]; ok {
		return message
	}
	return "未知错误"
}
