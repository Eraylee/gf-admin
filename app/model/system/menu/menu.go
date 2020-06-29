package menu

import "gf-admin/app/model/base"

// QueryMenuReq 分页查询
type QueryMenuReq struct {
	Name     string `p:"name"`
	ParentID int    `p:"parentId"`
	Type     string `p:"type"`
	Target   string `p:"target"`
	Enabled  int    `p:"enabled"`

	base.PagingQueryReq
}

// QueryTreeReq 查询菜单树
type QueryTreeReq struct {
	Name     string `p:"name"`
	Type     int    `p:"type"`
	Target   string `p:"target"`
	Visiable int    `p:"visiable"`
}

// CreateMenuReq 新增菜单资料请求参数
type CreateMenuReq struct {
	Name       string `p:"name"  v:"required#请输入菜单名称"`
	ParentID   int    `p:"parentId"`
	Icon       string `p:"icon"`
	Type       int    `p:"type"`
	Target     string `p:"target"`
	Action     string `p:"action"`
	Sort       int    `p:"sort"`
	Permission string `p:"permission"`
	Enabled    int    `p:"enabled"`
	Visiable   int    `p:"visiable"`
}

//UpdateMenuReq 修改菜单资料请求参数
type UpdateMenuReq struct {
	ID         int    `p:"id"  v:"required#请输入id"`
	Name       string `p:"name"  v:"required#请输入菜单名称"`
	ParentID   int    `p:"parentId"`
	Icon       string `p:"icon"`
	Action     string `p:"action"`
	Type       int    `p:"type"`
	Target     string `p:"target"`
	Sort       int    `p:"sort"`
	Permission string `p:"permission"`
	Enabled    int    `p:"enabled"`
	Visiable   int    `p:"visiable"`
}

//TreeItem 树形菜单子项
type TreeItem struct {
	ID         int    `json:"id"`
	Name       string `json:"name"`
	Icon       string `json:"icon"`
	Action     string `json:"action"`
	Type       int    `json:"type"`
	Target     string `json:"target"`
	Sort       int    `json:"sort"`
	Permission string `json:"permission"`
	Enabled    int    `json:"enabled"`
	Visiable   int    `json:"visiable"`

	Children []TreeItem `json:"children"`
}

// Menus 菜单数组
type Menus []Entity

// ToIDs 转换成id数组
func (e *Menus) ToIDs() []int {
	ids := make([]int, 0)
	for _, item := range *e {
		ids = append(ids, item.ID)
	}
	return ids
}
