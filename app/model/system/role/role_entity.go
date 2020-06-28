package role

import "time"

// Entity is the golang structure for table user.
type Entity struct {
	ID        int       `json:"id" xorm:"id BIGSERIAL pk"`
	Name      string    `json:"name" xorm:"VARCHAR(30) notnull"`
	Code      string    `json:"code" xorm:"VARCHAR(30) notnull"`
	Sort      int       `json:"sort" xorm:"INT notnull"`
	Enabled   int       `json:"enabled" xorm:"TINYINT notnull"`
	Admin     int       `json:"admin" xorm:"TINYINT notnull"`
	CreatedAt time.Time `json:"createdAt" xorm:"created notnull"`
	UpdatedAt time.Time `json:"updatedAt" xorm:"updated notnull" `
	DeletedAt time.Time `json:"-" xorm:"deleted"`
}

//TableName 表名
func (Entity) TableName() string {
	return "role"
}
