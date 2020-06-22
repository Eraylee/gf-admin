package role_menu

//CancelConnectReq 取消关联
type CancelConnectReq struct {
	RoleID  int   `p:"roleId"`
	MenuIDs []int `p:"menuIds"`
}
