package role

import "gf-admin/app/model/base"

// QueryRoleReq 分页查询
type QueryRoleReq struct {
	Name    string `p:"name"`
	Code    string `p:"code"`
	Admin   string `p:"admin"`
	Enabled string `p:"enabled"`

	base.PagingQueryReq
}

// CreateRoleReq 新增角色资料请求参数
type CreateRoleReq struct {
	Name    string `p:"name"  v:"required|length:5,30#请输入角色名称|角色名称长度:min到:max位"`
	Code    string `p:"code"  v:"required|length:5,30#请输入角色编码|角色编码长度:min到:max位"`
	Sort    int    `p:"sort"`
	Admin   bool   `p:"admin"`
	Enabled bool   `p:"enabled"`
}

// UpdateRoleReq 修改角色请求参数
type UpdateRoleReq struct {
	ID      int    `p:"id"  v:"required#请输入id"`
	Name    string `p:"name"  v:"length:5,30#请输入角色名称"`
	Sort    int    `p:"sort"`
	Admin   bool   `p:"admin"`
	Enabled bool   `p:"enabled"`
}
