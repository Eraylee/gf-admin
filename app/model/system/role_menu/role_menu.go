package role_menu

//CancelConnectReq 取消关联
type CancelConnectReq struct {
	RoleID  int   `p:"roleId"`
	MenuIDs []int `p:"menuIds"`
}

// RoleMenus 用户菜单
type RoleMenus []Entity

// ToRoleIDs 转换成角色id数组
func (e *RoleMenus) ToRoleIDs() []int {
	ids := make([]int, 0)
	for _, item := range *e {
		ids = append(ids, item.RoleID)
	}
	return ids
}

// ToMenuIDs 转换成菜单id数组
func (e *RoleMenus) ToMenuIDs() []int {
	ids := make([]int, 0)
	for _, item := range *e {
		ids = append(ids, item.MenuID)
	}
	return ids
}
