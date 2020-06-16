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

// AuthUser gets the auth user if any
func (api *API) AuthUser(ctx *gin.Context) user.DTO {
	iss := ctx.GetString("iss")
	if iss != "" {
		user := api.userService.GetByIss(iss)
		if user.ID != 0 {
			return user
		}
	}
	ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"status": "The user can not perform this action"})
	return user.DTO{}
}

// FindAll finds all available lists
func (api *API) FindAll(ctx *gin.Context) {
	lists := api.service.FindAll()
	if len(lists) > 0 {
		ctx.JSON(http.StatusOK, gin.H{"lists": lists})
	} else {
		ctx.JSON(http.StatusNotFound, gin.H{"status": "Not lists found"})
	}
}

// Get a list by id
func (api *API) Get(ctx *gin.Context) {
	id, err := strconv.ParseInt(ctx.Param("id"), 10, 64)
	if err != nil {
		ctx.JSON(
			http.StatusBadRequest,
			gin.H{
				"status": fmt.Sprintf("Not a valid identifier of a list: %s, error: %s", ctx.Param("id"), err.Error()),
			},
		)
	}

	list := api.service.Get(id)
	if list.ID != 0 {
		ctx.JSON(http.StatusOK, gin.H{"list": list})
	} else {
		ctx.JSON(http.StatusNotFound, gin.H{"status": "The list doesn't exist"})
	}
}

// Create a list
func (api *API) Create(ctx *gin.Context) {
	user := api.AuthUser(ctx)
	if !ctx.IsAborted() {
		var command commands.CreateList
		err := ctx.BindJSON(&command)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"status": fmt.Sprintf("Error found in the params: %s", err.Error())})
		}
		list := api.service.Create(user.ID, command)
		ctx.JSON(http.StatusOK, gin.H{"list": list})
	}
}

// Delete a list
func (api *API) Delete(ctx *gin.Context) {
	user := api.AuthUser(ctx)
	if !ctx.IsAborted() {
		id, err := strconv.ParseInt(ctx.Param("id"), 10, 64)
		if err != nil {
			ctx.JSON(
				http.StatusBadRequest,
				gin.H{
					"status": fmt.Sprintf("Not a valid identifier of a list: %s, error: %s", ctx.Param("id"), err.Error()),
				},
			)
		}

		list := api.service.Get(id)
		if list.ID != 0 {
			if list.Owner.ID == user.ID {
				list = api.service.Delete(list.ID)
				ctx.JSON(http.StatusOK, gin.H{"list": list})
			} else {
				ctx.JSON(http.StatusNotFound, gin.H{"status": http.StatusNotFound})
			}
		} else {
			ctx.JSON(http.StatusNotFound, gin.H{"status": http.StatusNotFound})
		}
	}
}

// CreateItem creates a item for a list
func (api *API) CreateItem(ctx *gin.Context) {
	user := api.AuthUser(ctx)
	if !ctx.IsAborted() {
		id, err := strconv.ParseInt(ctx.Param("id"), 10, 64)
		if err != nil {
			ctx.JSON(
				http.StatusBadRequest,
				gin.H{
					"status": fmt.Sprintf("Not a valid identifier of a list: %s, error: %s", ctx.Param("id"), err.Error()),
				},
			)
		}

		var command commands.CreateItem
		err = ctx.BindJSON(&command)
		if err != nil {
			panic(err.Error())
		}

		list := api.service.Get(id)
		if list.ID != 0 {
			if list.Owner.ID == user.ID {
				item := api.service.CreateItem(id, command)
				ctx.JSON(http.StatusOK, gin.H{"item": item})
			} else {
				ctx.JSON(http.StatusUnauthorized, gin.H{"status": "Can not create a item for a list that does not belong to this user"})
			}
		} else {
			ctx.JSON(http.StatusNotFound, gin.H{"status": "List does not exists"})
		}
	}
}

// DeleteItem deletes a item for a list
func (api *API) DeleteItem(ctx *gin.Context) {
	user := api.AuthUser(ctx)
	if !ctx.IsAborted() {
		id, err := strconv.ParseInt(ctx.Param("id"), 10, 64)
		if err != nil {
			ctx.JSON(
				http.StatusBadRequest,
				gin.H{
					"status": fmt.Sprintf("Not a valid identifier of a list: %s, error: %s", ctx.Param("id"), err.Error()),
				},
			)
		}
		list := api.service.Get(id)
		if list.ID != 0 {
			if list.Owner.ID == user.ID {
				itemID, err := strconv.ParseInt(ctx.Param("itemID"), 10, 64)
				if err != nil {
					ctx.JSON(
						http.StatusBadRequest,
						gin.H{
							"status": fmt.Sprintf("Not a valid identifier of a iten: %s, error: %s", ctx.Param("itemID"), err.Error()),
						},
					)
				}
				item := api.service.GetItem(itemID)
				if err != nil {
					panic(err.Error())
				}
				if item.ID != 0 && item.ListID == id {
					item := api.service.DeleteItem(itemID)
					ctx.JSON(http.StatusOK, gin.H{"item": item})
				} else {
					ctx.JSON(http.StatusNotFound, gin.H{"status": http.StatusNotFound})
				}
			} else {
				ctx.JSON(http.StatusUnauthorized, gin.H{"status": http.StatusUnauthorized})
			}
		} else {
			ctx.JSON(http.StatusNotFound, gin.H{"status": http.StatusNotFound})
		}
	}
}

// Fav a list
func (api *API) Fav(ctx *gin.Context) {
	user := api.AuthUser(ctx)
	if !ctx.IsAborted() {
		id, err := strconv.ParseInt(ctx.Param("id"), 10, 64)
		if err != nil {
			ctx.JSON(
				http.StatusBadRequest,
				gin.H{
					"status": fmt.Sprintf("Not a valid identifier of a list: %s, error: %s", ctx.Param("id"), err.Error()),
				},
			)
		}

		list := api.service.Get(id)
		if list.ID != 0 {
			if user.ID != list.Owner.ID {
				list := api.service.Fav(id, user.ID)
				ctx.JSON(http.StatusOK, gin.H{"list": list})
			} else {
				ctx.JSON(http.StatusNotFound, gin.H{"status": http.StatusNotFound})
			}
		} else {
			ctx.JSON(http.StatusNotFound, gin.H{"status": http.StatusNotFound})
		}
	}
}

// Unfav a list
func (api *API) Unfav(ctx *gin.Context) {
	user := api.AuthUser(ctx)
	if !ctx.IsAborted() {
		id, err := strconv.ParseInt(ctx.Param("id"), 10, 64)
		if err != nil {
			ctx.JSON(
				http.StatusBadRequest,
				gin.H{
					"status": fmt.Sprintf("Not a valid identifier of a list: %s, error: %s", ctx.Param("id"), err.Error()),
				},
			)
		}

		list := api.service.Get(id)
		if list.ID != 0 {
			list := api.service.Unfav(id, user.ID)
			ctx.JSON(http.StatusOK, gin.H{"list": list})
		} else {
			ctx.JSON(http.StatusNotFound, gin.H{"status": http.StatusNotFound})
		}
	}
}
