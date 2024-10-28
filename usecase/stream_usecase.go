package usecase

import (
	"go-redpanda-streaming/domain"
	"go-redpanda-streaming/repository"
)

type StreamUsecase struct {
    repo *repository.KafkaRepository
}

// func NewStreamUsecase(repo domain.StreamRepository) *StreamUsecase {
//     return &StreamUsecase{repo: repo}
// }
func NewStreamUsecase(kafkaRepo *repository.KafkaRepository) *StreamUsecase {
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

// func (u *StreamUsecase) GetResults(streamID string) ([]domain.Message, error) {
//     return u.repo.GetResults(streamID)
// }