package menu

import (
	"gf-admin/app/model/base"
	casbinRuleModel "gf-admin/app/model/system/casbin_rule"
	menuModel "gf-admin/app/model/system/menu"
	roleModel "gf-admin/app/model/system/role"
	menuRoleModel "gf-admin/app/model/system/role_menu"
	"gf-admin/library/orm"
	"gf-admin/library/paging"

	"github.com/gogf/gf/errors/gerror"
)

// Create 创建
func Create(req *menuModel.CreateMenuReq) (int, error) {
	db := orm.Instance()
	menu := menuModel.Entity{
		Name:     req.Name,
		Sort:     req.Sort,
		Enabled:  req.Enabled,
		Visiable: req.Visiable,
		ParentID: req.ParentID,
		Icon:     req.Icon,
		Type:     req.Type,
		Target:   req.Target,
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
	if req.Target != "" {
		menu.Target = req.Target
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
func QueryTree(req *menuModel.QueryTreeReq) ([]menuModel.TreeItem, error) {
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

	res := make([]menuModel.Entity, 0)
	if err := db.Find(&res); err != nil {
		return nil, err
	}

	return getTree(res, 0), nil

}

// getTree 获取树节点
func getTree(data menuModel.Menus, ID int) []menuModel.TreeItem {
	tree := make([]menuModel.TreeItem, 0)
	for _, v := range data {
		if v.ParentID != ID {
			continue
		}
		children := getTree(data, v.ID)
		item := menuModel.TreeItem{
			ID:       v.ID,
			Name:     v.Name,
			Sort:     v.Sort,
			Visiable: v.Visiable,
			Action:   v.Action,
			Icon:     v.Icon,
			Type:     v.Type,
			Target:   v.Target,
			Children: children,
		}
		tree = append(tree, item)
	}
	return tree
}

// CancelConnectByMenuID 取消关联
func CancelConnectByMenuID(req *menuRoleModel.CancelConnectReq) (bool, error) {
	var menuRole menuRoleModel.Entity
	session := orm.Instance().NewSession()
	if err := session.Begin(); err != nil {
		return false, err
	}
	defer session.Close()

	var role roleModel.Entity
	if has, err := session.Where("role_id = ?", req.RoleID).Get(&role); err != nil {
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
	if _, err := session.In("v1", menus.ToIDs).And("V0 = ?", role.Code).Delete(new(casbinRuleModel.Entity)); err != nil {
		return false, err
	}
	return true, session.Commit()
}
