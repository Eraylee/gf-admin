package user

import (
	"gf-admin/app/model/base"
	roleModel "gf-admin/app/model/system/role"
	"time"
)

// QueryUserReq 分页查询
type QueryUserReq struct {
	// 用户名
	Username string `p:"username"`
	// 手机号
	Phone string `p:"phone"`
	// 邮箱
	Email string `p:"email"`
	// 昵称
	Nickname string `p:"nickname"`
	// 是否启用
	Enabled int `p:"enabled"`

	base.PagingQueryReq
}

// CreateUserReq 新增用户资料请求参数
// swagger:parameters
// in: body
type CreateUserReq struct {
	// 用户名
	Username string `p:"username"  v:"required|length:1,30#请输入用户名称|用户名长度为:min到:max位"`
	// 手机号
	Phone string `p:"phone"  v:"required|phone#请输入手机号码|请输入正确的手机号码"`
	// 邮箱
	Email string `p:"email"  v:"required|email#请输入邮箱地址|请输入正确的电子邮箱"`
	// 昵称
	Nickname string `p:"nickname" v:"required|length:1,30#请输入昵称|昵称长度为:min到:max位"`
	// 是否启用
	Enabled int `p:"enabled"`
	// 角色
	RoleIDs []int `p:"roleIds"`
}

// UpdateUserReq 修改用户请求参数
type UpdateUserReq struct {
	ID int `p:"id"  v:"required#请输入id"`
	// 手机号
	Phone string `p:"phone"  v:"phone#请输入正确的手机号码"`
	// 邮箱
	Email string `p:"email"  v:"email#请输入正确的电子邮箱"`
	// 昵称
	Nickname string `p:"nickname" v:"length:1,30#昵称长度为:min到:max位"`
	// 是否启用
	Enabled int `p:"enabled"`
	// 角色
	RoleIDs []int `p:"roleIds"`
}

// UpdatePasswordReq 修改用户密码
type UpdatePasswordReq struct {
	// 旧密码
	OldPassword string `p:"oldPassword"  v:"required#请输入原密码"`
	// 新密码
	Password string `p:"password"  v:"required#请输入新密码"`
}

// LoginReq 登录参数
type LoginReq struct {
	// 用户名
	Username string `p:"username"  v:"required#请输入用户名称"`
	// 角色
	Password string `p:"password"  v:"required#请输入密码"`
}

// Result 返回数据
type Result struct {
	ID        int             `json:"id" xorm:"id"` //
	Email     string          `json:"email"`        //
	Phone     string          `json:"phone"`        //
	Nickname  string          `json:"nickname"`     //
	Enabled   int             `json:"enabled"`      //
	Username  string          `json:"username"`     //
	Roles     roleModel.Roles `json:"roles"`
	CreatedAt time.Time       `json:"createdAt"` //
	UpdatedAt time.Time       `json:"updatedAt"` //
}

// Users 用户列表
type Users []Entity

// ToIDs 转换成id数组
func (e *Users) ToIDs() []int {
	ids := make([]int, 0)
	for _, item := range *e {
		ids = append(ids, item.ID)
	}
	return ids
}

// Results 用户返回列表
type Results []Result

// ToIDs 转换成id数组
func (e *Results) ToIDs() []int {
	ids := make([]int, 0)
	for _, item := range *e {
		ids = append(ids, item.ID)
	}
	return ids
}
