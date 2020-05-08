package list

import (
	"backend/internal/list/commands"
	"net/http"

	"github.com/gin-gonic/gin"
)

// API exposes an API
type API struct {
	service Service
}

// NewAPI constructor of API
func NewAPI(serv Service) API {
	return API{service: serv}
}

// FindAll finds all available lists
func (api *API) FindAll(ctx *gin.Context) {
	lists := api.service.FindAll()
	if len(lists) > 0 {
		ctx.JSON(http.StatusOK, gin.H{"lists": lists})
	} else {
		ctx.JSON(http.StatusNotFound, gin.H{"status": "Not found"})
	}

}

// Get a list by id
func (api *API) Get(ctx *gin.Context) {
	id := ctx.Param("id")
	list := api.service.Get(id)
	if list.ID != 0 {
		ctx.JSON(http.StatusOK, gin.H{"list": list})
	} else {
		ctx.JSON(http.StatusNotFound, gin.H{"status": "Not found"})
	}
}

// CreateList create a list
func (api *API) CreateList(ctx *gin.Context) {
	var createListCommand commands.CreateList
	err := ctx.BindJSON(&createListCommand)
	if err != nil {
		panic(err.Error())
	}
	userID := createListCommand.UserID
	userDTO := api.service.Get(userID)
	if userDTO.ID != 0 {
		if err != nil {
			panic(err.Error())
		}
		id := api.service.CreateList(userID, createListCommand)
		ctx.JSON(http.StatusOK, gin.H{"id": id})
	} else {
		ctx.JSON(http.StatusNotFound, gin.H{"status": http.StatusNotFound})
	}
}
