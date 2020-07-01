package menu

import (
	"gf-admin/app/model/base"
	casbinRuleModel "gf-admin/app/model/system/casbin_rule"
	menuModel "gf-admin/app/model/system/menu"
	roleModel "gf-admin/app/model/system/role"
	roleMenuModel "gf-admin/app/model/system/role_menu"
	"gf-admin/library/orm"
	"gf-admin/library/paging"

	"github.com/gogf/gf/errors/gerror"
)

// Create 创建
func Create(req *menuModel.CreateMenuReq) (int, error) {
	db := orm.Instance()
	menu := menuModel.Entity{
		Name:     req.Name,
		Action:   req.Action,
		Sort:     req.Sort,
		Enabled:  req.Enabled,
		Visiable: req.Visiable,
		ParentID: req.ParentID,
		Icon:     req.Icon,
		Type:     req.Type,
		Path:     req.Path,
	}

	if _, err := db.Insert(&menu); err != nil {
		return 0, err
	}

	return menu.ID, nil
}

// Update 更新
func Update(req *menuModel.UpdateMenuReq) (int, error) {
	db := orm.Instance()
	var menu menuModel.Entity

	if has, err := db.ID(req.ID).Get(&menu); err != nil {
		return 0, err
	} else if !has {
		return 0, gerror.New("角色不存在")
	}

	if req.Name != "" {
		menu.Name = req.Name
	}
	if req.Action != "" {
		menu.Action = req.Action
	}
	if req.Path != "" {
		menu.Path = req.Path
	}
	if req.Type != 0 {
		menu.Type = req.Type
	}
	if req.ParentID != 0 {
		menu.ParentID = req.ParentID
	}
	if req.Visiable != 0 {
		menu.Visiable = req.Visiable
	}
	if req.Sort != 0 {
		menu.Sort = req.Sort
	}
	if req.Enabled != 0 {
		menu.Enabled = req.Enabled
	}

	if _, err := db.Update(&menu); err != nil {
		return 0, err
	}

	return menu.ID, nil
}

// QueryPage 分页查询
func QueryPage(req *menuModel.QueryMenuReq) (menuModel.Menus, error) {
	var menu menuModel.Entity
	//  orderColumn , orderType :=  "created_at" , ""
	db := orm.Instance()

	if req.Name != "" {
		db.Where("name like ?", "%"+req.Name+"%")
	}

	if req.Type != "" {
		db.Where("type like ?", "%"+req.Type+"%")
	}

	if req.Enabled != 0 {
		db.Where("enabled = ?", req.Enabled)
	}

	total, err := db.Count(&menu)

	if err != nil {
		return nil, gerror.New("读取行数失败")
	}

	p := paging.Create(req.PageNum, req.PageSize, int(total))

	db.OrderBy(req.OrderColumn + " " + req.OrderType + " ")
	db.Select("id , name , type , visiable ,  action , icon , type ,target , createdAt ,updatedAt ,parent_id ")
	db.Limit(p.PageSize, p.StartNum)

	var menus menuModel.Menus
	err = db.Table(&menu).Desc("created_at").Find(&menus)
	return menus, err
}

// QueryByID 通过id查询
func QueryByID(ID int) (*menuModel.Entity, error) {
	var menu menuModel.Entity
	db := orm.Instance()
	if _, err := db.ID(ID).Get(&menu); err != nil {
		return nil, err
	}
	return &menu, nil
}

// Delete 删除
func Delete(req *base.DeleteReq) (int64, error) {
	var menu menuModel.Entity
	db := orm.Instance()
	res, err := db.In("id", req.Ids).Delete(&menu)
	if err != nil {
		return 0, err
	}
	return res, nil
}

// QueryTree 查询菜单树
func QueryTree(req *menuModel.QueryTreeReq) (menuModel.Tree, error) {
	var menu menuModel.Entity
	//  orderColumn , orderType :=  "created_at" , ""
	db := orm.Instance().Table(&menu)

	if req.Name != "" {
		db.Where("name like ?", "%"+req.Name+"%")
	}

	if req.Type != 0 {
		db.Where("type = ?", req.Type)
	}

	if req.Visiable != 0 {
		db.Where("visiable = ?", req.Visiable)
	}

	db.Select("id , name , type , visiable ,  action , icon , type ,target ,parent_id ")

	menus := make(menuModel.Menus, 0)
	if err := db.Find(&menus); err != nil {
		return nil, err
	}

	return menus.GetTree(menus, 0), nil

}

// func QueryTreeBy

//QueryRoleMenus 查询菜单角色关系
func QueryRoleMenus(roleIDs []int) (roleMenuModel.RoleMenus, error) {
	db := orm.Instance()
	roleMenus := make(roleMenuModel.RoleMenus, 0)
	if err := db.Table(new(roleMenuModel.Entity)).In("role_id", roleIDs).Find(&roleMenus); err != nil {
		return nil, err
	}
	return roleMenus, nil
}

// QueryMenus 查询菜单
func QueryMenus(menuIDs []int) (menuModel.Menus, error) {
	db := orm.Instance()
	menus := make(menuModel.Menus, 0)
	if err := db.Table(new(menuModel.Entity)).In("id", menuIDs).Find(&menus); err != nil {
		return nil, err
	}
	return menus, nil
}

// CancelConnectByMenuID 取消关联
func CancelConnectByMenuID(req *roleMenuModel.CancelConnectReq) (bool, error) {
	var menuRole roleMenuModel.Entity
	session := orm.Instance().NewSession()
	if err := session.Begin(); err != nil {
		return false, err
	}
	defer session.Close()

	var role roleModel.Entity
	if has, err := session.ID(req.RoleID).Get(&role); err != nil {
		return false, err
	} else if !has {
		return false, gerror.New("角色不存在")
	}
	menus := make(menuModel.Menus, 0)
	if err := session.In("id", req.MenuIDs).Find(&menus); err != nil {
		return false, err
	}
	if len(menus) <= 0 {
		return false, gerror.New("菜单不存在")
	}
	if _, err := session.In("menu_id", req.MenuIDs).And("role_id = ?", req.RoleID).Delete(&menuRole); err != nil {
		return false, err
	}
	paths := menus.ToPaths()
	if _, err := session.In("v1", paths).And("V0 = ?", role.Code).Delete(new(casbinRuleModel.Entity)); err != nil {
		return false, err
	}
	return true, session.Commit()
}
