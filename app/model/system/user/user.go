package user

import "gf-admin/app/model/base"

// CreateUserReq 新增用户资料请求参数
type CreateUserReq struct {
	Username string `p:"username"  v:"required|passport#请输入用户名称|用户名必须是字母开头，只能包含字母、数字和下划线，长度在6~18之间"`
	Phone    string `p:"phone"  v:"required|phone#请输入手机号码|请输入正确的手机号码"`
	Email    string `p:"email"  v:"required|email#请输入邮箱地址|请输入正确的电子邮箱"`
	Nickname string `p:"nickname" v:"required|length:1,30#请输入昵称|昵称长度为:min到:max位"`
	Enabled  bool   `p:"enabled"`
	RoleIds  []int  `p:"roleIds"`
}

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
