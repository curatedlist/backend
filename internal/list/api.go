package list

import (
	"backend/internal/list/commands"
	"backend/internal/user"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// API exposes an API
type API struct {
	service     Service
	userService user.Service
}

// NewAPI constructor of API
func NewAPI(serv Service, userServ user.Service) API {
	return API{service: serv, userService: userServ}
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

// CreateItem create a item for a list
func (api *API) CreateItem(ctx *gin.Context) {
	var createItemCommand commands.CreateItem
	err := ctx.BindJSON(&createItemCommand)
	if err != nil {
		panic(err.Error())
	}
	userID := createItemCommand.UserID
	userDTO := api.userService.Get(userID)
	if userDTO.ID != 0 {
		if err != nil {
			panic(err.Error())
		}
		listID := createItemCommand.ListID
		listDTO := api.service.Get(listID)
		if listDTO.ID != 0 {
			if err != nil {
				panic(err.Error())
			}
			item := api.service.CreateItem(listID, createItemCommand)
			ctx.JSON(http.StatusOK, gin.H{"item": item})
		} else {
			ctx.JSON(http.StatusNotFound, gin.H{"status": http.StatusNotFound})
		}
	} else {
		ctx.JSON(http.StatusNotFound, gin.H{"status": http.StatusNotFound})
	}
}

// DeleteItem deletes a item for a list
func (api *API) DeleteItem(ctx *gin.Context) {
	listID := ctx.Param("listID")
	listDTO := api.service.Get(listID)
	if listDTO.ID != 0 {
		itemID := ctx.Param("itemID")
		itemDTO := api.service.GetItem(itemID)
		fmt.Println(itemDTO.ListID)
		listID, err := strconv.ParseUint(listID, 10, 64)
		if err != nil {
			panic(err.Error())
		}
		if itemDTO.ID != 0 && uint64(itemDTO.ListID) == listID {
			fmt.Println(itemDTO.ID)
			item := api.service.DeleteItem(itemID)
			ctx.JSON(http.StatusOK, gin.H{"item": item})
		} else {
			ctx.JSON(http.StatusNotFound, gin.H{"status": http.StatusNotFound})
		}
	} else {
		ctx.JSON(http.StatusNotFound, gin.H{"status": http.StatusNotFound})
	}
}
