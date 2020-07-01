package permission

import (
	permissionLib "gf-admin/library/permission"
	"gf-admin/library/response"

	jwt "github.com/gogf/gf-jwt"
	"github.com/gogf/gf/net/ghttp"
)

// Middleware 中间件
func Middleware(r *ghttp.Request) {
	permission := permissionLib.Instance()
	payload := r.Get("JWT_PAYLOAD")
	user := payload.(jwt.MapClaims)
	codes := user["roleCodes"].([]interface{})

	for _, v := range codes {
		if pass, err := permission.Enforce(v, r.Request.RequestURI, r.Request.Method); err != nil {
			response.Res(r).BadRequest(err.Error())
			return
		} else if !pass {
			response.Res(r).Forbidden("对不起，您没有该接口访问权限，请联系管理员")
			return
		}
	}
	r.Middleware.Next()
}
