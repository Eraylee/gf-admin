package paging

import (
	"math"
)

// Paging 分页参数
type Paging struct {
	PageNum   int `json:"pageNum"`   //当前页
	PageSize  int `json:"pageSize"`  //每页条数
	Total     int `json:"total"`     //总条数
	PageCount int `json:"pageCount"` //总页数
	StartNum  int `json:"-"`         //起始行
}

// New 创建分页
func Create(pageNum, pageSize, total int) *Paging {
	if pageNum < 1 {
		pageNum = 1
	}
	if pageSize < 1 {
		pageSize = 10
	}

	pageCount := math.Ceil(float64(total) / float64(pageSize))
	startNum := pageSize * (pageNum - 1)
	paging := new(Paging)
	paging.PageNum = pageNum
	paging.PageSize = pageSize
	paging.Total = total
	paging.PageCount = int(pageCount)
	paging.StartNum = startNum
	return paging
}
