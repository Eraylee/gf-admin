package orm

import (
	"fmt"
	"log"

	"github.com/gogf/gf/frame/g"
	_ "github.com/lib/pq"
	"xorm.io/xorm"

	userModel "gf-admin/app/model/system/user"
)

var engine *xorm.Engine

func init() {
	var err error
	pgsql := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		g.Cfg().GetString("database.host"),
		g.Cfg().GetString("database.port"),
		g.Cfg().GetString("database.user"),
		g.Cfg().GetString("database.pass"),
		g.Cfg().GetString("database.name"))

	engine, err = xorm.NewEngine("postgres", pgsql)
	if err != nil {
		log.Printf("数据库连接错误:%v \n", err.Error())
		return
	}
	err = engine.Ping()
	if err != nil {
		log.Printf("数据库连接错误:%v \n", err.Error())
		return
	}
	engine.Sync2(new(userModel.Entity))
	engine.ShowSQL(true)
}

// Instance 获取数据库实例
func Instance() (engine *xorm.Engine) {
	return engine
}
