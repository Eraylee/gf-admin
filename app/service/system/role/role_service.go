package role

import (
	"errors"
	"gf-admin/app/model/base"
	casbinRuleModel "gf-admin/app/model/system/casbin_rule"
	menuModel "gf-admin/app/model/system/menu"
	roleModel "gf-admin/app/model/system/role"
	roleMenuModel "gf-admin/app/model/system/role_menu"
	userRoleModel "gf-admin/app/model/system/user_role"
	menuService "gf-admin/app/service/system/menu"
	"gf-admin/library/orm"
	"gf-admin/library/paging"

	"github.com/gogf/gf/container/gmap"
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
	if has, err := session.Where("code = ?", req.Code).Exist(&role); err != nil {
		return 0, err
	} else if has {
		return 0, gerror.New("code已存在")
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
			menuRoles := make([]roleMenuModel.Entity, 0)
			casbinRules := make([]casbinRuleModel.Entity, 0)

			for i := range menus {
				menu := menus[i]

				userRole := roleMenuModel.Entity{
					RoleID: role.ID,
					MenuID: menu.ID,
				}
				casbinRule := casbinRuleModel.Entity{
					PType: "p",
					V0:    role.Code,
					V1:    menu.Path,
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
		menuRoles := make([]roleMenuModel.Entity, 0)
		casbinRules := make([]casbinRuleModel.Entity, 0)

		for _, menu := range menus {
			userRole := roleMenuModel.Entity{
				RoleID: role.ID,
				MenuID: menu.ID,
			}
			casbinRule := casbinRuleModel.Entity{
				PType: "p",
				V0:    role.Code,
				V1:    menu.Path,
				V2:    menu.Action,
			}
			menuRoles = append(menuRoles, userRole)
			casbinRules = append(casbinRules, casbinRule)
		}

		if len(menuRoles) > 0 {
			if _, err := session.Where("menu_id = ?", req.ID).Delete(new(roleMenuModel.Entity)); err != nil {
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
func QueryPage(req *roleModel.QueryRoleReq) (*base.PagingRes, error) {
	// 分页查询角色信息
	roles, p, err := queryRoles(req)
	if err != nil {
		return nil, err
	}
	//获取角色对应的关系
	roleIDs := roles.ToIDs()
	roleMenus, err := menuService.QueryRoleMenus(roleIDs)
	if err != nil {
		return nil, err
	}
	//获取所有关联菜单的详情
	menuIDs := roleMenus.ToMenuIDs()
	menus, err := menuService.QueryMenus(menuIDs)
	if err != nil {
		return nil, err
	}

	res := buildResults(roles, menus, roleMenus)

	return &base.PagingRes{
		Data:   res,
		Paging: p,
	}, err
}

// queryRoles 查询角色
func queryRoles(req *roleModel.QueryRoleReq) (roleModel.Results, *paging.Paging, error) {
	var role roleModel.Entity
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

	total, err := db.Count(&role)

	if err != nil {
		return nil, nil, errors.New("读取行数失败")
	}

	p := paging.Create(req.PageNum, req.PageSize, int(total))

	db.OrderBy(req.OrderColumn + " " + req.OrderType + " ")

	db.Limit(p.PageSize, p.StartNum)

	roles := make(roleModel.Results, 0)
	err = db.Table(&role).Find(&roles)
	return roles, p, err
}

// buildResults 生成用户返回值
func buildResults(roles roleModel.Results, menus menuModel.Menus, roleMenus roleMenuModel.RoleMenus) roleModel.Results {
	menuMap := gmap.New()
	for _, role := range menus {
		menuMap.Set(role.ID, role)
	}
	roleRes := make(roleModel.Results, 0)
	for _, role := range roles {
		m := make(menuModel.Menus, 0)
		for _, roleMenu := range roleMenus {
			if role.ID == roleMenu.RoleID {
				menu := menuMap.Get(roleMenu.MenuID).(menuModel.Entity)
				m = append(m, menu)
			}
		}
		role.Menus = m
		roleRes = append(roleRes, role)
	}
	return roleRes
}

// buildSingleResult 生成单用户数据
func buildSingleResult(role roleModel.Result, menus menuModel.Menus, roleMenus roleMenuModel.RoleMenus) roleModel.Result {
	menuMap := gmap.New()
	for _, role := range menus {
		menuMap.Set(role.ID, role)
	}
	m := make(menuModel.Menus, 0)
	for _, roleMenu := range roleMenus {
		if role.ID == roleMenu.RoleID {
			menu := menuMap.Get(roleMenu.MenuID).(menuModel.Entity)
			m = append(m, menu)
		}
	}
	role.Menus = m
	return role
}

// QueryByID 通过id查询
func QueryByID(ID int) (*roleModel.Result, error) {
	var role roleModel.Result
	db := orm.Instance()
	if _, err := db.Table(new(roleModel.Entity)).ID(ID).Get(&role); err != nil {
		return nil, err
	}
	roleMenus, err := menuService.QueryRoleMenus([]int{role.ID})
	if err != nil {
		return nil, err
	}
	menuIDs := roleMenus.ToMenuIDs()
	menus, err := menuService.QueryMenus(menuIDs)
	if err != nil {
		return nil, err
	}
	res := buildSingleResult(role, menus, roleMenus)
	return &res, nil
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
