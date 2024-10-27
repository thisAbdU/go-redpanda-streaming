package repository

import (
    "context"
    "go-redpanda-streaming/domain"
    "log"
    "github.com/segmentio/kafka-go"
    "sync"
)

type KafkaRepository struct {
    writer  *kafka.Writer
    readers map[string]*kafka.Reader
    mu      sync.Mutex
}

func NewKafkaRepository(brokers []string) *KafkaRepository {
    writer := kafka.NewWriter(kafka.WriterConfig{
        Brokers: brokers,
    })
    return &KafkaRepository{
        writer:  writer,
        readers: make(map[string]*kafka.Reader),
    }
}

func (r *KafkaRepository) StartStream(streamID string) error {
    conn, err := kafka.DialLeader(context.Background(), "tcp", "localhost:9092", streamID, 0)
    if err != nil {
        log.Printf("Failed to dial leader for stream %s: %v", streamID, err)
        return err
    }
    defer conn.Close()

    // Create topic if it doesn't exist
    err = conn.CreateTopics(kafka.TopicConfig{
        Topic:             streamID,
        NumPartitions:     1,
        ReplicationFactor: 1,
    })
    if err != nil {
        log.Printf("Failed to create topic for stream %s: %v", streamID, err)
        return err
    }

    log.Printf("Stream %s started successfully", streamID)
    return nil
}

func (r *KafkaRepository) SendMessage(streamID string, message domain.Message) error {
    go func() {
        err := r.writer.WriteMessages(context.Background(), kafka.Message{
            Key:   []byte(streamID),
            Value: []byte(message.Payload),
        })
        if err != nil {
            log.Printf("Failed to send message to stream %s: %v", streamID, err)
        } else {
            log.Printf("Message sent to stream %s: %s", streamID, message.Payload)
        }
    }()
    return nil
}

func (r *KafkaRepository) ReceiveMessages(streamID string) (<-chan domain.Message, error) {
    r.mu.Lock()
    reader, exists := r.readers[streamID]
    if !exists {
        reader = kafka.NewReader(kafka.ReaderConfig{
            Brokers: []string{"localhost:9092"},
            Topic:   streamID,
            GroupID: streamID,
        })
        r.readers[streamID] = reader
    }
    r.mu.Unlock()

    ch := make(chan domain.Message, 100) // Buffered channel for high throughput

    go func() {
        defer reader.Close()
        for {
            msg, err := reader.ReadMessage(context.Background())
            if err != nil {
                log.Printf("Failed to read message from stream %s: %v", streamID, err)
                close(ch)
                return
            }
            log.Printf("Message received from stream %s: %s", streamID, string(msg.Value))
            processedMsg := domain.Message{
                StreamID: streamID,
                Payload:  "Processed: " + string(msg.Value),
            }
            ch <- processedMsg

            // Send processed message to WebSocket clients
            Hub.Broadcast([]byte(processedMsg.Payload))
        }
    }()
    return ch, nil
}

func (r *KafkaRepository) GetResults(streamID string) ([]domain.Message, error) {
    // Simulate real-time processing
    messages := []domain.Message{}
    ch, err := r.ReceiveMessages(streamID)
    if err != nil {
        return nil, err
    }

    for msg := range ch {
        // Simulate processing (e.g., transformation or aggregation)
        processedMsg := domain.Message{
            StreamID: msg.StreamID,
            Payload:  "Processed: " + msg.Payload,
        }
        messages = append(messages, processedMsg)
    }
    return messages, nil
}
