package domain

type Stream struct {
    StreamID string
    Data     string
}

type Message struct {
    StreamID string
    Payload  string
}

// repository.go
type StreamRepository interface {
    StartStream(streamID string) error
    SendMessage(streamID string, message Message) error
    ReceiveMessages(streamID string) (<-chan Message, error)
}
