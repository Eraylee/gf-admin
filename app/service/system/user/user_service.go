package user

import (
	"errors"
	"gf-admin/app/model/base"
	userModel "gf-admin/app/model/system/user"
	userRoleModel "gf-admin/app/model/system/user_role"
	"gf-admin/library/orm"
	"gf-admin/library/paging"

	"github.com/gogf/gf/crypto/gmd5"
	"github.com/gogf/gf/errors/gerror"
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/util/grand"
)

// Create 创建
func Create(req *userModel.CreateUserReq) (int, error) {
	session := orm.Instance().NewSession()
	defer session.Close()
	// add Begin() before any action
	if err := session.Begin(); err != nil {
		return 0, err
	}
	//生成密码
	salt := grand.S(6)
	password := g.Cfg().GetString("app.DefaultPassword") + salt
	user := userModel.Entity{
		Phone:    req.Phone,
		Enabled:  req.Enabled,
		Nickname: req.Nickname,
		Email:    req.Email,
		Username: req.Username,
		Salt:     salt,
		Password: gmd5.MustEncryptString(password),
	}

	if total, err := session.Table(&user).Where("username = ?", user.Username).Count(); err != nil {
		return 0, err
	} else if total > 0 {
		return 0, gerror.New("用户名已存在")
	}

	if _, err := session.Insert(&user); err != nil {
		return 0, err
	}

	if len(req.RoleIDs) > 0 {
		userRoles := make([]userRoleModel.Entity, 0)
		for i := range req.RoleIDs {
			roleID := req.RoleIDs[i]
			if roleID != 0 {
				userRole := userRoleModel.Entity{
					UserID: user.ID,
					RoleID: roleID,
				}
				userRoles = append(userRoles, userRole)
			}
		}

		if len(userRoles) > 0 {
			if _, err := session.Insert(userRoles); err != nil {
				return 0, err
			}
		}

	}

	return user.ID, session.Commit()
}

// Update 修改
func Update(req *userModel.UpdateUserReq) (int, error) {
	session := orm.Instance().NewSession()
	defer session.Close()
	var user userModel.Entity
	if has, err := session.ID(req.ID).Get(&user); err != nil {
		return 0, err
	} else if !has {
		return 0, gerror.New("暂无用户")
	}

	if req.Email != "" {
		user.Email = req.Email
	}
	if req.Nickname != "" {
		user.Nickname = req.Nickname
	}
	if req.Enabled != 0 {
		user.Enabled = req.Enabled
	}
	if req.Phone != "" {
		user.Phone = req.Phone
	}

	if _, err := session.ID(req.ID).Update(&user); err != nil {
		return 0, err
	}

	if len(req.RoleIDs) > 0 {
		userRoles := make([]userRoleModel.Entity, 0)
		for i := range req.RoleIDs {
			roleID := req.RoleIDs[i]
			if roleID != 0 {
				userRole := userRoleModel.Entity{
					UserID: req.ID,
					RoleID: roleID,
				}
				userRoles = append(userRoles, userRole)
			}
		}

		if len(userRoles) > 0 {
			if _, err := session.Where("user_id = ?", req.ID).Delete(new(userRoleModel.Entity)); err != nil {
				return 0, err
			}
			if _, err := session.Insert(userRoles); err != nil {
				return 0, err
			}
		}

	}

	return req.ID, session.Commit()
}

// QueryPage 分页查询
func QueryPage(req *userModel.QueryUserReq) ([]userModel.Res, error) {
	var userEntity userModel.Entity
	//  orderColumn , orderType :=  "created_at" , ""
	db := orm.Instance()
	if req.Username != "" {
		db.Where("username like ?", "%"+req.Username+"%")
	}
	if req.Nickname != "" {
		db.Where("nickname like ?", "%"+req.Nickname+"%")
	}
	if req.Phone != "" {
		db.Where("phone like ?", "%"+req.Phone+"%")
	}
	if req.Email != "" {
		db.Where("email like ?", "%"+req.Email+"%")
	}
	if req.Enabled != 0 {
		db.Where("enabled = ?", req.Enabled)
	}

	total, err := db.Count(&userEntity)

	if err != nil {
		return nil, errors.New("读取行数失败")
	}

	p := paging.Create(req.PageNum, req.PageSize, int(total))

	db.Select("id , email , phone , nickname , username , enabled , created_at , updated_at")

	db.OrderBy(req.OrderColumn + " " + req.OrderType + " ")

	db.Limit(p.PageSize, p.StartNum)

	res := make([]userModel.Res, 0)
	err = db.Table(&userEntity).Desc("created_at").Find(&res)
	return res, err
}

// QueryByID 通过id查询
func QueryByID(ID int) (*userModel.Res, error) {
	var userEntity userModel.Entity
	var res userModel.Res
	db := orm.Instance()
	if _, err := db.Table(&userEntity).ID(ID).Get(&res); err != nil {
		return nil, err
	}
	return &res, nil
}

// ResetPassword 重置密码
func ResetPassword(ID int) (bool, error) {
	var user userModel.Entity
	db := orm.Instance()

	if has, err := db.ID(ID).Get(&user); err != nil {
		return false, err
	} else if !has {
		return false, gerror.New("用户不存在")
	}

	//生成密码
	password := g.Cfg().GetString("app.DefaultPassword") + user.Salt
	user.Password = gmd5.MustEncryptString(password)

	if _, err := db.Update(&user); err != nil {
		return false, err
	}
	return true, nil
}

// UpdatePassword 重置密码
func UpdatePassword(ID int, req *userModel.UpdatePasswordReq) (bool, error) {
	var user userModel.Entity
	db := orm.Instance()

	if has, err := db.ID(ID).Get(&user); err != nil {
		return false, err
	} else if !has {
		return false, gerror.New("用户不存在")
	}

	OldPassword := req.OldPassword + user.Salt

	if user.Password != gmd5.MustEncryptString(OldPassword) {
		return false, gerror.New("原始密码不正确")
	}
	newPassword := req.Password + user.Salt
	user.Password = gmd5.MustEncryptString(newPassword)
	if _, err := db.Update(&user); err != nil {
		return false, err
	}
	return true, nil
}

// Delete 删除
func Delete(req *base.DeleteReq) (int64, error) {
	var user userModel.Entity
	db := orm.Instance()
	res, err := db.In("id", req.Ids).Delete(&user)
	if err != nil {
		return 0, err
	}
	return res, nil
}
