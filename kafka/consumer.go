package kafka

import "github.com/IBM/sarama"

type Consumer struct {
    consumer sarama.Consumer
}

func NewConsumer(brokers []string) (*Consumer, error) {
    consumer, err := sarama.NewConsumer(brokers, nil)
    if err != nil {
        return nil, err
    }
    return &Consumer{consumer: consumer}, nil
}

func (c *Consumer) ConsumePartition(topic string, partition int32) (sarama.PartitionConsumer, error) {
    return c.consumer.ConsumePartition(topic, partition, sarama.OffsetNewest)
}