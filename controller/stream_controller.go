package controller

import (
    "github.com/gin-gonic/gin"
    "go-redpanda-streaming/domain"
    "go-redpanda-streaming/usecase"
)

type StreamController struct {
    usecase *usecase.StreamUsecase
}

func NewStreamController(usecase *usecase.StreamUsecase) *StreamController {
    return &StreamController{usecase: usecase}
}

func (c *StreamController) StartStream(ctx *gin.Context) {
    streamID := ctx.Param("stream_id")
    if err := c.usecase.StartStream(streamID); err != nil {
        ctx.JSON(500, gin.H{"error": err.Error()})
        return
    }
    ctx.JSON(200, gin.H{"status": "Stream started"})
}
