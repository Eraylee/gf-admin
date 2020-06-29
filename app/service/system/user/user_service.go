package user

import (
	"errors"
	"gf-admin/app/model/base"
	roleModel "gf-admin/app/model/system/role"
	userModel "gf-admin/app/model/system/user"
	userRoleModel "gf-admin/app/model/system/user_role"
	roleService "gf-admin/app/service/system/role"
	"gf-admin/library/orm"
	"gf-admin/library/paging"

	"github.com/gogf/gf/container/gmap"
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
		var userRoles userRoleModel.UserRoles
		for _, roleID := range req.RoleIDs {
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
		var userRoles userRoleModel.UserRoles
		for _, roleID := range req.RoleIDs {
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
func QueryPage(req *userModel.QueryUserReq) (*base.PagingRes, error) {
	// 分页查询用户信息
	users, p, err := queryUsers(req)
	if err != nil {
		return nil, err
	}
	//获取用户对应的关系
	userIDs := users.ToIDs()
	userRoles, err := roleService.QueryUserRoles(userIDs)
	if err != nil {
		return nil, err
	}
	//获取所有关联角色的详情
	roleIDs := userRoles.ToRoleIDs()
	roles, err := roleService.QueryRoles(roleIDs)
	if err != nil {
		return nil, err
	}

	res := buildResults(users, roles, userRoles)

	return &base.PagingRes{
		Data:   res,
		Paging: p,
	}, err
}

// buildSingleResult 生成单用户数据
func buildSingleResult(user userModel.Result, roles roleModel.Roles, userRoles userRoleModel.UserRoles) userModel.Result {
	roleMap := gmap.New()
	for _, role := range roles {
		roleMap.Set(role.ID, role)
	}
	r := make(roleModel.Roles, 0)
	for _, userRole := range userRoles {
		if user.ID == userRole.UserID {
			role := roleMap.Get(userRole.RoleID).(roleModel.Entity)
			r = append(r, role)
		}
	}
	user.Roles = r
	return user
}

// buildResults 生成用户返回值
func buildResults(users userModel.Results, roles roleModel.Roles, userRoles userRoleModel.UserRoles) userModel.Results {
	roleMap := gmap.New()
	for _, role := range roles {
		roleMap.Set(role.ID, role)
	}
	userRes := make(userModel.Results, 0)
	for _, user := range users {
		r := make([]roleModel.Entity, 0)
		for _, userRole := range userRoles {
			if user.ID == userRole.UserID {
				role := roleMap.Get(userRole.RoleID).(roleModel.Entity)
				r = append(r, role)
			}
		}
		user.Roles = r
		userRes = append(userRes, user)
	}
	return userRes
}

// 查询用户
func queryUsers(req *userModel.QueryUserReq) (userModel.Results, *paging.Paging, error) {
	var userEntity userModel.Entity
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
		return nil, nil, errors.New("读取行数失败")
	}
	p := paging.Create(req.PageNum, req.PageSize, int(total))

	db.OrderBy(req.OrderColumn + " " + req.OrderType + " ")
	db.Limit(p.PageSize, p.StartNum)
	users := make(userModel.Results, 0)
	err = db.Table(&userEntity).Select("id , email , phone , nickname , username , enabled , created_at , updated_at").Find(&users)
	if err != nil {
		return nil, nil, err
	}
	return users, p, nil
}

// func queryUserRole

// QueryByID 通过id查询
func QueryByID(ID int) (*userModel.Result, error) {
	var user userModel.Result
	db := orm.Instance()
	if _, err := db.Table(new(userModel.Entity)).
		Select("id , email , phone , nickname , username , created_at , updated_at").
		Where("id = ?", ID).
		Get(&user); err != nil {
		return nil, err
	}
	userRoles, err := roleService.QueryUserRoles([]int{user.ID})
	if err != nil {
		return nil, err
	}
	roleIDs := userRoles.ToRoleIDs()
	roles, err := roleService.QueryRoles(roleIDs)
	if err != nil {
		return nil, err
	}
	res := buildSingleResult(user, roles, userRoles)
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
