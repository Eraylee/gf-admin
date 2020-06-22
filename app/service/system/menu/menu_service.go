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

// CancelConnectByMenuID 取消关联
func CancelConnectByMenuID(req *menuRoleModel.CancelConnectReq) (bool, error) {
	var menuRole menuRoleModel.Entity
	db := orm.Instance()
	if _, err := db.Where("role_id = ? AND menu_id in (?)", req.RoleID, req.MenuIDs).Delete(&menuRole); err != nil {
		return false, err
	}
	return true, nil
}
