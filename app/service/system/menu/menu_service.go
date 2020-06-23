package menu

import (
	"gf-admin/app/model/base"
	menuModel "gf-admin/app/model/system/menu"
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
	if req.Type != "" {
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
func QueryPage(req *menuModel.QueryMenuReq) ([]menuModel.Entity, error) {
	var userEntity menuModel.Entity
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

	total, err := db.Count(&userEntity)

	if err != nil {
		return nil, gerror.New("读取行数失败")
	}

	p := paging.Create(req.PageNum, req.PageSize, int(total))

	db.OrderBy(req.OrderColumn + " " + req.OrderType + " ")
	db.Select("name , type , visiable ,  action , icon , type ,target , createdAt ,updatedAt ,parentId ")
	db.Limit(p.PageSize, p.StartNum)

	res := make([]menuModel.Entity, 0)
	err = db.Table(&userEntity).Desc("created_at").Find(&res)
	return res, err
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

	if req.Type != "" {
		db.Where("type like ?", "%"+req.Type+"%")
	}

	if req.Visiable != 0 {
		db.Where("visiable = ?", req.Visiable)
	}

	db.Select("name , type , visiable ,  action , icon , type ,target ,parentId ")

	res := make([]menuModel.Entity, 0)
	if err := db.Find(&res); err != nil {
		return nil, err
	}
	return toTreeData(res), nil

}

// toTreeData 转换成树形数据
func toTreeData(data []menuModel.Entity) []menuModel.TreeItem {

	tree := make([]menuModel.TreeItem, 0)
	for _, v := range data {
		if v.ParentID == 0 {
			continue
		}
		children := getChildren(data, v)
		menu := menuModel.TreeItem{
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
		tree = append(tree, menu)
	}
	return tree
}

// getChildren 获取子节点
func getChildren(data []menuModel.Entity, item menuModel.Entity) []menuModel.TreeItem {
	tree := make([]menuModel.TreeItem, 0)
	for _, v := range data {

		if v.ParentID != item.ID {
			continue
		}

		children := getChildren(data, v)
		menu := menuModel.TreeItem{
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
		tree = append(tree, menu)
	}
	return tree
}

// CancelConnectByMenuID 取消关联
func CancelConnectByMenuID(req *menuRoleModel.CancelConnectReq) (bool, error) {
	var menuRole menuRoleModel.Entity
	db := orm.Instance()
	if _, err := db.Where("role_id = ? AND menu_id in (?)", req.RoleID, req.MenuIDs).Delete(&menuRole); err != nil {
		return false, err
	}
	return true, nil
}
