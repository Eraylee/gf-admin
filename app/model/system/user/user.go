package user

// CreateUserReq 新增用户资料请求参数
type CreateUserReq struct {
	Username string `p:"username"  v:"required|passport#请输入用户名称|用户名必须是字母开头，只能包含字母、数字和下划线，长度在6~18之间"`
	Phone    string `p:"phone"  v:"required|phone#请输入手机号码|请输入正确的手机号码"`
	Email    string `p:"email"  v:"required|email#请输入邮箱地址|请输入正确的电子邮箱"`
	Nickname string `p:"nickname" v:"required|length:1,30#请输入昵称|昵称长度为:min到:max位"`
	Enabled  bool   `p:"enabled"`
	RoleIds  []int  `p:"roleIds"`
}
