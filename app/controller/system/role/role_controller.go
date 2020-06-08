package role

import (
	"gf-admin/app/model/base"
	roleModel "gf-admin/app/model/system/role"
	roleService "gf-admin/app/service/system/role"
	"gf-admin/library/response"

	"github.com/gogf/gf/net/ghttp"
)

// Create 创建
// QueryPage 创建
// @Summary 创建
// @Description 创建
// @Tags 系统 角色
// @accept json
// @Produce  json
// @Param body body role.CreateRoleReq true "model.SwagGroupAdd"
// @Success 200 {object} response.Response
// @Router /system/role/create [post]
// @Security ApiKeyAuth
func Create(r *ghttp.Request) {
	var req *roleModel.CreateRoleReq
	if err := r.Parse(&req); err != nil {
		response.Res(r).BadRequest(err.Error())
	}
	id, err := roleService.Create(req)
	if err != nil {
		response.Res(r).BadRequest(err.Error())
	}
	response.Res(r).Success(id)
}

// Update 修改
// @Summary 创建
// @Description 创建
// @Tags 系统 角色
// @accept json
// @Produce  json
// @Param body body role.UpdateRoleReq true "model.SwagGroupAdd"
// @Success 200 {object} response.Response
// @Router /system/role/update [post]
// @Security ApiKeyAuth
func Update(r *ghttp.Request) {
	var req *roleModel.UpdateRoleReq
	if err := r.Parse(&req); err != nil {
		response.Res(r).BadRequest(err.Error())
	}
	id, err := roleService.Update(req)
	if err != nil {
		response.Res(r).BadRequest(err.Error())
	}
	response.Res(r).Success(id)
}

// QueryPage 分页查询
// @Summary 分页查询
// @Description 分页查询
// @Tags 系统 角色
// @accept json
// @Produce  json
// @Param query query role.QueryRoleReq true "model.SwagGroupAdd"
// @Success 200 {object} response.Response
// @Router /system/role/queryPage [get]
// @Security ApiKeyAuth
func QueryPage(r *ghttp.Request) {
	var req *roleModel.QueryRoleReq
	if err := r.Parse(&req); err != nil {
		response.Res(r).BadRequest(err.Error())
	}
	res, err := roleService.QueryPage(req)
	if err != nil {
		response.Res(r).BadRequest(err.Error())
	}
	response.Res(r).Success(res)
}

// QuertByID 根据ID查询角色
// @Summary 根据ID查询角色
// @Description 根据ID查询角色
// @Tags 系统 角色
// @accept json
// @Produce  json
// @Param query id int true "model.SwagGroupAdd"
// @Success 200 {object} response.Response
// @Router /system/role/queryById [get]
// @Security ApiKeyAuth
func QuertByID(r *ghttp.Request) {
	ID := r.GetQueryInt("id")
	res, err := roleService.QueryByID(ID)
	if err != nil {
		response.Res(r).BadRequest(err.Error())
	}
	response.Res(r).Success(res)
}

// Delete 删除角色
// @Summary 删除角色
// @Description 删除角色
// @Tags 系统 角色
// @accept json
// @Produce  json
// @Param query data base.DeleteReq true "model.SwagGroupAdd"
// @Success 200 {object} response.Response
// @Router /system/role/delete [get]
// @Security ApiKeyAuth
func Delete(r *ghttp.Request) {
	var req base.DeleteReq
	if err := r.Parse(&req); err != nil {
		response.Res(r).BadRequest(err.Error())
	}
	res, err := roleService.Delete(&req)
	if err != nil {
		response.Res(r).BadRequest(err.Error())
	}
	response.Res(r).Success(res)
}
