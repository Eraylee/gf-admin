package menu

import "gf-admin/app/model/base"

// QueryMenuReq 分页查询
type QueryMenuReq struct {
	// 名称
	Name     string `p:"name"`
	ParentID int    `p:"parentId"`
	// 类型
	Type string `p:"type"`
	// 路径
	Path string `p:"path"`
	// 是否可用
	Enabled int `p:"enabled"`

	base.PagingQueryReq
}

// QueryTreeReq 查询菜单树
type QueryTreeReq struct {
	// 名称
	Name string `p:"name"`
	// 类型
	Type int `p:"type"`
	// 路径
	Path string `p:"path"`
	// 是否可见
	Visiable int `p:"visiable"`
}

// CreateMenuReq 新增菜单资料请求参数
type CreateMenuReq struct {
	// 名称
	Name     string `p:"name"  v:"required#请输入菜单名称"`
	ParentID int    `p:"parentId"`
	// 图标
	Icon string `p:"icon"`
	// 类型
	Type int `p:"type"`
	// 路径
	Path string `p:"path"`
	// 动作
	Action string `p:"action"`
	// 排序
	Sort int `p:"sort"`
	// 权限
	Permission string `p:"permission"`
	// 是否可用
	Enabled int `p:"enabled"`
	// 是否可见
	Visiable int `p:"visiable"`
}

//UpdateMenuReq 修改菜单资料请求参数
type UpdateMenuReq struct {
	ID int `p:"id"  v:"required#请输入id"`
	// 名称
	Name     string `p:"name"  v:"required#请输入菜单名称"`
	ParentID int    `p:"parentId"`
	// 图标
	Icon string `p:"icon"`
	// 动作
	Action string `p:"action"`
	// 类型
	Type int `p:"type"`
	// 路径
	Path string `p:"path"`
	// 排序
	Sort int `p:"sort"`
	// 权限
	Permission string `p:"permission"`
	// 是否可用
	Enabled int `p:"enabled"`
	// 是否可见
	Visiable int `p:"visiable"`
}

//TreeItem 树形菜单子项
type TreeItem struct {
	ID         int    `json:"id"`
	Name       string `json:"name"`
	Icon       string `json:"icon"`
	Action     string `json:"action"`
	Type       int    `json:"type"`
	Path       string `json:"path"`
	Sort       int    `json:"sort"`
	Permission string `json:"permission"`
	Enabled    int    `json:"enabled"`
	Visiable   int    `json:"visiable"`

	Children Tree `json:"children"`
}

// Tree 菜单树
type Tree []TreeItem

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

// ToPaths 转换成路径
func (e *Menus) ToPaths() []string {
	ids := make([]string, 0)
	for _, item := range *e {
		ids = append(ids, item.Path)
	}
	return ids
}

//GetTree 获取菜单树
func (e *Menus) GetTree(data Menus, ID int) Tree {
	tree := make(Tree, 0)
	for _, v := range data {
		if v.ParentID != ID {
			continue
		}
		children := e.GetTree(data, v.ID)
		item := TreeItem{
			ID:       v.ID,
			Name:     v.Name,
			Sort:     v.Sort,
			Visiable: v.Visiable,
			Action:   v.Action,
			Icon:     v.Icon,
			Type:     v.Type,
			Path:     v.Path,
			Children: children,
		}
		tree = append(tree, item)
	}
	return tree
}
