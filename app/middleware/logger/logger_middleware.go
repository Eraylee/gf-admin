package logger

import (
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
	"github.com/gogf/gf/os/gtime"
)

// Middleware 中间件
func Middleware(r *ghttp.Request) {
	// 开始时间
	startTime := gtime.Now()
	r.Middleware.Next()
	// 结束时间
	endTime := gtime.Now()
	// 执行时间
	executionTime := endTime.Sub(startTime)
	// 方法
	method := r.Request.Method
	// 地址
	url := r.Request.RequestURI
	// 状态码
	status := r.Response.Status
	// 请求ip
	ip := r.Request.RemoteAddr

	g.Log().Infof("请求 ip:%s,url:%s,method:%s,status:%d,executionTime:%s", ip, url, method, status, executionTime)
}
