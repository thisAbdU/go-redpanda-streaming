package domain

type Stream struct {
    StreamID string
    Data     string
}

type Message struct {
    StreamID string
    Payload  string
}

type StreamData struct {
    Data string `json:"data"`
}
