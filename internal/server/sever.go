package server

import (
	"backend/internal/list"
	"backend/internal/middleware"
	"backend/internal/user"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	ginprometheus "github.com/zsais/go-gin-prometheus"
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
	return server.withProm().withCors().registerAllRoutes()
}

// Run starts the http server
func (server *Server) Run() *Server {
	err := server.router.Run()
	if err != nil {
		panic(err)
	}
	return server
}

func (server *Server) withProm() *Server {
	prom := ginprometheus.NewPrometheus("gin")
	prom.Use(server.router)
	return server
}

func (server *Server) withCors() *Server {
	config := cors.DefaultConfig()
	config.AllowHeaders = []string{"Origin", "Content-Length", "Content-Type", "Authorization"}
	config.AllowAllOrigins = true
	server.router.Use(cors.New(config))
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

	authenticated := lists.Group("/")
	authenticated.Use(middleware.TokenAuthMiddleware())
	{
		authenticated.POST("/", server.listAPI.Create)
		authenticated.DELETE("/:id", server.listAPI.Delete)
		authenticated.POST("/:id/items/", server.listAPI.CreateItem)
		authenticated.PATCH("/:id/items/:itemID/delete", server.listAPI.DeleteItem)
		authenticated.POST("/:id/fav", server.listAPI.Fav)
		authenticated.DELETE("/:id/unfav", server.listAPI.Unfav)
	}
	return server
}

func (server *Server) registerUserRoutes() *Server {
	users := server.router.Group("/users")
	users.GET("/username/:username", server.userAPI.GetByUsername)
	users.GET("/username/:username/lists", server.userAPI.GetListsByUsername)
	users.GET("/username/:username/favs", server.userAPI.GetFavsByUsername)

	authenticated := users.Group("/")
	authenticated.Use(middleware.TokenAuthMiddleware())
	{
		authenticated.POST("/login", server.userAPI.Login)
		authenticated.PUT("/id/:id", server.userAPI.Update)
		authenticated.POST("/", server.userAPI.Create)
	}
	return server
}
