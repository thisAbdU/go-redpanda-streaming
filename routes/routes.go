package routes

import (
	"go-redpanda-streaming/controller"
	"go-redpanda-streaming/domain"
	"go-redpanda-streaming/middleware"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

// setupRouter initializes the router and sets up the routes
func SetupRouter(streamController *controller.StreamController, log *logrus.Logger,  apiKeyStore domain.APIKeyStore) *gin.Engine {
	router := gin.Default()

	// Use the logging middleware
	router.Use(middleware.LoggerMiddleware(log))

	// Use the API key authentication middleware
	router.Use(middleware.APIKeyAuthMiddleware(apiKeyStore))

	// Define your routes here
	router.GET("/ws/:stream_id", streamController.HandleWebSocket)

	return router
}
