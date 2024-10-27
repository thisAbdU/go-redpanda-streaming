package middleware

import (
	"go-redpanda-streaming/domain"
	"net/http"

	"github.com/gin-gonic/gin"
)

// APIKeyAuthMiddleware checks for a valid API key in the request headers
func APIKeyAuthMiddleware(apiKeyStore domain.APIKeyStore) gin.HandlerFunc {
    return func(c *gin.Context) {
        apiKey := c.GetHeader("X-API-Key")
        if !apiKeyStore.IsValidAPIKey(apiKey) {
            c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
            c.Abort()
            return
        }
        c.Next()
    }
}
