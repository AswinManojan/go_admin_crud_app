package controllers

import (
	"full_domain/handler"
	"full_domain/middleware"
)

type UserRouters struct {
	router *ServerStruct
	user   *handler.UserHandler
	jwt    *middleware.JwtUtil
}

func (as *UserRouters) URoutes() {
	as.router.r.POST("api/user/register", as.user.RegisterUser)
	as.router.r.POST("api/user/login", as.user.Login)
	as.router.r.GET("api/user/home", as.jwt.ValidateToken("user"), as.user.Home)
}

func NewUserRoute(a *handler.UserHandler, server *ServerStruct, jwt *middleware.JwtUtil) *UserRouters {
	return &UserRouters{
		router: server,
		user:   a,
		jwt:    jwt,
	}
}
