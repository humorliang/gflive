package defs

type Err struct {
	Error     string `json:"error"`
	ErrorCode string `json:"error_code"`
}

type ErroResponse struct {
	HttpSC int
	Error  Err
}

//定义错误
var (
	ErrorRequestBodyParseFailed = ErroResponse{HttpSC: 400, Error: Err{Error: "请求参数有误", ErrorCode: "001"}}
	ErrorNotAuthUser            = ErroResponse{HttpSC: 401, Error: Err{Error: "用户未认证", ErrorCode: "002"}}
)
