package user

import (
	"time"
)

// Entity is the golang structure for table user.
type Entity struct {
	ID        int64     `json:"id" xorm:"SERIAL comment('用户ID')"`
	Email     string    `json:"email" xorm:"email VARCHAR(50) comment('邮箱')"`
	Phone     string    `json:"phone" xorm:"phone VARCHAR(11) comment('手机号码')"`
	Nickname  string    `json:"nickname" xorm:"nickname VARCHAR(50) comment('昵称')"`
	Enabled   bool      `json:"enabled" xorm:"enabled BOOL comment('启用状态')"`
	Username  string    `json:"username" xorm:"username VARCHAR(30) comment('用户名')"`
	Password  string    `json:"password" xorm:"password VARCHAR(50) comment('密码')"`
	LoginIP   string    `json:"login_ip" xorm:"login_ip VARCHAR(50) comment('最后登录ip')"`
	LoginDate time.Time `json:"login_at" xorm:"login_at comment('最后登录时间')"`
	CreatedAt time.Time `json:"created_at" xorm:"created"`
	UpdatedAt time.Time `json:"updated_at" xorm:"updated"`
	DeletedAt time.Time `json:"deleted_at" xorm:"deleted"`
}
