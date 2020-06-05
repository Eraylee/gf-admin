package menu

import "gf-admin/app/model/base"

// QueryMenuReq 分页查询
type QueryMenuReq struct {
	Name     string `p:"name"`
	ParentID int    `p:"parentId"`
	Type     string `p:"type"`
	Target   string `p:"target"`
	Enabled  string `p:"enabled"`

	base.PagingQueryReq
}

// CreateMenuReq 新增菜单资料请求参数
type CreateMenuReq struct {
	Name       string `p:"name"  v:"required#请输入菜单名称"`
	Code       string `p:"code"  v:"required#请输入菜单编码"`
	ParentID   int    `p:"parentId"`
	Icon       string `p:"icon"`
	Type       string `p:"type"`
	Target     string `p:"target"`
	Sort       int    `p:"sort"`
	Permission string `p:"permission"`
	Enabled    int    `p:"enabled"`
}

//UpdateMenuReq 修改菜单资料请求参数
type UpdateMenuReq struct {
	ID         int    `p:"id"  v:"required#请输入id"`
	Name       string `p:"name"  v:"required#请输入菜单名称"`
	Code       string `p:"code"  v:"required#请输入菜单编码"`
	ParentID   int    `p:"parentId"`
	Icon       string `p:"icon"`
	Type       string `p:"type"`
	Target     string `p:"target"`
	Sort       int    `p:"sort"`
	Permission string `p:"permission"`
	Enabled    int    `p:"enabled"`
}
