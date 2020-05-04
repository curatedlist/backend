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
	if len(lists) > 0 {
		ctx.JSON(http.StatusOK, gin.H{"lists": lists})
	} else {
		ctx.JSON(http.StatusNotFound, gin.H{"status": "Not found"})
	}

}

// Get a list by id
func (api *API) Get(ctx *gin.Context) {
	id := ctx.Param("id")
	list := api.Service.Get(id)
	if list.ID != 0 {
		ctx.JSON(http.StatusOK, gin.H{"list": list})
	} else {
		ctx.JSON(http.StatusNotFound, gin.H{"status": "Not found"})
	}
}
