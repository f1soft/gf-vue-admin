package v1

import (
	"github.com/gogf/gf/net/ghttp"
)

// @Tags Customer
// @Summary 创建客户
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Customer true "创建客户"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /customer/customer [post]
func CreateCustomer(r *ghttp.Request) {

}

// @Tags Customer
// @Summary 删除客户
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Customer true "删除客户"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /customer/customer [delete]
func DeleteCustomer(r *ghttp.Request) {
}

// @Tags Customer
// @Summary 更新客户信息
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Customer true "创建客户"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /customer/customer [put]
func UpdateCustomer(r *ghttp.Request) {
}

// @Tags Customer
// @Summary 获取单一客户信息
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Customer true "获取单一客户信息"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /customer/customer [get]
func FindCustomer(r *ghttp.Request) {
}

// @Tags Customer
// @Summary 获取权限客户列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.PageInfo true "获取权限客户列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /customer/customerList [get]
func GetCustomerList(r *ghttp.Request) {
}
