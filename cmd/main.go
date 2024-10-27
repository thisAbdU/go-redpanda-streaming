package main

import (
	"fmt"
	"go-redpanda-streaming/config"
	"go-redpanda-streaming/controller"
	"go-redpanda-streaming/repository"
	"go-redpanda-streaming/routes"
	"go-redpanda-streaming/usecase"
	"go-redpanda-streaming/utils"

	"github.com/sirupsen/logrus"
)

var log = logrus.New()

func main() {
    // Load config
    config, err := config.LoadConfig()
    if err != nil {
        log.Fatalf("Error loading config: %v", err)
    }

    // Setup repositories
    kafkaRepo := repository.NewKafkaRepository(config.Kafka.Brokers)

    // Setup use cases
    streamUsecase := usecase.NewStreamUsecase(kafkaRepo)

    apiKeyStore := repository.NewAPIKeyRepository()

    // Generate and store API keys for streams
    streamID1 := "stream_id_1"
    apiKey1 := utils.GenerateAPIKey(streamID1)
    fmt.Println("apie kyye", apiKey1)
    apiKeyStore.AddAPIKey(apiKey1, streamID1)

    streamID2 := "stream_id_2"
    apiKey2 := utils.GenerateAPIKey(streamID2)
    apiKeyStore.AddAPIKey(apiKey2, streamID2)

    // Setup controller
    streamController := controller.NewStreamController(streamUsecase, apiKeyStore)

    // Start server
    router := routes.SetupRouter(streamController, log, apiKeyStore)
    log.Fatal(router.Run(":" + config.WebSocketPort)) // Use the port from config
}