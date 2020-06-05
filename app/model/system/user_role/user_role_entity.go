package user_role

// Entity is the golang structure for table user.
type Entity struct {
	UserID int `json:"userId" xorm:"user_id BIGSERIAL pk"`
	RoleID int `json:"roleId" xorm:"role_id BIGSERIAL pk"`
}

//TableName 表名
func (Entity) TableName() string {
	return "user_role"
}
