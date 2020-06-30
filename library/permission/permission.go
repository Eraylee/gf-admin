package permission

import (
	"gf-admin/library/orm"
	"gf-admin/library/response"
	"log"

	"github.com/casbin/casbin/v2"
	xormadapter "github.com/casbin/xorm-adapter/v2"
	jwt "github.com/gogf/gf-jwt"
	"github.com/gogf/gf/net/ghttp"
)

var permission *casbin.Enforcer

func init() {
	db := orm.Instance()

	a, err := xormadapter.NewAdapterByEngine(db)

	if err != nil {
		log.Printf("casbin适配器连接数据库失败err : %s\n", err.Error())
		return
	}
	// a, _ := xormadapter.NewAdapter("mysql", "mysql_username:mysql_password@tcp(127.0.0.1:3306)/") // Your driver and data source.
	permission, err = casbin.NewEnforcer("config/rbac_model.conf", a)
	if err != nil {
		log.Printf("casbin策略加载失败err : %s\n", err.Error())
		return
	}
	// Load the policy from DB.
	permission.LoadPolicy()
}

// Instance 获取实例
func Instance() *casbin.Enforcer {
	return permission
}

// Middleware 中间件
func Middleware(r *ghttp.Request) {
	payload := r.Get("JWT_PAYLOAD")
	user := payload.(jwt.MapClaims)
	codes := user["roleCodes"].([]interface{})

	for _, v := range codes {
		if pass, err := permission.Enforce(v, r.Request.RequestURI, r.Request.Method); err != nil {
			response.Res(r).BadRequest(err.Error())
			return
		} else if !pass {
			response.Res(r).Unauthorized("对不起，您没有该接口访问权限，请联系管理员")
			return
		}
	}
	r.Middleware.Next()
}
