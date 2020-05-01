package server

import "github.com/gin-gonic/gin"

// Server the http Server
type Server struct {
	router *gin.Engine
}

// Init initialize the http server
func Init() *Server {
	server := Server{router: gin.Default()}
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
	return server.registerStatusRoutes()
}

func (server *Server) registerStatusRoutes() *Server {
	server.router.GET("/status", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "OK",
		})
	})
	return server
}
