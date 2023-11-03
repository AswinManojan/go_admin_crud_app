package controllers

import (
	"full_domain/handler"
	"full_domain/middleware"
)

type AdminRouters struct {
	router *ServerStruct
	admin  *handler.AdminHandler
	jwt    *middleware.JwtUtil
}

func (as *AdminRouters) Routes() {
	as.router.r.POST("api/admin/add", as.jwt.ValidateToken("admin"), as.admin.AddUser)
	as.router.r.PUT("api/admin/update/:id", as.jwt.ValidateToken("admin"), as.admin.UpdateUser)
	as.router.r.GET("api/admin/block/:id", as.jwt.ValidateToken("admin"), as.admin.BlockUser)
	as.router.r.GET("api/admin/unblock/:id", as.jwt.ValidateToken("admin"), as.admin.UnBlockUser)
	as.router.r.DELETE("api/admin/delete/:id", as.jwt.ValidateToken("admin"), as.admin.DeleteUser)
	as.router.r.GET("api/admin/findall", as.jwt.ValidateToken("admin"), as.admin.FindAllUsers)
	as.router.r.GET("api/admin/find/:id", as.jwt.ValidateToken("admin"), as.admin.FindUser)
	as.router.r.GET("api/admin/search", as.jwt.ValidateToken("admin"), as.admin.SearchUser)
	as.router.r.POST("api/admin/login", as.admin.Login)
}

func NewAdminRoute(a *handler.AdminHandler, server *ServerStruct, jwt *middleware.JwtUtil) *AdminRouters {
	return &AdminRouters{
		router: server,
		admin:  a,
	}
}
