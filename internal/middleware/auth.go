package middleware

import (
	"strings"

	"github.com/dhighwayman/go-magic"
	"github.com/gin-gonic/gin"
)

//TokenAuthMiddleware middleware for gin to auth token
func TokenAuthMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		req := ctx.Request
		authToken := req.Header.Get("Authorization")
		token := strings.Split(authToken, "Bearer ")[1]
		if token == "" {
			ctx.AbortWithStatusJSON(401, gin.H{"status": "API token required"})
			return
		}

		magic, err := magic.New(token)
		if err != nil {
			ctx.AbortWithStatusJSON(401, gin.H{"status": err.Error()})
			return
		}

		err = magic.Validate()
		if err != nil {
			ctx.AbortWithStatusJSON(401, gin.H{"status": "Invalid API token"})
			return
		}

		iss := magic.Issuer()
		if err != nil {
			ctx.AbortWithStatusJSON(401, gin.H{"status": "Invalid API token"})
			return
		}

		ctx.Set("iss", iss)
		ctx.Next()
	}
}
