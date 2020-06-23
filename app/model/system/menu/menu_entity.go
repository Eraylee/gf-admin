package menu

import "time"

// Entity is the golang structure for table user.
type Entity struct {
	ID        int       `json:"id" xorm:"id BIGSERIAL pk"`
	Name      string    `json:"name" xorm:"VARCHAR(30) notnull"`
	Sort      int       `json:"sort" xorm:"INT notnull"`
	Enabled   int       `json:"enabled" xorm:"TINYINT notnull"`
	Visiable  int       `json:"visiable" xorm:"TINYINT notnull"`
	ParentID  int       `json:"parentId" xorm:"parent_id BIGINT"`
	Action    string    `json:"action" xorm:"VARCHAR(16)"`
	Icon      string    `json:"icon" xorm:"VARCHAR(30) notnull"`
	Type      string    `json:"type" xorm:"VARCHAR(30) notnull"`
	Target    string    `json:"target" xorm:"VARCHAR(30) notnull"`
	CreatedAt time.Time `json:"createdAt" xorm:"created notnull"`
	UpdatedAt time.Time `json:"updatedAt" xorm:"updated notnull" `
	DeletedAt time.Time `json:"-" xorm:"deleted"`
}

//TableName 表名
func (Entity) TableName() string {
	return "menu"
}
