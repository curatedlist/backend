package list

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// API exposes an API
type API struct {
	Service Service
}

// NewAPI constructor of API
func NewAPI(serv Service) API {
	return API{Service: serv}
}

// FindAll finds all available lists
func (api *API) FindAll(ctx *gin.Context) {
	lists := api.Service.FindAll()
	ctx.JSON(http.StatusOK, gin.H{"lists": lists})
}
