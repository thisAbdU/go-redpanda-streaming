package middleware

import (
    "github.com/gin-gonic/gin"
    "net/http"
)

func APIKeyAuthMiddleware(apiKey string) gin.HandlerFunc {
    return func(c *gin.Context) {
        if c.GetHeader("X-API-KEY") != apiKey {
            c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "unauthorized"})
            return
        }
        c.Next()
    }
}
