package role

import (
	"gf-admin/app/model/base"
	"gf-admin/app/model/system/menu"
)

// QueryRoleReq 分页查询
type QueryRoleReq struct {
	// 姓名
	Name string `p:"name"`
	// 编码
	Code string `p:"code"`
	// 是否超级管理
	Admin int `p:"admin"`
	// 是否可用
	Enabled int `p:"enabled"`

	base.PagingQueryReq
}

// CreateRoleReq 新增角色资料请求参数
type CreateRoleReq struct {
	// 姓名
	Name string `p:"name"  v:"required|length:1,20#请输入角色名称|角色名称长度:min到:max位"`
	// 编码
	Code string `p:"code"  v:"required|length:1,20#请输入角色编码|角色编码长度:min到:max位"`
	// 排序
	Sort int `p:"sort"`
	// 是否超级管理
	Admin int `p:"admin"`
	// 是否可用
	Enabled int `p:"enabled"`
	// 菜单ids
	MenuIDs []int `p:"menuIds"`
}

// UpdateRoleReq 修改角色请求参数
type UpdateRoleReq struct {
	ID int `p:"id"  v:"required#请输入id"`
	// 姓名
	Name string `p:"name"  v:"length:1,20#角色名称长度:min到:max位"`
	// 排序
	Sort int `p:"sort"`
	// 是否超级管理
	Admin int `p:"admin"`
	// 是否可用
	Enabled int `p:"enabled"`
	// 菜单ids
	MenuIDs []int `p:"menuIds"`
}

// Result 返回值
type Result struct {
	Entity
	Menus menu.Menus `json:"menus"`
}

// Results 返回值集合
type Results []Result

// Roles 角色数组
type Roles []Entity

// ToIDs 转换成id数组
func (e *Roles) ToIDs() []int {
	ids := make([]int, 0)
	for _, item := range *e {
		ids = append(ids, item.ID)
	}
	return ids
}

// ToIDs 转换成id数组
func (e *Results) ToIDs() []int {
	ids := make([]int, 0)
	for _, item := range *e {
		ids = append(ids, item.ID)
	}
	return ids
}
