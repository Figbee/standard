package e

//正确结果返回封装
type Resp struct {
	Code int         `json:"code"`
	Data interface{} `json:"data"`
	Msg  string      `json:"msg"`
}

type PageInfo struct {
}
