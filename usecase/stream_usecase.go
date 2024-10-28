package usecase

import (
	"go-redpanda-streaming/domain"
	"go-redpanda-streaming/repository"
)

type StreamUsecase struct {
    repo *repository.KafkaRepository
}

func NewStreamUsecase(kafkaRepo *repository.KafkaRepository) domain.StreamUsecase {
    return &StreamUsecase{repo: kafkaRepo}
}

func (u *StreamUsecase) StartStream(streamID string) error {
    return u.repo.StartStream(streamID)
}

func (u *StreamUsecase) SendData(streamID string, data domain.StreamData) error {
    message := domain.Message{
        StreamID: streamID,
        Payload:  data.Data,
    }
    return u.repo.SendMessage(streamID, message)
}

func (u *StreamUsecase) GetResults(streamID string) ([]domain.Message, error) {
    messages, err := u.repo.ReceiveMessages(streamID)
    if err != nil {
        return nil, err
    }

    // Collect messages into a slice
    var result []domain.Message
    for msg := range messages {
        result = append(result, msg)
    }

    return result, nil
}

