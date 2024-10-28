package tests

import (
	"go-redpanda-streaming/controller"
	"go-redpanda-streaming/domain/mocks"
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

func TestSendDataController(t *testing.T){
	t.Run("")
}