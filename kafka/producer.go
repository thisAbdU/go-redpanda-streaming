package kafka

// import (
// 	"context"
// 	"fmt"
// 	"log"
// 	"sync"

// 	"github.com/Shopify/sarama"
// 	"go-redpanda-streaming/domain"
// )

// type KafkaRepository struct {
// 	producer sarama.SyncProducer
// 	readers  map[string]*sarama.Consumer
// 	mu       sync.Mutex
// }

// func NewKafkaRepository(brokers []string) *KafkaRepository {
// 	config := sarama.NewConfig()
// 	config.Producer.RequiredAcks = sarama.WaitForAll
// 	config.Producer.Partitioner = sarama.NewRandomPartitioner
// 	config.Producer.Return.Successes = true

// 	producer, err := sarama.NewSyncProducer(brokers, config)
// 	if err != nil {
// 		log.Fatalf("Error creating producer: %v", err)
// 	}

// 	return &KafkaRepository{
// 		producer: producer,
// 		readers:  make(map[string]*sarama.Consumer),
// 	}
// }

// func (r *KafkaRepository) StartStream(streamID string) error {
// 	// Validate the streamID to ensure it's a valid topic name
// 	if !domain.IsValidTopicName(streamID) {
// 		return fmt.Errorf("invalid topic name: %s", streamID)
// 	}

// 	// Create the topic if it doesn't exist
// 	topic := "streaming"
// 	err := r.createTopicIfNotExists(topic)
// 	if err != nil {
// 		return fmt.Errorf("failed to create topic %s: %v", topic, err)
// 	}

// 	log.Printf("Stream %s started successfully", streamID)
// 	return nil
// }

// func (r *KafkaRepository) createTopicIfNotExists(topic string) error {
// 	// Check if the topic already exists
// 	adminClient, err := sarama.NewClusterAdmin([]string{"localhost:9092"}, nil)
// 	if err != nil {
// 		return fmt.Errorf("failed to create admin client: %v", err)
// 	}
// 	defer adminClient.Close()

// 	// Check if the topic exists
// 	topics, err := adminClient.ListTopics()
// 	if err != nil {
// 		return fmt.Errorf("failed to list topics: %v", err)
// 	}

// 	if _, exists := topics[topic]; !exists {
// 		// Create the topic with 1 partition and a replication factor of 1
// 		err = adminClient.CreateTopic(topic, &sarama.TopicDetail{
// 			NumPartitions:     1,
// 			ReplicationFactor: 1,
// 		}, false)
// 		if err != nil {
// 			return fmt.Errorf("failed to create topic %s: %v", topic, err)
// 		}
// 		log.Printf("Topic %s created successfully", topic)
// 	} else {
// 		log.Printf("Topic %s already exists", topic)
// 	}

// 	return nil
// }

// func (r *KafkaRepository) SendMessage(streamID string, message domain.Message) error {
// 	msg := &sarama.ProducerMessage{
// 		Topic: streamID,
// 		Key:   sarama.StringEncoder(message.StreamID), // Use StreamID as the key
// 		Value: sarama.StringEncoder(message.Payload),   // Use Payload as the message value
// 	}

// 	_, _, err := r.producer.SendMessage(msg)
// 	if err != nil {
// 		log.Printf("Failed to send message to stream %s: %v", streamID, err)
// 		return err
// 	}

// 	log.Printf("Message sent to stream %s: %s", streamID, message.Payload)
// 	return nil
// }

// func (r *KafkaRepository) Close() {
// 	if err := r.producer.Close(); err != nil {
// 		log.Fatalf("Error closing producer: %v", err)
// 	}
// }