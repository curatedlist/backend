package server

import (
	"backend/internal/list"

	"github.com/gin-gonic/gin"
)

// Server the http Server
type Server struct {
	router  *gin.Engine
	listAPI list.API
}

// Init initialize the http server
func Init(listAPI list.API) *Server {
	server := Server{router: gin.Default(), listAPI: listAPI}
	return server.registerAllRoutes()
}

// Run starts the http server
func (server *Server) Run() *Server {
	err := server.router.Run()
	if err != nil {
		panic(err)
	}
	return server
}

func (server *Server) registerAllRoutes() *Server {
	return server.registerStatusRoutes().registerListRoutes()
}

func (server *Server) registerStatusRoutes() *Server {
	server.router.GET("/status", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "OK",
		})
	})
	return server
}

func (server *Server) registerListRoutes() *Server {
	lists := server.router.Group("/lists")
	lists.GET("", server.listAPI.FindAll)
	return server
}
