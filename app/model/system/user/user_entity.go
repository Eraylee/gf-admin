package user

import "time"

// Entity is the golang structure for table user.
type Entity struct {
	ID        int       `json:"id" xorm:"id BIGSERIAL pk"`
	Email     string    `json:"email" xorm:"VARCHAR(50)"`
	Phone     string    `json:"phone" xorm:"VARCHAR(11)"`
	Nickname  string    `json:"nickname" xorm:"VARCHAR(50) notnull"`
	Enabled   int       `json:"enabled" xorm:"TINYINT notnull"`
	Username  string    `json:"username" xorm:"VARCHAR(30) notnull"`
	Salt      string    `json:"salt" xorm:"VARCHAR(30) notnull"`
	Password  string    `json:"password" xorm:"VARCHAR(50) notnull"`
	LoginIP   string    `json:"loginIp" xorm:"VARCHAR(50)"`
	LoginDate time.Time `json:"loginAt"`
	CreatedAt time.Time `json:"createdAt" xorm:"created notnull"`
	UpdatedAt time.Time `json:"updatedAt" xorm:"updated notnull" `
	DeletedAt time.Time `json:"deletedAt" xorm:"deleted"`
}

//TableName 表名
func (Entity) TableName() string {
	return "user"
}
