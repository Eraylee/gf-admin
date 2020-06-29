package role

import "gf-admin/app/model/base"

// QueryRoleReq 分页查询
type QueryRoleReq struct {
	Name    string `p:"name"`
	Code    string `p:"code"`
	Admin   int    `p:"admin"`
	Enabled int    `p:"enabled"`

	base.PagingQueryReq
}

// CreateRoleReq 新增角色资料请求参数
type CreateRoleReq struct {
	Name    string `p:"name"  v:"required|length:1,20#请输入角色名称|角色名称长度:min到:max位"`
	Code    string `p:"code"  v:"required|length:1,20#请输入角色编码|角色编码长度:min到:max位"`
	Sort    int    `p:"sort"`
	Admin   int    `p:"admin"`
	Enabled int    `p:"enabled"`
	MenuIDs []int  `p:"menuIds"`
}

// UpdateRoleReq 修改角色请求参数
type UpdateRoleReq struct {
	ID      int    `p:"id"  v:"required#请输入id"`
	Name    string `p:"name"  v:"length:1,20#角色名称长度:min到:max位"`
	Sort    int    `p:"sort"`
	Admin   int    `p:"admin"`
	Enabled int    `p:"enabled"`
	MenuIDs []int  `p:"menuIds"`
}

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
