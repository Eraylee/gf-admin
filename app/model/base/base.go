package base

import "gf-admin/library/paging"

// DeleteReq 删除参数
type DeleteReq struct {
	// 删除ids
	Ids []int `p:"ids"  v:"required#请输入ids"`
}

//PagingQueryReq 通用查询参数
type PagingQueryReq struct {
	//当前页码
	PageNum int `p:"pageNum"`
	//每页数
	PageSize int `p:"pageSize"`
	//排序字段
	OrderColumn string `p:"orderColumn"`
	//排序方式
	OrderType string `p:"orderType"`
}

//PagingRes 分页查询返回
type PagingRes struct {
	Data interface{} `json:"data"`
	*paging.Paging
}
