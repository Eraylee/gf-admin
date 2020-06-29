package role

import (
	"errors"
	"gf-admin/app/model/base"
	casbinRuleModel "gf-admin/app/model/system/casbin_rule"
	menuModel "gf-admin/app/model/system/menu"
	roleModel "gf-admin/app/model/system/role"
	menuRoleModel "gf-admin/app/model/system/role_menu"
	userRoleModel "gf-admin/app/model/system/user_role"
	"gf-admin/library/orm"
	"gf-admin/library/paging"

	"github.com/gogf/gf/errors/gerror"
)

// Create 创建
func Create(req *roleModel.CreateRoleReq) (int, error) {
	session := orm.Instance().NewSession()
	defer session.Close()
	// add Begin() before any action
	if err := session.Begin(); err != nil {
		return 0, err
	}
	role := roleModel.Entity{
		Name:    req.Name,
		Code:    req.Code,
		Sort:    req.Sort,
		Admin:   req.Admin,
		Enabled: req.Enabled,
	}
	if _, err := session.Insert(&role); err != nil {
		return 0, err
	}

	if len(req.MenuIDs) > 0 {
		var menus []menuModel.Entity
		if err := session.Table(new(menuModel.Entity)).In("id", req.MenuIDs).Find(&menus); err != nil {
			return 0, err
		}

		if len(menus) > 0 {
			menuRoles := make([]menuRoleModel.Entity, 0)
			casbinRules := make([]casbinRuleModel.Entity, 0)

			for i := range menus {
				menu := menus[i]

				userRole := menuRoleModel.Entity{
					RoleID: role.ID,
					MenuID: menu.ID,
				}
				casbinRule := casbinRuleModel.Entity{
					PType: "p",
					V0:    role.Code,
					V1:    menu.Target,
					V2:    menu.Action,
				}
				menuRoles = append(menuRoles, userRole)
				casbinRules = append(casbinRules, casbinRule)
			}

			if len(menuRoles) > 0 {
				if _, err := session.Insert(menuRoles); err != nil {
					return 0, err
				}
				if _, err := session.Insert(casbinRules); err != nil {
					return 0, err
				}
			}

		}
	}

	return role.ID, session.Commit()

}

// Update 更新
func Update(req *roleModel.UpdateRoleReq) (int, error) {
	session := orm.Instance().NewSession()
	defer session.Close()
	// add Begin() before any action
	if err := session.Begin(); err != nil {
		return 0, err
	}
	var role roleModel.Entity
	if has, err := session.ID(req.ID).Get(&role); err != nil {
		return 0, err
	} else if !has {
		return 0, gerror.New("角色不存在")
	}
	menus := make(menuModel.Menus, 0)
	if err := session.Table(new(menuModel.Entity)).In("id", req.MenuIDs).Find(&menus); err != nil {
		return 0, err
	}

	if req.Name != "" {
		role.Name = req.Name
	}
	if req.Admin != 0 {
		role.Admin = req.Admin
	}
	if req.Sort != 0 {
		role.Sort = req.Sort
	}
	if req.Enabled != 0 {
		role.Enabled = req.Enabled
	}
	if _, err := session.Update(&role); err != nil {
		return 0, err
	}

	if len(req.MenuIDs) > 0 {
		menuRoles := make([]menuRoleModel.Entity, 0)
		casbinRules := make([]casbinRuleModel.Entity, 0)

		for _, menu := range menus {
			userRole := menuRoleModel.Entity{
				RoleID: role.ID,
				MenuID: menu.ID,
			}
			casbinRule := casbinRuleModel.Entity{
				PType: "p",
				V0:    role.Code,
				V1:    menu.Target,
				V2:    menu.Action,
			}
			menuRoles = append(menuRoles, userRole)
			casbinRules = append(casbinRules, casbinRule)
		}

		if len(menuRoles) > 0 {
			if _, err := session.Where("menu_id = ?", req.ID).Delete(new(menuRoleModel.Entity)); err != nil {
				return 0, err
			}
			if _, err := session.Where("V0 = ?", role.Code).Delete(new(casbinRuleModel.Entity)); err != nil {
				return 0, err
			}
			if _, err := session.Insert(menuRoles); err != nil {
				return 0, err
			}
			if _, err := session.Insert(casbinRules); err != nil {
				return 0, err
			}
		}

	}
	return role.ID, session.Commit()
}

// QueryPage 分页查询
func QueryPage(req *roleModel.QueryRoleReq) (roleModel.Roles, error) {
	var userEntity roleModel.Entity
	//  orderColumn , orderType :=  "created_at" , ""
	db := orm.Instance()

	if req.Name != "" {
		db.Where("name like ?", "%"+req.Name+"%")
	}

	if req.Code != "" {
		db.Where("code like ?", "%"+req.Code+"%")
	}

	if req.Enabled != 0 {
		db.Where("enabled = ?", req.Enabled)
	}

	if req.Admin != 0 {
		db.Where("admin = ?", req.Admin)
	}

	total, err := db.Count(&userEntity)

	if err != nil {
		return nil, errors.New("读取行数失败")
	}

	p := paging.Create(req.PageNum, req.PageSize, int(total))

	db.OrderBy(req.OrderColumn + " " + req.OrderType + " ")

	db.Limit(p.PageSize, p.StartNum)

	var roles roleModel.Roles
	err = db.Table(&userEntity).Find(&roles)
	return roles, err
}

// QueryByID 通过id查询
func QueryByID(ID int) (*roleModel.Entity, error) {
	var role roleModel.Entity
	db := orm.Instance()
	if _, err := db.ID(ID).Get(&role); err != nil {
		return nil, err
	}
	return &role, nil
}

// Delete 删除
func Delete(req *base.DeleteReq) (int64, error) {
	var role roleModel.Entity
	db := orm.Instance()
	res, err := db.In("id", req.Ids).Delete(&role)
	if err != nil {
		return 0, err
	}
	return res, nil
}

// CancelConnectByUserID a取消关联
func CancelConnectByUserID(req *userRoleModel.CancelConnectReq) (bool, error) {
	var userRole userRoleModel.Entity
	db := orm.Instance()
	if _, err := db.In("role_id", req.RoleIDs).And("user_id = ?)", req.UserID).Delete(&userRole); err != nil {
		return false, err
	}
	return true, nil
}

//QueryUserRoles 查询用户角色关系
func QueryUserRoles(userIDs []int) (userRoleModel.UserRoles, error) {
	db := orm.Instance()
	userRoles := make(userRoleModel.UserRoles, 0)
	if err := db.Table(new(userRoleModel.Entity)).In("user_id", userIDs).Select("user_id,role_id").Find(&userRoles); err != nil {
		return nil, err
	}
	return userRoles, nil
}

// QueryRoles 查询角色
func QueryRoles(roleIDs []int) (roleModel.Roles, error) {
	db := orm.Instance()
	roles := make(roleModel.Roles, 0)
	if err := db.Table(new(roleModel.Entity)).In("id", roleIDs).Select("id,name,code,Admin,created_at,updated_at").Find(&roles); err != nil {
		return nil, err
	}
	return roles, nil
}
