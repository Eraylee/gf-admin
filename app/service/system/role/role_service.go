package role

import (
	"errors"
	"gf-admin/app/model/base"
	roleModel "gf-admin/app/model/system/role"
	"gf-admin/library/orm"
	"gf-admin/library/paging"

	"github.com/gogf/gf/errors/gerror"
)

// Create 创建
func Create(req *roleModel.CreateRoleReq) (int, error) {
	db := orm.Instance().NewSession()
	role := roleModel.Entity{
		Name:    req.Name,
		Code:    req.Code,
		Sort:    req.Sort,
		Admin:   req.Admin,
		Enabled: req.Enabled,
	}

	if _, err := db.Insert(&role); err != nil {
		return 0, err
	}

	return role.ID, nil

}

// Update 更新
func Update(req *roleModel.UpdateRoleReq) (int, error) {
	db := orm.Instance()
	var role roleModel.Entity
	if has, err := db.ID(req.ID).Get(&role); err != nil {
		return 0, err
	} else if !has {
		return 0, gerror.New("角色不存在")
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
	if _, err := db.Update(&role); err != nil {
		return 0, err
	}
	return role.ID, nil
}

// QueryPage 分页查询
func QueryPage(req *roleModel.QueryRoleReq) ([]roleModel.Entity, error) {
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

	res := make([]roleModel.Entity, 0)
	err = db.Table(&userEntity).Desc("created_at").Find(&res)
	return res, err
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
func Delete(IDs *base.DeleteReq) (int64, error) {
	var role roleModel.Entity
	db := orm.Instance()

	res, err := db.Where("id = ?", IDs).Delete(&role)
	if err != nil {
		return 0, err
	}

	return res, nil
}
