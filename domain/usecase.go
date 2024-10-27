package domain

type StreamUsecase interface {
    StartStream(streamID string) error
    SendData(streamID string, data StreamData) error
    GetResults(streamID string) ([]Message, error)
}