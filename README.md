# Go Redpanda Streaming API

## Description

This API, built with Go, interacts with Redpanda (a Kafka-compatible streaming platform) to allow users to start streams, send messages, and retrieve stream results. It includes WebSocket support for real-time message delivery and API key authentication.

## Table of Contents

- [Features](#features)
- [Technologies](#technologies)
- [Setup Instructions](#setup-instructions)
- [API Usage](#api-usage)
- [Testing](#testing)
- [Performance Benchmarking](#performance-benchmarking)

## Features

- Start a stream
- Send and retrieve messages in a stream
- WebSocket for real-time messaging
- API key authentication

## Technologies

- Go (Golang)
- Gin (HTTP web framework)
- Sarama (Kafka client)
- Redpanda
- Docker (for running Redpanda)

## Setup Instructions

### Prerequisites

- Go 1.16+
- Docker

### Installation

1. Install dependencies:
   ```bash
   go mod tidy

2. Create a .env file in the root directory with the following content:
KAFKA_BROKER_URL=localhost:9092
KAFKA_TOPIC=streaming
API_KEY=your_api_key
WEBSOCKET_PORT=8080

3. Run Redpanda using Docker:
```
docker run -d --name redpanda -p 9092:9092 vectorized/redpanda:latest redpanda start --overprovisioned --smp 1 --memory 1G --reserve-memory 0M --disable-idle-detection
```
4. Create the topic:
```
docker exec -it redpanda rpk topic create streaming

```
5. Start the Go application:
```
go run main.go
```

API Usage
Start a Stream
Endpoint: POST /stream/start?stream_id={stream_id}

Example:
```
curl -X POST "http://localhost:8080/stream/start?stream_id=test_stream"
```

Send Data to a Stream
Endpoint: POST /stream/send/{stream_id}

Request Body:
```
{
  "payload": "Your message here"
}
```

Example:
```
curl -X POST http://localhost:8080/stream/send/test_stream -H "Content-Type: application/json" -d '{"payload": "Hello, World!"}'
```
Get Results from a Stream
Endpoint: GET /stream/results/{stream_id}

Example:
```
curl -X GET "http://localhost:8080/stream/results/test_stream"
```
Testing
Run Unit Tests
```
go test ./tests -v
```

Run Integration Tests
Ensure that Redpanda is running, then execute:
```
go test ./tests/integration_test.go -v

```



