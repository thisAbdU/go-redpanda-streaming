package config

import (
    "github.com/joho/godotenv"
    "os"
    "strings"
)

type Config struct {
    Kafka struct {
        Brokers []string
        Topic   string
    }
    APIKey string
    WebSocketPort string
}

func LoadConfig() (*Config, error) {
    _ = godotenv.Load()
    config := &Config{
        Kafka: struct {
            Brokers []string
            Topic   string
        }{
            Brokers: strings.Split(os.Getenv("KAFKA_BROKER_URL"), ","),
            Topic:   os.Getenv("KAFKA_TOPIC"),
        },
        APIKey: os.Getenv("API_KEY"),
        WebSocketPort: os.Getenv("WEBSOCKET_PORT"),
    }
    return config, nil
}