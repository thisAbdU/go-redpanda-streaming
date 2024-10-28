package repository

import (
	// "context"
	"fmt"
	"go-redpanda-streaming/domain"
	"go-redpanda-streaming/utils"
	"log"
	"sync"

	"github.com/IBM/sarama"
)

type KafkaRepository struct {
	producer sarama.SyncProducer
	readers  map[string]*sarama.Consumer
	mu       sync.Mutex
}

func NewKafkaRepository(brokers []string) *KafkaRepository {
	config := sarama.NewConfig()
	config.Producer.RequiredAcks = sarama.WaitForAll
	config.Producer.Partitioner = sarama.NewRandomPartitioner
	config.Producer.Return.Successes = true

	producer, err := sarama.NewSyncProducer(brokers, config)
	if err != nil {
		log.Fatalf("Error creating producer: %v", err)
	}

	return &KafkaRepository{
		producer: producer,
		readers:  make(map[string]*sarama.Consumer),
	}
}

func (r *KafkaRepository) StartStream(streamID string) error {
    // Validate the streamID to ensure it's a valid topic name
    
    if !utils.IsValidTopicName(streamID) {
        return fmt.Errorf("invalid topic name: %s", streamID)
    }

	// Create the topic if it doesn't exist
	topic := "streaming"
	err := r.createTopicIfNotExists(topic)
	if err != nil {
		return fmt.Errorf("failed to create topic %s: %v", topic, err)
	}

	log.Printf("Stream %s started successfully", streamID)
	return nil
}

func (r *KafkaRepository) createTopicIfNotExists(topic string) error {
	// Check if the topic already exists
	adminClient, err := sarama.NewClusterAdmin([]string{"localhost:9092"}, nil)
	if err != nil {
		return fmt.Errorf("failed to create admin client: %v", err)
	}
	defer adminClient.Close()

	// Check if the topic exists
	topics, err := adminClient.ListTopics()
	if err != nil {
		return fmt.Errorf("failed to list topics: %v", err)
	}

	if _, exists := topics[topic]; !exists {
		// Create the topic with 1 partition and a replication factor of 1
		err = adminClient.CreateTopic(topic, &sarama.TopicDetail{
			NumPartitions:     1,
			ReplicationFactor: 1,
		}, false)
		if err != nil {
			return fmt.Errorf("failed to create topic %s: %v", topic, err)
		}
		log.Printf("Topic %s created successfully", topic)
	} else {
		log.Printf("Topic %s already exists", topic)
	}

	return nil
}

func (r *KafkaRepository) SendMessage(streamID string, message domain.Message) error {
	msg := &sarama.ProducerMessage{
		Topic: streamID,
		Key:   sarama.StringEncoder(message.StreamID), // Use StreamID as the key
		Value: sarama.StringEncoder(message.Payload),   // Use Payload as the message value
	}

	_, _, err := r.producer.SendMessage(msg)
	if err != nil {
		log.Printf("Failed to send message to stream %s: %v", streamID, err)
		return err
	}

	log.Printf("Message sent to stream %s: %s", streamID, message.Payload)
	return nil
}
func (r *KafkaRepository) ReceiveMessages(streamID string) (<-chan domain.Message, error) {
	r.mu.Lock()
	defer r.mu.Unlock()

	// Create a new consumer if it doesn't exist
	if _, exists := r.readers[streamID]; !exists {
		consumer, err := sarama.NewConsumer([]string{"localhost:9092"}, nil)
		if err != nil {
			return nil, fmt.Errorf("failed to create consumer: %v", err)
		}
		r.readers[streamID] = &consumer // Store the consumer interface directly
	}

	// Create a channel to return messages
	ch := make(chan domain.Message)

	go func() {
		defer close(ch)
		partitionConsumer, err := (*r.readers[streamID]).ConsumePartition("streaming", 0, sarama.OffsetNewest)
		if err != nil {
			log.Printf("Failed to start partition consumer: %v", err)
			return
		}
		defer partitionConsumer.Close()

		for msg := range partitionConsumer.Messages() {
			// Create a domain.Message from the consumed message
			receivedMessage := domain.Message{
				StreamID: string(msg.Key),
				Payload:  string(msg.Value),
			}
			ch <- receivedMessage
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
