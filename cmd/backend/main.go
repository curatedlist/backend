package main

import (
	"backend/internal/config"
	"backend/internal/database"
	"backend/internal/list"
	"backend/internal/server"
	"backend/internal/user"
)

func main() {
	conf := config.New()
	db := database.NewDB(conf.DB)
	listRepository := list.NewRepository(db)
	listService := list.NewService(listRepository)
	listAPI := list.NewAPI(listService)
	userRepository := user.NewRepository(db)
	userService := user.NewService(userRepository)
	userAPI := user.NewAPI(userService)

	server.Init(listAPI, userAPI).Run()
}
