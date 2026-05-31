package main

import (
	"backend/internal/config"
	"backend/internal/database"
	"backend/internal/list"
	"backend/internal/search"
	"backend/internal/server"
	"backend/internal/user"
)

func main() {
	conf := config.New()
	db := database.NewDB(conf.DB)
	userRepository := user.NewRepository(db)
	userService := user.NewService(userRepository)
	userAPI := user.NewAPI(userService)
	listRepository := list.NewRepository(db)
	listService := list.NewService(listRepository)
	listAPI := list.NewAPI(listService, userService)
	searchService := search.NewService(conf.TMDBKey)
	searchAPI := search.NewAPI(searchService)

	server.Init(listAPI, userAPI, searchAPI, conf.GoogleClientID).Run()
}
