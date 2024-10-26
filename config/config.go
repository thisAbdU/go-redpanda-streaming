package config

import (
    "github.com/joho/godotenv"
    "os"
)

type Config struct {

    Kafka struct {
        BrokerURL string
        Topic     string
    }

    APIKey string
}

func LoadConfig() (*Config, error) {
    _ = godotenv.Load()
    config := &Config{
        Kafka: struct {
            BrokerURL string
            Topic     string
        }{
            BrokerURL: os.Getenv("KAFKA_BROKER_URL"),
            Topic:     "stream_topic",
        },
        APIKey: os.Getenv("API_KEY"),
    }
    return config, nil
}
