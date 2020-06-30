package router

import (
	"gf-admin/app/controller/system/menu"
	"gf-admin/app/controller/system/role"
	"gf-admin/app/controller/system/user"
	"gf-admin/library/auth"
	"gf-admin/library/permission"

	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
)

func init() {
	s := g.Server()
	// s.BindHandler("POST:/login", auth.GfJWTMiddleware.LoginHandler)
	s.Group("/system", func(g *ghttp.RouterGroup) {

		g.POST("/login", auth.GfJWTMiddleware.LoginHandler)
		g.Middleware(auth.Middleware)
		g.Middleware(permission.Middleware)
		g.ALL("/refreshToken", auth.GfJWTMiddleware.RefreshHandler)

		g.Group("/user", func(g *ghttp.RouterGroup) {

			g.GET("/queryPage", user.QueryPage)
			g.GET("/queryById", user.QueryByID)
			g.GET("/resetPassword", user.ResetPassword)
			g.POST("/create", user.Create)
			g.POST("/update", user.Update)
			g.POST("/delete", user.Delete)
			g.POST("/updatePassword", user.UpdatePassword)
		})

		g.Group("/role", func(g *ghttp.RouterGroup) {
			g.GET("/queryPage", role.QueryPage)
			g.GET("/queryById", role.QueryByID)
			g.POST("/create", role.Create)
			g.POST("/update", role.Update)
			g.POST("/delete", role.Delete)
			g.POST("/cancelUserConnect", role.CancelUserConnect)
		})

		g.Group("/menu", func(g *ghttp.RouterGroup) {
			g.GET("/queryPage", menu.QueryPage)
			g.GET("/queryTree", menu.QueryTree)
			g.GET("/queryById", menu.QueryByID)
			g.POST("/create", menu.Create)
			g.POST("/update", menu.Update)
			g.POST("/delete", menu.Delete)
			g.POST("/cancelMenuConnect", menu.CancelMenuConnect)
		})
	})
}
