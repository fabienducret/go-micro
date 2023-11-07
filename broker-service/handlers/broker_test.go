package handlers_test

import (
	"broker/handlers"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBroker(t *testing.T) {
	// Arrange
	sut := http.HandlerFunc(handlers.Broker)
	rr := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodPost, "/", nil)

	// Act
	sut.ServeHTTP(rr, req)
	var reply struct {
		Message string `json:"message"`
	}
	json.Unmarshal(rr.Body.Bytes(), &reply)

	// Assert
	assert.Equal(t, rr.Code, http.StatusOK)
	assert.Equal(t, reply.Message, "Hit the broker")
}
