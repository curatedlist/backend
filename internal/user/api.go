package user

import (
	"backend/internal/user/commands"
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

// Get a user by id
func (api *API) Get(ctx *gin.Context) {
	id := ctx.Param("id")
	user := api.service.Get(id)
	if user.ID != 0 {
		ctx.JSON(http.StatusOK, gin.H{"user": user})
	} else {
		ctx.JSON(http.StatusNotFound, gin.H{"status": http.StatusNotFound})
	}
}

// GetByEmail an user by Email
func (api *API) GetByEmail(ctx *gin.Context) {
	email := ctx.Param("email")
	user := api.service.GetByEmail(email)
	if user.ID != 0 {
		ctx.JSON(http.StatusOK, gin.H{"user": user})
	} else {
		ctx.JSON(http.StatusNotFound, gin.H{"status": http.StatusNotFound})
	}
}

// GetByUsername an user by Username
func (api *API) GetByUsername(ctx *gin.Context) {
	username := ctx.Param("username")
	user := api.service.GetByUsername(username)
	if user.ID != 0 {
		ctx.JSON(http.StatusOK, gin.H{"user": user})
	} else {
		ctx.JSON(http.StatusNotFound, gin.H{"status": http.StatusNotFound})
	}
}

// CreateUser create an user
func (api *API) CreateUser(ctx *gin.Context) {
	var registerCommand commands.Register
	err := ctx.BindJSON(&registerCommand)
	if err != nil {
		panic(err.Error())
	}
	id := api.service.CreateUser(registerCommand.Email)
	ctx.JSON(http.StatusOK, gin.H{"id": id})
}

// UpdateUser create an user
func (api *API) UpdateUser(ctx *gin.Context) {
	id := ctx.Param("id")
	var updateCommand commands.Update
	err := ctx.BindJSON(&updateCommand)
	if err != nil {
		panic(err.Error())
	}
	uid := api.service.UpdateUser(id, updateCommand)
	ctx.JSON(http.StatusOK, gin.H{"id": uid})
}
