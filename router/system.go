package router

import (
	"gf-admin/app/controller/system/role"
	"gf-admin/app/controller/system/user"
	"gf-admin/library/auth"

	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
)

func init() {
	s := g.Server()
	// s.BindHandler("POST:/login", auth.GfJWTMiddleware.LoginHandler)
	s.Group("/system", func(g *ghttp.RouterGroup) {

		g.POST("/login", auth.GfJWTMiddleware.LoginHandler)
		g.Middleware(auth.MiddlewareAuth)
		g.ALL("/refreshToken", auth.GfJWTMiddleware.RefreshHandler)

		g.Group("/user", func(g *ghttp.RouterGroup) {
			g.POST("/create", user.Create)
			g.POST("/update", user.Update)
			g.POST("/delete", user.Delete)
			g.GET("/queryPage", user.QueryPage)
			g.GET("/queryById", user.QuertByID)
			g.GET("/resetPassword", user.ResetPassword)
			g.POST("/updatePassword", user.UpdatePassword)
		})

		g.Group("/role", func(g *ghttp.RouterGroup) {
			g.POST("/create", role.Create)
			g.POST("/update", role.Update)
			g.POST("/delete", role.Delete)
			g.GET("/queryPage", role.QueryPage)
			g.GET("/queryById", role.QuertByID)
		})
	})

}
