package utils

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"regexp"

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


// isValidTopicName checks if the topic name is valid
func IsValidTopicName(topic string) bool {
    // Define a regex pattern for valid topic names
    validTopicNamePattern := `^[a-zA-Z0-9._-]+$`
    matched, _ := regexp.MatchString(validTopicNamePattern, topic)
    return matched
}