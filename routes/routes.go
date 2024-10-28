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

	router.POST("/generate-api-key/:stream_id", streamController.GenerateAPIKey)

	router.POST("/stream/start", streamController.StartStream)

	router.POST("/stream/send/:stream_id", streamController.SendData)

    // router.GET("/stream/results/:stream_id", streamController.GetResults)
	// Use the API key authentication middleware
	router.Use(middleware.APIKeyAuthMiddleware(apiKeyStore))

	// Define your routes here
	router.GET("/ws/:stream_id", streamController.HandleWebSocket)

	return router
}
