package router

import (
	"gf-admin/app/controller/system/user"
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
)

func init() {
	s := g.Server()
	s.Group("/system", func(group *ghttp.RouterGroup) {
		//group.Middleware(auth.Auth)
		group.Group("/user", func(group *ghttp.RouterGroup) {

			group.POST("/create", user_c.Create)

		})

	})
}
