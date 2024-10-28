package tests

import (
	"errors"
	"go-redpanda-streaming/domain"
	"go-redpanda-streaming/domain/mocks"
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
)


func TestStartStream(t *testing.T)  {
	mockrepository := new(mocks.StreamRepository)

	t.Run("success", func(t *testing.T) {
		mockrepository.On("StartStream", "test").Return(nil)
		err := mockrepository.StartStream("test")
		assert.NoError(t, err)
		mockrepository.AssertExpectations(t)
	})
}

func TestSendData(t *testing.T)  {
	mockrepository := new(mocks.StreamRepository)

	t.Run("success", func(t *testing.T) {
		mockStreamIdD := "test"
		mockPayload := "test"
		mockData := domain.Message{StreamID: mockStreamIdD, Payload: mockPayload }

		mockrepository.On("SendMessage", mockStreamIdD, mockData).Return(nil)
		err := mockrepository.SendMessage("test", mockData) 
		assert.NoError(t, err)
		mockrepository.AssertExpectations(t)
	})
}
func TestGetResults(t *testing.T) {
    mockRepo := new(mocks.StreamRepository)

    t.Run("success", func(t *testing.T) {
        mockStreamID := "testMockID"

        // Create a bidirectional channel and populate it
        messagesChannel := make(chan domain.Message, 2)
        mockMessages := []domain.Message{
            {StreamID: "msg1", Payload: "Hello"},
            {StreamID: "msg2", Payload: "World"},
        }

        // Send messages to the channel and close it
        for _, message := range mockMessages {
            messagesChannel <- message
        }
        close(messagesChannel)

        // Convert the channel to a receive-only channel
        receiveOnlyMessagesChannel := (<-chan domain.Message)(messagesChannel)

        // Mock the repository to return the receive-only channel
        mockRepo.On("ReceiveMessages", mockStreamID).Return(receiveOnlyMessagesChannel, nil)

        // Call the GetResults function
        result, err := mockRepo.ReceiveMessages(mockStreamID)

        // Verify no error occurred
        if err != nil {
            t.Fatalf("Expected no error, got %v", err)
        }

        // Collect messages from the channel into a slice
        var receivedMessages []domain.Message
        for msg := range result {
            receivedMessages = append(receivedMessages, msg)
        }

        // Check if the receivedMessages matches the expected mockMessages
        if !reflect.DeepEqual(receivedMessages, mockMessages) {
            t.Errorf("Expected result %v, got %v", mockMessages, receivedMessages)
        }

        // Assert expectations
        mockRepo.AssertExpectations(t)
    })

    t.Run("error from repository", func(t *testing.T) {
        mockStreamID := "errorStreamID"

        expectedError := errors.New("repository error")
        mockRepo.On("ReceiveMessages", mockStreamID).Return(nil, expectedError)

        result, err := mockRepo.ReceiveMessages(mockStreamID)

        if err == nil || err.Error() != expectedError.Error() {
            t.Fatalf("Expected error %v, got %v", expectedError, err)
        }

        if result != nil {
            t.Errorf("Expected result to be nil, got %v", result)
        }

        mockRepo.AssertExpectations(t)
    })
}
