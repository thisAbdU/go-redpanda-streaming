package usecase

import "go-redpanda-streaming/domain"

type StreamUsecase struct {
    repo domain.StreamRepository
}

func NewStreamUsecase(repo domain.StreamRepository) *StreamUsecase {
    return &StreamUsecase{repo: repo}
}

func (u *StreamUsecase) StartStream(streamID string) error {
    return u.repo.StartStream(streamID)
}

func (u *StreamUsecase) SendMessage(streamID string, message domain.Message) error {
    return u.repo.SendMessage(streamID, message)
}

func (u *StreamUsecase) ReceiveMessages(streamID string) (<-chan domain.Message, error) {
    return u.repo.ReceiveMessages(streamID)
}
