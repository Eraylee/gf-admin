package role_menu

// Entity is the golang structure for table user.
type Entity struct {
	RoleID int `json:"roleId" xorm:"role_id BIGSERIAL pk"`
	MenuID int `json:"menuId" xorm:"menu_id BIGSERIAL pk"`
}

//TableName 表名
func (Entity) TableName() string {
	return "role_menu"
}
