package main

import (
	_ "gf-admin/boot"
	_ "gf-admin/router"

	"github.com/gogf/gf/frame/g"
)

func main() {
	g.Server().Run()
}
