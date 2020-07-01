package response

import (
	"net/http"

	"github.com/gogf/gf/net/ghttp"
)

// AppRes 结构体
type AppRes struct {
	r *ghttp.Request
}

// Response 返回值
type Response struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

// Res 响应
func Res(r *ghttp.Request) *AppRes {
	return &AppRes{r}
}

// Response setting gin.JSON
func (a *AppRes) response(code int, message string, data interface{}) {
	_ = a.r.Response.WriteJsonExit(Response{
		Code:    code,
		Message: message,
		Data:    data,
	})
}

//Success 成功
func (a *AppRes) Success(data interface{}) {
	a.response(SUCCESS, "成功", data)
}

// BadRequest 错误请求
func (a *AppRes) BadRequest(message string) {
	// logError(a, BAD_REQUEST, message)
	a.response(BAD_REQUEST, message, nil)
}

// NotFound 404
func (a *AppRes) NotFound() {
	// logError(a, NOT_FOUND, http.StatusText(http.StatusNotFound))
	a.response(NOT_FOUND, http.StatusText(http.StatusNotFound), nil)
}

// InternalServerError 500
func (a *AppRes) InternalServerError(message string) {
	// logError(a, INTERNAL_SERVER_ERROR, message)
	a.response(INTERNAL_SERVER_ERROR, message, nil)
}

// Unauthorized 401
func (a *AppRes) Unauthorized(message string) {
	// logError(a, UNAUTHORIZED, message)
	a.response(UNAUTHORIZED, message, nil)
}

// Forbidden 403
func (a *AppRes) Forbidden(message string) {
	// logError(a, FORBIDDEN, message)
	a.response(FORBIDDEN, message, nil)
}

// func logError(a *AppRes, code int, message string) {
// 	url := a.r.URL
// 	method := a.r.Method
// 	g.Log().Infof("请求失败 url:%s,method:%s,code:%d,error:%s", url, method, code, message)
// }
