package middleware

import (
    "github.com/gin-gonic/gin"
    "github.com/sirupsen/logrus"
)

// LoggerMiddleware logs incoming requests and responses
func LoggerMiddleware(log *logrus.Logger) gin.HandlerFunc {
    return func(c *gin.Context) {
        log.WithFields(logrus.Fields{
            "method": c.Request.Method,
            "path":   c.Request.URL.Path,
        }).Info("Incoming request")

        c.Next()

        if len(c.Errors) > 0 {
            for _, err := range c.Errors {
                log.WithFields(logrus.Fields{
                    "error":  err.Error(),
                    "status": c.Writer.Status(),
                }).Error("Error occurred")
            }
        } else {
            log.WithFields(logrus.Fields{
                "status": c.Writer.Status(),
            }).Info("Response sent")
        }
    }
}