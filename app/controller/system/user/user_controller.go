package user

import (
	userModel "gf-admin/app/model/system/user"
	userService "gf-admin/app/service/system/user"
	"gf-admin/library/response"

	"github.com/gogf/gf/net/ghttp"
)

// Create 创建
// QueryPage 创建
// @Summary 创建
// @Description 创建
// @Tags 系统 用户
// @accept json
// @Produce  json
// @Param body body user.CreateUserReq true "model.SwagGroupAdd"
// @Success 200 {object} response.Response
// @Router /system/user/create [post]
// @Security ApiKeyAuth
func Create(r *ghttp.Request) {
	var req *userModel.CreateUserReq
	if err := r.Parse(&req); err != nil {
		response.Res(r).BadRequest(err.Error())
	}
	id, err := userService.Create(req)
	if err != nil {
		response.Res(r).BadRequest(err.Error())
	}
	response.Res(r).Success(id)
}

// Update 修改
// @Summary 创建
// @Description 创建
// @Tags 系统 用户
// @accept json
// @Produce  json
// @Param body body user.UpdateUserReq true "model.SwagGroupAdd"
// @Success 200 {object} response.Response
// @Router /system/user/update [post]
// @Security ApiKeyAuth
func Update(r *ghttp.Request) {
	var req *userModel.UpdateUserReq
	if err := r.Parse(&req); err != nil {
		response.Res(r).BadRequest(err.Error())
	}
	id, err := userService.Update(req)
	if err != nil {
		response.Res(r).BadRequest(err.Error())
	}
	response.Res(r).Success(id)
}

// QueryPage 分页查询
// @Summary 分页查询
// @Description 分页查询
// @Tags 系统 用户
// @accept json
// @Produce  json
// @Param query query user.QueryUserReq true "model.SwagGroupAdd"
// @Success 200 {object} response.Response
// @Router /system/user/queryPage [get]
// @Security ApiKeyAuth
func QueryPage(r *ghttp.Request) {
	var req *userModel.QueryUserReq
	if err := r.Parse(&req); err != nil {
		response.Res(r).BadRequest(err.Error())
	}
	res, err := userService.QueryPage(req)
	if err != nil {
		response.Res(r).BadRequest(err.Error())
	}
	response.Res(r).Success(res)
}
