package router

import (
	"gf-admin/app/controller/system/user"
	"gf-admin/library/auth"

	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
)

func init() {
	s := g.Server()
	// s.BindHandler("POST:/login", auth.GfJWTMiddleware.LoginHandler)
	s.Group("/system", func(group *ghttp.RouterGroup) {
		group.POST("/login", auth.GfJWTMiddleware.LoginHandler)
		group.Middleware(auth.MiddlewareAuth)
		group.ALL("/refreshToken", auth.GfJWTMiddleware.RefreshHandler)
		group.Group("/user", func(group *ghttp.RouterGroup) {

			group.POST("/create", user.Create)
			group.POST("/update", user.Update)
			group.GET("/queryPage", user.QueryPage)

		})

	})
}
