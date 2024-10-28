package controller

import (
	"go-redpanda-streaming/domain"
	"go-redpanda-streaming/usecase"
	"go-redpanda-streaming/utils"

	"net/http"

	"go-redpanda-streaming/repository"

	"github.com/gin-gonic/gin"
)

type StreamController struct {
    usecase *usecase.StreamUsecase
    apiKeyStore domain.APIKeyStore
}

func NewStreamController(usecase *usecase.StreamUsecase, apiKeyStore domain.APIKeyStore) *StreamController {
    return &StreamController{
        usecase: usecase,
        apiKeyStore: apiKeyStore,
    }
}

func (c *StreamController) StartStream(ctx *gin.Context) {
    streamID := ctx.Query("stream_id")
    
    if err := c.usecase.StartStream(streamID); err != nil {
        ctx.JSON(500, gin.H{"error": err.Error()})
        return
    }

    ctx.JSON(http.StatusOK, gin.H{"status": "Stream started"})
}

func (c *StreamController) SendData(ctx *gin.Context) {
    streamID := ctx.Param("stream_id")
    var data domain.StreamData
    if err := ctx.ShouldBindJSON(&data); err != nil {
        ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }
    if err := c.usecase.SendData(streamID, data); err != nil {
        ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }
    ctx.JSON(http.StatusOK, gin.H{"status": "Data sent"})
}

func (c *StreamController) GetResults(ctx *gin.Context) {
    streamID := ctx.Param("stream_id")
    messages, err := c.usecase.GetResults(streamID)
    if err != nil {
        ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    // Return the messages as a JSON response
    ctx.JSON(http.StatusOK, messages)
}

func (c *StreamController) HandleWebSocket(ctx *gin.Context) {
    streamID := ctx.Param("stream_id")
    // Check if the user has access to the stream_id
    if !c.hasAccessToStream(ctx, streamID) {
        ctx.JSON(http.StatusForbidden, gin.H{"error": "Access to this stream is forbidden"})
        return
    }
        
	repository.Hub.HandleConnections(ctx)
}

// hasAccessToStream checks if the user has access to the specified stream_id
func (c *StreamController) hasAccessToStream(ctx *gin.Context, streamID string) bool {
    apiKey := ctx.GetHeader("X-API-Key")

    // Retrieve the stream ID associated with the API key
    if allowedStreamID, exists := c.apiKeyStore.GetStreamID(apiKey); exists {
        return allowedStreamID == streamID
    }

    return false
}

func (c *StreamController)GenerateAPIKey(ctx *gin.Context) {
    streamID := ctx.Param("stream_id")

    apiKey := utils.GenerateAPIKey(streamID) 
    c.apiKeyStore.AddAPIKey(apiKey, streamID)

    ctx.JSON(http.StatusOK, gin.H{"apiKey": apiKey})
}