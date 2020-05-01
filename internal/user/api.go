package user

import (
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
	ctx.JSON(http.StatusOK, gin.H{"user": user})
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
