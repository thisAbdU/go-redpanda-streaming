package main

import (
    "log"
    "go-redpanda-streaming/config"
    "go-redpanda-streaming/controller"
    "go-redpanda-streaming/repository"
    "go-redpanda-streaming/usecase"
)

func main() {
    // Load config
    config, err := config.LoadConfig()
    if err != nil {
        log.Fatalf("Error loading config: %v", err)
    }

    // Setup repositories
    kafkaRepo := repository.NewKafkaRepository(config.Kafka)

    // Setup use cases
    streamUsecase := usecase.NewStreamUsecase(kafkaRepo)

    // Setup controller
    streamController := controller.NewStreamController(streamUsecase)

    // Start server
    router := setupRouter(streamController)
    log.Fatal(router.Run(":8080"))
}
