package search

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// API exposes the external-catalog search endpoint.
type API struct {
	service Service
}

// NewAPI constructor of API
func NewAPI(serv Service) API {
	return API{service: serv}
}

// Search queries external catalogs for the `q` query param and returns the
// combined, normalized results. Short/empty queries return an empty list.
func (api *API) Search(ctx *gin.Context) {
	query := ctx.Query("q")
	if len(query) < 2 {
		ctx.JSON(http.StatusOK, gin.H{"results": []Result{}})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"results": api.service.Search(query)})
}
