package user

import (
	userModel "gf-admin/app/model/system/user"
	"gf-admin/app/model/system/user_role"
	"gf-admin/library/orm"

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
		Nickname: req.Nickname,
		Email:    req.Email,
		Username: req.Username,
		Password: gmd5.MustEncryptString(password),
	}

	if total, err := session.Where("username = ?", user.Username).Count(&user); err != nil {
		return 0, err
	} else if total > 0 {
		return 0, gerror.New("用户名已存在")
	}

	if _, err := session.Insert(&user); err != nil {
		return 0, err
	}

	if len(req.RoleIds) > 0 {
		userRoles := make([]user_role.Entity, 0)
		for i := range req.RoleIds {
			roleID := req.RoleIds[i]
			if roleID != 0 {
				userRole := user_role.Entity{
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
