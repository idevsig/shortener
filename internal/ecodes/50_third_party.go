package ecodes

/*
第三方服务错误码 (50xxx)

错误码范围	类别	示例代码	说明
50000-50099	短信服务错误	50001	短信发送失败
50100-50199	邮件服务错误	50101	邮件发送失败
50200-50299	云存储错误	50201	文件上传失败
50300-50399	地图服务错误	50301	地理编码失败
50400-50499	支付网关错误	50401	支付网关连接失败
*/

const (
	ErrCodeSMSFailed            = 50001
	ErrCodeEmailFailed          = 50101
	ErrCodeFileUploadFailed     = 50201
	ErrCodeGeocodingFailed      = 50301
	ErrCodePaymentGatewayFailed = 50401
)
