package menu

import (
	"gf-admin/app/model/base"
	menuModel "gf-admin/app/model/system/menu"
	RoleMenuService "gf-admin/app/model/system/role_menu"
	menuService "gf-admin/app/service/system/menu"
	"gf-admin/library/response"

	"github.com/gogf/gf/net/ghttp"
)

// Create 创建
// QueryPage 创建
// @Summary 创建
// @Description 创建
// @Tags 系统 菜单
// @accept json
// @Produce  json
// @Param data body menu.CreateMenuReq true "model.SwagGroupAdd"
// @Success 200 {object} response.Response
// @Router /system/menu/create [post]
// @Security ApiKeyAuth
func Create(r *ghttp.Request) {
	var req *menuModel.CreateMenuReq
	if err := r.Parse(&req); err != nil {
		response.Res(r).BadRequest(err.Error())
	}
	id, err := menuService.Create(req)
	if err != nil {
		response.Res(r).BadRequest(err.Error())
	}
	response.Res(r).Success(id)
}

// Update 修改
// @Summary 修改
// @Description 修改
// @Tags 系统 菜单
// @accept json
// @Produce  json
// @Param data body menu.UpdateMenuReq true "model.SwagGroupAdd"
// @Success 200 {object} response.Response
// @Router /system/menu/update [post]
// @Security ApiKeyAuth
func Update(r *ghttp.Request) {
	var req *menuModel.UpdateMenuReq
	if err := r.Parse(&req); err != nil {
		response.Res(r).BadRequest(err.Error())
	}
	id, err := menuService.Update(req)
	if err != nil {
		response.Res(r).BadRequest(err.Error())
	}
	response.Res(r).Success(id)
}

// QueryPage 分页查询
// @Summary 分页查询
// @Description 分页查询
// @Tags 系统 菜单
// @accept json
// @Produce  json
// @Param data query menu.QueryMenuReq true "model.SwagGroupAdd"
// @Success 200 {object} response.Response
// @Router /system/menu/queryPage [get]
// @Security ApiKeyAuth
func QueryPage(r *ghttp.Request) {
	var req *menuModel.QueryMenuReq
	if err := r.Parse(&req); err != nil {
		response.Res(r).BadRequest(err.Error())
	}
	res, err := menuService.QueryPage(req)
	if err != nil {
		response.Res(r).BadRequest(err.Error())
	}
	response.Res(r).Success(res)
}

// QueryTree 查询树形菜单
// @Summary 查询树形菜单
// @Description 查询树形菜单
// @Tags 系统 菜单
// @accept json
// @Produce  json
// @Param data query menu.QueryTreeReq true "model.SwagGroupAdd"
// @Success 200 {object} response.Response
// @Router /system/menu/queryTree [get]
// @Security ApiKeyAuth
func QueryTree(r *ghttp.Request) {
	var req *menuModel.QueryTreeReq
	if err := r.Parse(&req); err != nil {
		response.Res(r).BadRequest(err.Error())
	}
	res, err := menuService.QueryTree(req)
	if err != nil {
		response.Res(r).BadRequest(err.Error())
	}
	response.Res(r).Success(res)
}

// QueryByID 根据ID查询菜单
// @Summary 根据ID查询菜单
// @Description 根据ID查询菜单
// @Tags 系统 菜单
// @accept json
// @Produce  json
// @Param id query int true "model.SwagGroupAdd"
// @Success 200 {object} response.Response
// @Router /system/menu/queryById [get]
// @Security ApiKeyAuth
func QueryByID(r *ghttp.Request) {
	ID := r.GetQueryInt("id")
	res, err := menuService.QueryByID(ID)
	if err != nil {
		response.Res(r).BadRequest(err.Error())
	}
	response.Res(r).Success(res)
}

// Delete 删除菜单
// @Summary 删除菜单
// @Description 删除菜单
// @Tags 系统 菜单
// @accept json
// @Produce  json
// @Param data query base.DeleteReq true "model.SwagGroupAdd"
// @Success 200 {object} response.Response
// @Router /system/menu/delete [post]
// @Security ApiKeyAuth
func Delete(r *ghttp.Request) {
	var req base.DeleteReq
	if err := r.Parse(&req); err != nil {
		response.Res(r).BadRequest(err.Error())
	}
	res, err := menuService.Delete(&req)
	if err != nil {
		response.Res(r).BadRequest(err.Error())
	}
	response.Res(r).Success(res)
}

// CancelMenuConnect 取消关联菜单
// @Summary 取消关联菜单
// @Description 取消关联菜单
// @Tags 系统 菜单
// @accept json
// @Produce  json
// @Param data body role_menu.CancelConnectReq true "model.SwagGroupAdd"
// @Success 200 {object} response.Response
// @Router /system/menu/cancelMenuConnect [post]
// @Security ApiKeyAuth
func CancelMenuConnect(r *ghttp.Request) {
	var req RoleMenuService.CancelConnectReq
	if err := r.Parse(&req); err != nil {
		response.Res(r).BadRequest(err.Error())
	}
	res, err := menuService.CancelConnectByMenuID(&req)
	if err != nil {
		response.Res(r).BadRequest(err.Error())
	}
	response.Res(r).Success(res)
}
