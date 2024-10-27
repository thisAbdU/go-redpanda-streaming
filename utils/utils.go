package utils

import (
    "crypto/sha256"
    "encoding/hex"
    "fmt"
    "github.com/google/uuid"
)

// GenerateAPIKey generates a unique API key based on the stream ID
func GenerateAPIKey(streamID string) string {
    // Generate a UUID for uniqueness
    uuid := uuid.New().String()
    
    // Combine the stream ID and UUID, then hash them
    hash := sha256.New()
    hash.Write([]byte(fmt.Sprintf("%s-%s", streamID, uuid)))
    
    // Return the hex representation of the hash
    return hex.EncodeToString(hash.Sum(nil))
}