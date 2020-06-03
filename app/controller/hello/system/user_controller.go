package user

import (
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
	var req *userService.CreateUserReq
	if err := r.Parse(&req); err != nil {
		response.Res(r).BadRequest(err.Error())
	}
	id, err := userService.Create(req)
	if err != nil {
		response.Res(r).BadRequest(err.Error())
	}
	response.Res(r).Success(id)
}
