package handlers_test

import (
	"broker/handlers"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestBroker(t *testing.T) {
	// Arrange
	h := http.HandlerFunc(handlers.Broker)
	rr := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodPost, "/", nil)

	// Act
	h.ServeHTTP(rr, req)
	message := messageFrom(rr)

	// Assert
	assertStatusCode(t, rr.Code, http.StatusOK)
	assertEqual(t, message, "Hit the broker")

}
