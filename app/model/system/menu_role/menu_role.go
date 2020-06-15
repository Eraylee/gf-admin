package menu_role

//CancelConnectReq 取消关联
type CancelConnectReq struct {
	MenuID  int   `p:"menuId"`
	RoleIDs []int `p:"roleIds"`
}
