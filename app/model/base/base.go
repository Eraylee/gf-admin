package base

import "gf-admin/library/paging"

// DeleteReq 删除参数
type DeleteReq struct {
	Ids []int `p:"ids"  v:"required#请输入ids"`
}

//PagingQueryReq 通用查询参数
type PagingQueryReq struct {
	PageNum     int    `p:"pageNum"`     //当前页码
	PageSize    int    `p:"pageSize"`    //每页数
	OrderColumn string `p:"orderColumn"` //排序字段
	OrderType   string `p:"orderType"`   //排序方式
}

//PagingRes 分页查询返回
type PagingRes struct {
	Data interface{} `json:"data"`
	*paging.Paging
}
