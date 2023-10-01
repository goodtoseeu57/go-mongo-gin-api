package authMiddleware

import (
	"context"
	"my-api/module/config"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

func FirebaseAuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.Request.Header.Get("Authorization")
		idToken := strings.Split(authHeader, " ")[1]

		token, err := config.AuthClient.VerifyIDToken(context.Background(), idToken)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "error verifying ID token"})
			c.Abort()
			return
		}

		c.Set("token", token.UID)
		c.Next()
	}
}
