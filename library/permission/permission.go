package permission

import (
	"gf-admin/library/orm"

	"github.com/casbin/casbin/v2"
	xormadapter "github.com/casbin/xorm-adapter/v2"
	"github.com/gogf/gf/os/glog"
)

var permission *casbin.Enforcer

func init() {
	db := orm.Instance()

	a, err := xormadapter.NewAdapterByEngine(db)

	if err != nil {
		glog.Error("casbin适配器连接数据库失败err : %s\n", err.Error())
		return
	}
	// a, _ := xormadapter.NewAdapter("mysql", "mysql_username:mysql_password@tcp(127.0.0.1:3306)/") // Your driver and data source.
	permission, err = casbin.NewEnforcer("config/rbac_model.conf", a)
	if err != nil {
		glog.Error("casbin策略加载失败err : %s\n", err.Error())
		return
	}
	// Load the policy from DB.
	permission.LoadPolicy()
}

// Instance 获取实例
func Instance() *casbin.Enforcer {
	return permission
}
