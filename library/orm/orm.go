package orm

import (
	"fmt"
	"log"

	menuModel "gf-admin/app/model/system/menu"
	menuRole "gf-admin/app/model/system/menu_role"
	roleModel "gf-admin/app/model/system/role"
	userModel "gf-admin/app/model/system/user"
	userRole "gf-admin/app/model/system/user_role"

	"github.com/gogf/gf/frame/g"
	_ "github.com/lib/pq"
	"xorm.io/xorm"
)

var engine *xorm.Engine

func init() {
	var err error
	pgsqlStr := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		g.Cfg().GetString("database.host"),
		g.Cfg().GetString("database.port"),
		g.Cfg().GetString("database.user"),
		g.Cfg().GetString("database.pass"),
		g.Cfg().GetString("database.name"))

	engine, err = xorm.NewEngine("postgres", pgsqlStr)
	if err != nil {
		log.Printf("数据库连接错误:%v \n", err.Error())
		return
	}
	err = engine.Ping()
	if err != nil {
		log.Printf("数据库连接错误:%v \n", err.Error())
		return
	}
	err = engine.Sync2(
		new(userModel.Entity),
		new(menuModel.Entity),
		new(roleModel.Entity),
		new(menuRole.Entity),
		new(userRole.Entity))
	if err != nil {
		log.Printf("同步数据库错误:%v \n", err.Error())
		return
	}
	engine.ShowSQL(true)

}

// Instance 获取数据库实例
func Instance() *xorm.Engine {
	return engine
}
