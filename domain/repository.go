package domain

type StreamRepository interface {
    StartStream(streamID string) error
    SendMessage(streamID string, message Message) error
    ReceiveMessages(streamID string) (<-chan Message, error)
}

type APIKeyStore interface {
    AddAPIKey(apiKey, streamID string)
    GetStreamID(apiKey string) (string, bool)
    IsValidAPIKey(apiKey string) bool
}