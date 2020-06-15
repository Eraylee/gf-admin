package user

import (
	"gf-admin/app/model/base"
	"time"
)

// QueryUserReq 分页查询
type QueryUserReq struct {
	Username string `p:"username"`
	Phone    string `p:"phone"`
	Email    string `p:"email"`
	Nickname string `p:"nickname"`
	Enabled  int    `p:"enabled"`

	base.PagingQueryReq
}

// CreateUserReq 新增用户资料请求参数
type CreateUserReq struct {
	Username string `p:"username"  v:"required|passport#请输入用户名称|用户名必须是字母开头，只能包含字母、数字和下划线，长度在6~18之间"`
	Phone    string `p:"phone"  v:"required|phone#请输入手机号码|请输入正确的手机号码"`
	Email    string `p:"email"  v:"required|email#请输入邮箱地址|请输入正确的电子邮箱"`
	Nickname string `p:"nickname" v:"required|length:1,30#请输入昵称|昵称长度为:min到:max位"`
	Enabled  int    `p:"enabled"`
	RoleIDs  []int  `p:"roleIds"`
}

// UpdateUserReq 修改用户请求参数
type UpdateUserReq struct {
	ID       int    `p:"id"  v:"required#请输入id"`
	Phone    string `p:"phone"  v:"phone#请输入正确的手机号码"`
	Email    string `p:"email"  v:"email#请输入正确的电子邮箱"`
	Nickname string `p:"nickname" v:"length:1,30#昵称长度为:min到:max位"`
	Enabled  int    `p:"enabled"`
	RoleIDs  []int  `p:"roleIds"`
}

// UpdatePasswordReq 修改用户密码
type UpdatePasswordReq struct {
	OldPassword string `p:"oldPassword"  v:"required#请输入原密码"`
	Password    string `p:"password"  v:"required#请输入新密码"`
}

// LoginReq 登录参数
type LoginReq struct {
	Username string `p:"username"  v:"required#请输入用户名称"`
	Password string `p:"password"  v:"required#请输入密码"`
}

// Res 返回数据
type Res struct {
	ID        int       `json:"id" xorm:"id"`                //
	Email     string    `json:"email" xorm:"email"`          //
	Phone     string    `json:"phone" xorm:"phone"`          //
	Nickname  string    `json:"nickname" xorm:"nickname"`    //
	Enabled   int       `json:"enabled" xorm:"enabled"`      //
	Username  string    `json:"username" xorm:"username"`    //
	CreatedAt time.Time `json:"createdAt" xorm:"created_at"` //
	UpdatedAt time.Time `json:"updatedAt" xorm:"updated_at"` //
}
