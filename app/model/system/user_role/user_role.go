package user_role

//CancelConnectReq 取消关联
type CancelConnectReq struct {
	UserID  int   `p:"userId"`
	RoleIDs []int `p:"roleIds"`
}

// UserRoles 用户菜单
type UserRoles []Entity

// ToRoleIDs 转换成角色id数组
func (e *UserRoles) ToRoleIDs() []int {
	ids := make([]int, 0)
	for _, item := range *e {
		ids = append(ids, item.RoleID)
	}
	return ids
}

// ToUserIDs 转换成菜单id数组
func (e *UserRoles) ToUserIDs() []int {
	ids := make([]int, 0)
	for _, item := range *e {
		ids = append(ids, item.UserID)
	}
	return ids
}
