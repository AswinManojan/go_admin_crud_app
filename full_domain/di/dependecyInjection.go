package di

import (
	"full_domain/controllers"
	"full_domain/db"
	"full_domain/handler"
	"full_domain/middleware"
	"full_domain/repository"
	"full_domain/service"
)

func Init() *controllers.ServerStruct {
	db := db.ConnectDB()
	jwt := middleware.NewJwtUtil()
	adminRepository := repository.NewAdminRepository(db)
	userRepository := repository.NewUserRepository(db)

	adminService := service.NewAdminService(adminRepository, jwt)
	userService := service.NewUserService(userRepository, jwt)

	adminHandlers := handler.NewAdminHandler(adminService)
	userHandlers := handler.NewUserHandler(userService)

	server := controllers.NewServer()

	userRoutes := controllers.NewUserRoute(userHandlers, server, jwt)
	adminRoutes := controllers.NewAdminRoute(adminHandlers, server, jwt)
	adminRoutes.Routes()
	userRoutes.URoutes()

	return server
}
