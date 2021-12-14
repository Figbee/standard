package e

//通用错误
var (
	SUCCESS       = NewError(0, "success")
	ServerError   = NewError(1000, "服务器内部错误")
	InvalidParmas = NewError(1001, "参数错误")
	NotFound      = NewError(1002, "找不到相应资源")

	TokenNotEisit    = NewError(1100, "鉴权失败,找不到对应信息")
	TokenError       = NewError(1101, "鉴权失败,token错误")
	TokenGenerateErr = NewError(1102, "鉴权失败,token生成失败")
	TokenTimeout     = NewError(1103, "鉴权失败,token超时")
	TooManyRequest   = NewError(1104, "鉴权失败,token请求过多")

	AuthFaild    = NewError(1200, "权限鉴定失败")
	AuthNotAllow = NewError(1201, "没有该权限")
)

//模块错误
var (
	ErrorUploadFile = NewError(10001, "上传文件失败")
)
