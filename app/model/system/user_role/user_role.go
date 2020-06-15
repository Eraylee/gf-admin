package user_role

//CancelConnectReq 取消关联
type CancelConnectReq struct {
	UserID  int   `p:"userId"`
	RoleIDs []int `p:"roleIds"`
}
