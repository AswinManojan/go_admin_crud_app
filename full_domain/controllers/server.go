package controllers

import "github.com/gin-gonic/gin"

type ServerStruct struct {
	r *gin.Engine
}

func (s *ServerStruct) StartServer() {
	s.r.Run(":8080")
}

func NewServer() *ServerStruct {
	router := gin.Default()
	return &ServerStruct{
		r: router,
	}
}
