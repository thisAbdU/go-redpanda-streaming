// integration_test.go
package tests

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"go-redpanda-streaming/controller"
	"go-redpanda-streaming/repository"
	"go-redpanda-streaming/usecase"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestIntegration(t *testing.T) {
    // Set up the Gin router
    gin.SetMode(gin.TestMode)
    router := gin.Default()

    // Initialize the repository and use case
    kafkaRepo := repository.NewKafkaRepository([]string{"localhost:9092"}) // Ensure Kafka is running
    streamUsecase := usecase.NewStreamUsecase(kafkaRepo)
	mockAPIKeyStore := repository.NewAPIKeyRepository()
    streamController := controller.NewStreamController(streamUsecase, mockAPIKeyStore)

    // Define the routes
    router.POST("/stream/start", streamController.StartStream)
    router.POST("/stream/send/:stream_id", streamController.SendData)
    router.GET("/stream/results/:stream_id", streamController.GetResults)

    // Start a stream
    t.Run("Start Stream", func(t *testing.T) {
        reqStart, _ := http.NewRequest("POST", "/stream/start?stream_id=test_stream", nil)
        wStart := httptest.NewRecorder()
        router.ServeHTTP(wStart, reqStart)
        assert.Equal(t, http.StatusOK, wStart.Code)
    })

    // Send data to the stream
    t.Run("Send Data", func(t *testing.T) {
        reqSend, _ := http.NewRequest("POST", "/stream/send/test_stream", bytes.NewBuffer([]byte(`{"payload": "Hello, World!"}`)))
        reqSend.Header.Set("Content-Type", "application/json")
        wSend := httptest.NewRecorder()
        router.ServeHTTP(wSend, reqSend)
        assert.Equal(t, http.StatusOK, wSend.Code)

		time.Sleep(1 * time.Second)
    })
}