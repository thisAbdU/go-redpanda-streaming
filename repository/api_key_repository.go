package repository

import "sync"

// APIKeyRepository implements the APIKeyStore interface
type APIKeyRepository struct {
    mu      sync.RWMutex
    apiKeys map[string]string // map[apiKey]streamID
}

// NewAPIKeyRepository initializes a new API key repository
func NewAPIKeyRepository() *APIKeyRepository {
    return &APIKeyRepository{
        apiKeys: make(map[string]string),
    }
}

// AddAPIKey adds a new API key for a given stream ID
func (r *APIKeyRepository) AddAPIKey(apiKey, streamID string) {
    r.mu.Lock()
    defer r.mu.Unlock()
    r.apiKeys[apiKey] = streamID
}

// GetStreamID retrieves the stream ID associated with the given API key
func (r *APIKeyRepository) GetStreamID(apiKey string) (string, bool) {
    r.mu.RLock()
    defer r.mu.RUnlock()
    streamID, exists := r.apiKeys[apiKey]
    return streamID, exists
}

// IsValidAPIKey checks if the provided API key is valid
func (r *APIKeyRepository) IsValidAPIKey(apiKey string) bool {
    r.mu.RLock()
    defer r.mu.RUnlock()
    _, exists := r.apiKeys[apiKey]
    return exists
}