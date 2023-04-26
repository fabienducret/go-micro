package server_test

import (
	"broker/repositories"
	"broker/server"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

type replyPayload struct {
	Error   bool   `json:"error"`
	Message string `json:"message"`
}

func TestServer(t *testing.T) {
	s := server.NewServer(
		repositories.NewAuthenticationTestRepository(),
		repositories.NewLoggerTestRepository(),
		repositories.NewMailerTestRepository(),
	)

	t.Run("handle hit", func(t *testing.T) {
		request, _ := http.NewRequest(http.MethodPost, "/", nil)
		response := httptest.NewRecorder()

		mux := s.Routes()
		mux.ServeHTTP(response, request)

		var reply replyPayload
		json.Unmarshal(response.Body.Bytes(), &reply)

		assertStatusCode(t, response.Code, http.StatusOK)
		assertMessage(t, reply.Message, "Hit the broker")
	})

	t.Run("handle authenticate with error", func(t *testing.T) {
		payload := "{\"action\":\"auth\",\"auth\":{\"email\":\"admin@example.com\",\"password\":\"badpassword\"}}"

		request, _ := http.NewRequest(http.MethodPost, "/handle", strings.NewReader(payload))
		response := httptest.NewRecorder()

		mux := s.Routes()
		mux.ServeHTTP(response, request)

		assertStatusCode(t, response.Code, http.StatusUnauthorized)
	})

	t.Run("handle authenticate with success", func(t *testing.T) {
		payload := "{\"action\":\"auth\",\"auth\":{\"email\":\"admin@example.com\",\"password\":\"verysecret\"}}"

		request, _ := http.NewRequest(http.MethodPost, "/handle", strings.NewReader(payload))
		response := httptest.NewRecorder()

		mux := s.Routes()
		mux.ServeHTTP(response, request)

		var reply replyPayload
		json.Unmarshal(response.Body.Bytes(), &reply)

		assertStatusCode(t, response.Code, http.StatusAccepted)
		assertMessage(t, reply.Message, "Authenticated !")
	})

	t.Run("handle logger", func(t *testing.T) {
		payload := "{\"action\":\"log\",\"log\":{\"name\":\"event\",\"data\":\"hello world\"}}"

		request, _ := http.NewRequest(http.MethodPost, "/handle", strings.NewReader(payload))
		response := httptest.NewRecorder()

		mux := s.Routes()
		mux.ServeHTTP(response, request)

		var reply replyPayload
		json.Unmarshal(response.Body.Bytes(), &reply)

		assertStatusCode(t, response.Code, http.StatusAccepted)
		assertMessage(t, reply.Message, "Log handled for:event")
	})

	t.Run("handle mail", func(t *testing.T) {
		payload := "{\"action\":\"mail\",\"mail\":{\"from\":\"homer@gmail.com\",\"to\":\"simpson@gmail.com\"}}"

		request, _ := http.NewRequest(http.MethodPost, "/handle", strings.NewReader(payload))
		response := httptest.NewRecorder()

		mux := s.Routes()
		mux.ServeHTTP(response, request)

		var reply replyPayload
		json.Unmarshal(response.Body.Bytes(), &reply)

		assertStatusCode(t, response.Code, http.StatusAccepted)
		assertMessage(t, reply.Message, "Message sent to simpson@gmail.com")
	})
}

func assertStatusCode(t *testing.T, got, want int) {
	if got != want {
		t.Errorf("Test failed for route with status code %v", got)
	}
}

func assertMessage(t *testing.T, got, want string) {
	if got != want {
		t.Errorf("Test failed for route with message %s", got)
	}
}
