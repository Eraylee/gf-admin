package role

import (
	roleModel "gf-admin/app/model/system/role"
	"gf-admin/library/orm"

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
