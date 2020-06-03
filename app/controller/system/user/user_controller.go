package user_c

import (
	user_m "gf-admin/app/model/system/user"
	user_s "gf-admin/app/service/system/user"
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
	var req *user_m.CreateUserReq
	if err := r.Parse(&req); err != nil {
		response.Res(r).BadRequest(err.Error())
	}
	id, err := user_s.Create(req)
	if err != nil {
		response.Res(r).BadRequest(err.Error())
	}
	response.Res(r).Success(id)
}
