package main

import (
	"backend/internal/config"
	"backend/internal/list"
	"backend/internal/server"
	"backend/internal/user"
)

func main() {
	conf := config.New()
	listRepository := list.NewRepository(conf.DB)
	listService := list.NewService(listRepository)
	listAPI := list.NewAPI(listService)
	userRepository := user.NewRepository(conf.DB)
	userService := user.NewService(userRepository)
	userAPI := user.NewAPI(userService)

	server.Init(listAPI, userAPI).Run()
}
