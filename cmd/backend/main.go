package main

import (
	"backend/internal/config"
	"backend/internal/list"
	"backend/internal/server"
)

func main() {
	conf := config.New()
	listRepository := list.NewRepository(conf.DB)
	listService := list.NewService(listRepository)
	listAPI := list.NewAPI(listService)
	server.Init(listAPI).Run()
}
