package user_m

import (
	"time"
)

// Entity is the golang structure for table user.
type Entity struct {
	ID        int64     `json:"id" xorm:"id BIGSERIAL pk"`
	Email     string    `json:"email" xorm:"VARCHAR(50)"`
	Phone     string    `json:"phone" xorm:"VARCHAR(11)"`
	Nickname  string    `json:"nickname" xorm:"VARCHAR(50) notnull"`
	Enabled   int       `json:"enabled" xorm:"TINYINT notnull"`
	Admin     int       `json:"admin" xorm:"TINYINT notnull"`
	Username  string    `json:"username" xorm:"VARCHAR(30) notnull"`
	Salt      string    `json:"salt" xorm:"VARCHAR(30) notnull"`
	Password  string    `json:"password" xorm:"VARCHAR(50) notnull"`
	LoginIP   string    `json:"login_ip" xorm:"VARCHAR(50)"`
	LoginDate time.Time `json:"login_at"`
	CreatedAt time.Time `json:"created_at" xorm:"created notnull"`
	UpdatedAt time.Time `json:"updated_at" xorm:"updated notnull" `
	DeletedAt time.Time `json:"deleted_at" xorm:"deleted"`
}

func (Entity) TableName() string {
	return "user"
}
