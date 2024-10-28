package tests

import (
	"bytes"
	"encoding/json"
	"go-redpanda-streaming/controller"
	"go-redpanda-streaming/domain"
	"go-redpanda-streaming/domain/mocks"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)


func TestStartStreamController(t *testing.T){
	t.Run("success", func(t *testing.T){
		mockUsecase := new(mocks.StreamUsecase)
		mockAPIKeyStore := new(mocks.APIKeyStore)
		mockStreamID := "testID"

		mockUsecase.On("StartStream", mockStreamID).Return(nil)

		controller := controller.NewStreamController(mockUsecase, mockAPIKeyStore)
		router := gin.Default()

		router.GET("/stream/start", controller.StartStream)

		req, err := http.NewRequest("GET", "/stream/start?stream_id=testID", nil)
		if err != nil {
			t.Fatal(err)
		}

		w := httptest.NewRecorder()
        router.ServeHTTP(w, req)

		assert.Equal(t, http.StatusOK, w.Code)
	})
}
func TestSendDataController(t *testing.T) {
    t.Run("success", func(t *testing.T) {
        mockUsecase := new(mocks.StreamUsecase)
        mockAPIKeyStore := new(mocks.APIKeyStore)
        mockStreamID := "testID"
        mockData := domain.StreamData{Data: "testData"}

        mockUsecase.On("SendData", mockStreamID, mockData).Return(nil)

        controller := controller.NewStreamController(mockUsecase, mockAPIKeyStore)
        router := gin.Default()

        router.POST("/stream/send/:stream_id", controller.SendData)

        jsonData, err := json.Marshal(mockData)
        if err != nil {
            t.Fatal(err)
        }

        req, err := http.NewRequest("POST", "/stream/send/"+mockStreamID, bytes.NewBuffer(jsonData))
        if err != nil {
			log.Println("this si the error", err)
            t.Fatal(err)
        }
        req.Header.Set("Content-Type", "application/json")

        w := httptest.NewRecorder()
        router.ServeHTTP(w, req)

        assert.Equal(t, http.StatusOK, w.Code)
    })
}

func TestGetResultsController(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		mockUsecase := new(mocks.StreamUsecase)
		mockAPIKeyStore := new(mocks.APIKeyStore)
		mockStreamID := "testID"
		mockMessages := []domain.Message{
			{StreamID: "testID", Payload: "testPayload"},
		}

		mockUsecase.On("GetResults", mockStreamID).Return(mockMessages, nil)

		controller := controller.NewStreamController(mockUsecase, mockAPIKeyStore)
		router := gin.Default()

		router.GET("/stream/results/:stream_id", controller.GetResults)

		req, err := http.NewRequest("GET", "/stream/results/"+mockStreamID, nil)
		if err != nil {
			t.Fatal(err)
		}

		w := httptest.NewRecorder()
		router.ServeHTTP(w, req)

		assert.Equal(t, http.StatusOK, w.Code)
	})
}