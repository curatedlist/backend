package middleware

import (
	"context"
	"strings"

	"github.com/gin-gonic/gin"
	"google.golang.org/api/idtoken"
)

// TokenAuthMiddleware validates a Google ID token from the Authorization header.
// On success it stores the Google subject (a stable per-user id, persisted in
// user.iss) and the verified email on the gin context for downstream handlers.
func TokenAuthMiddleware(clientID string) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		authHeader := ctx.Request.Header.Get("Authorization")
		parts := strings.SplitN(authHeader, "Bearer ", 2)
		if len(parts) < 2 || parts[1] == "" {
			ctx.AbortWithStatusJSON(401, gin.H{"status": "API token required"})
			return
		}
		token := parts[1]

		payload, err := idtoken.Validate(context.Background(), token, clientID)
		if err != nil {
			ctx.AbortWithStatusJSON(401, gin.H{"status": "Invalid API token"})
			return
		}

		// Only trust verified emails for the email-based account migration.
		email, _ := payload.Claims["email"].(string)
		if verified, ok := payload.Claims["email_verified"].(bool); !ok || !verified {
			email = ""
		}

		ctx.Set("iss", payload.Subject)
		ctx.Set("email", email)
		ctx.Next()
	}
}
