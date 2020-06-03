package user_s

import (
	"gf-admin/app/model/system/user"
	"gf-admin/library/orm"
)

// Create 创建
func Create(req *user_m.CreateUserReq) (int64, error) {
	user := user_m.Entity{
		Nickname: req.Nickname,
		Email:    req.Email,
		Username: req.Username,
	}
	db := orm.Instance()
	_, err := db.Table(&user).Insert(&user)
	return user.ID, err
}
