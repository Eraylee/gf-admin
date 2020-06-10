package main

import (
	_ "gf-admin/boot"
	_ "gf-admin/router"

	"github.com/gogf/gf-swagger/swagger"
	"github.com/gogf/gf/frame/g"
)

// @title Golang Gin API
// @version 1.0
// @description ERAYLEE'S web api

// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization
func main() {
	s := g.Server()
	s.Plugin(&swagger.Swagger{})
	s.Run()
}
