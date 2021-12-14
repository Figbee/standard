package e

import (
	"fmt"
	"standard/internal/global"

	"github.com/gin-gonic/gin"
)

type Error struct {
	code    int      `json:"code"`
	msg     string   `json:"msg"`
	details []string `json:"details"`
}

var errinfo = map[int]string{}

func NewError(code int, msg string) Error {
	if _, ok := errinfo[code]; ok {
		global.Logger.Panicf("错误码%d 已经存在,请更换一个", code)
	}
	errinfo[code] = msg
	return Error{code: code, msg: msg}
}

func (e *Error) Error() string {
	return fmt.Sprintf("错误码%d,错误信息%s", e.code, e.msg)
}

func (e *Error) Code() int {
	return e.code
}

func (e *Error) Msg() string {
	return e.msg
}

func (e *Error) Msgf(args []interface{}) string {
	return fmt.Sprintf(e.msg, args...)
}

func (e *Error) Details() []string {
	return e.details
}

func (e *Error) withDetails(details ...string) *Error {
	newError := *e
	newError.details = []string{}
	for _, detail := range details {
		newError.details = append(newError.details, detail)
	}
	return &newError
}

func (e *Error) Gin() gin.H {
	return gin.H{
		"code":    e.Code(),
		"msg":     e.Msg(),
		"details": e.Details(),
	}
}
