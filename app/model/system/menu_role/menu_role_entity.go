package menu_role

// Entity is the golang structure for table user.
type Entity struct {
	MenuID int `json:"menuId" xorm:"menu_id BIGSERIAL pk"`
	RoleID int `json:"roleId" xorm:"role_id BIGSERIAL pk"`
}

//TableName 表名
func (Entity) TableName() string {
	return "menu_role"
}
