package server

import (
	"backend/internal/list"
	"backend/internal/user"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

// Server the http Server
type Server struct {
	router  *gin.Engine
	listAPI list.API
	userAPI user.API
}

// Init initialize the http server
func Init(listAPI list.API, userAPI user.API) *Server {
	server := Server{router: gin.Default(), listAPI: listAPI, userAPI: userAPI}
	return server.withCors().registerAllRoutes()
}

// Run starts the http server
func (server *Server) Run() *Server {
	err := server.router.Run()
	if err != nil {
		panic(err)
	}
	return server
}

func (server *Server) withCors() *Server {
	server.router.Use(cors.Default())
	return server
}

func (server *Server) registerAllRoutes() *Server {
	return server.registerStatusRoutes().registerListRoutes().registerUserRoutes()
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
	lists.GET("/", server.listAPI.FindAll)
	lists.GET("/id/:id", server.listAPI.Get)
	lists.POST("/", server.listAPI.CreateList)
	lists.POST("/:id/items/", server.listAPI.CreateItem)
	lists.PATCH("/:listID/items/:itemID/delete", server.listAPI.DeleteItem)
	lists.PUT("/:id/fav", server.listAPI.FavList)
	return server
}

func (server *Server) registerUserRoutes() *Server {
	users := server.router.Group("/users")
	users.GET("/id/:id", server.userAPI.Get)
	users.GET("/email/:email", server.userAPI.GetByEmail)
	users.GET("/username/:username", server.userAPI.GetByUsername)
	users.PUT("/id/:id", server.userAPI.UpdateUser)
	users.POST("/", server.userAPI.CreateUser)

	return server
}
