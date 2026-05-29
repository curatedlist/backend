package user

import (
	"backend/internal/user/commands"
	"fmt"
	"log"
	"net/http"
	"strconv"

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

// AuthUser gets the auth user if any.
// It resolves by Google subject (stored in user.iss); for users created under
// the old Magic auth it falls back to matching the verified email and re-links
// the account to the Google identity so all their data carries over.
func (api *API) AuthUser(ctx *gin.Context) DTO {
	iss := ctx.GetString("iss")
	email := ctx.GetString("email")
	log.Printf("Authenticated call, iss: %s goes into %s", iss, ctx.Request.RequestURI)
	if iss != "" {
		userDTO := api.service.GetByIss(iss)
		if userDTO.ID != 0 {
			return userDTO
		}
		if email != "" {
			existing := api.service.GetByEmail(email)
			if existing.ID != 0 {
				return api.service.Relink(existing.ID, iss)
			}
		}
	}
	ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"status": "Invalid token"})
	return DTO{}
}

// Get a user by id
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

	user := api.service.Get(id)
	if user.ID != 0 {
		ctx.JSON(http.StatusOK, gin.H{"user": user})
	} else {
		ctx.JSON(http.StatusNotFound, gin.H{"status": http.StatusNotFound})
	}
}

// Login an user. The identity is established entirely from the verified Google
// token (handled by AuthUser); returns 401 for an unknown user so the client
// can fall back to creating one.
func (api *API) Login(ctx *gin.Context) {
	user := api.AuthUser(ctx)
	if !ctx.IsAborted() {
		ctx.JSON(http.StatusOK, gin.H{"user": user})
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

// GetListsByUsername gets list of an user by username
func (api *API) GetListsByUsername(ctx *gin.Context) {
	username := ctx.Param("username")
	user := api.service.GetByUsername(username)
	if user.ID != 0 {
		lists := api.service.GetLists(user)
		ctx.JSON(http.StatusOK, gin.H{"lists": lists})
	} else {
		ctx.JSON(http.StatusNotFound, gin.H{"status": http.StatusNotFound})
	}
}

// GetFavsByUsername gets favorites list of an user by username
func (api *API) GetFavsByUsername(ctx *gin.Context) {
	username := ctx.Param("username")
	user := api.service.GetByUsername(username)
	if user.ID != 0 {
		lists := api.service.GetFavs(user)
		ctx.JSON(http.StatusOK, gin.H{"lists": lists})
	} else {
		ctx.JSON(http.StatusNotFound, gin.H{"status": http.StatusNotFound})
	}
}

// Create an user from the verified Google token (subject + email). The email is
// taken from the token rather than the request body so it can't be spoofed.
func (api *API) Create(ctx *gin.Context) {
	iss := ctx.GetString("iss")
	email := ctx.GetString("email")
	if iss != "" && email != "" {
		user := api.service.Create(email, iss)
		ctx.JSON(http.StatusOK, gin.H{"user": user})
	} else {
		ctx.JSON(http.StatusUnauthorized, gin.H{"status": "Invalid token"})
	}
}

// Update an user
func (api *API) Update(ctx *gin.Context) {
	userDTO := api.AuthUser(ctx)
	if !ctx.IsAborted() {
		var command commands.Update
		err := ctx.BindJSON(&command)
		if err != nil {
			panic(err.Error())
		}
		user := api.service.Update(userDTO.ID, command)
		ctx.JSON(http.StatusOK, gin.H{"user": user})
	}
}
