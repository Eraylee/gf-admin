package user

import (
	"errors"
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

	if len(req.RoleIds) > 0 {
		userRoles := make([]userRoleModel.Entity, 0)
		for i := range req.RoleIds {
			roleID := req.RoleIds[i]
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
	user := userModel.Entity{
		Phone:    req.Phone,
		Enabled:  req.Enabled,
		Nickname: req.Nickname,
		Email:    req.Email,
	}

	if _, err := session.ID(req.ID).Update(&user); err != nil {
		return 0, err
	}

	if len(req.RoleIds) > 0 {
		userRoles := make([]userRoleModel.Entity, 0)
		for i := range req.RoleIds {
			roleID := req.RoleIds[i]
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
	if req.Enabled != "" {
		db.Where("email = ?", req.Enabled)
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

// //Login 登录
// func Login(req *userModel.LoginReq) (string, error) {
// 	var userEntity userModel.Entity
// 	db := orm.Instance()
// 	if err := db.Table(&userEntity).Where("username = ?", req.Username).Find(userEntity); err != nil {
// 		return " ", err
// 	} else if userEntity == nil {
// 		return " ", gerror.New("用户名或密码错误")
// 	}
// 	password := userEntity.Password + userEntity.Salt
// 	if gmd5.MustEncryptString(password) != req.Password {
// 		return " ", gerror.New("用户名或密码错误")
// 	}

// }
