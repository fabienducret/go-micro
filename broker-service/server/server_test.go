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

	t.Run("handle authenticate with error", func(t *testing.T) {
		payload := "{\"action\":\"auth\",\"auth\":{\"email\":\"admin@example.com\",\"password\":\"badpassword\"}}"

		request, _ := http.NewRequest(http.MethodPost, "/handle", strings.NewReader(payload))
		response := httptest.NewRecorder()

		mux := s.Routes()
		mux.ServeHTTP(response, request)

		if response.Code != http.StatusUnauthorized {
			t.Error("Test failed for route authenticate")
		}
	})
	t.Run("handle authenticate with success", func(t *testing.T) {
		payload := "{\"action\":\"auth\",\"auth\":{\"email\":\"admin@example.com\",\"password\":\"verysecret\"}}"

		request, _ := http.NewRequest(http.MethodPost, "/handle", strings.NewReader(payload))
		response := httptest.NewRecorder()

		mux := s.Routes()
		mux.ServeHTTP(response, request)

		var reply replyPayload
		json.Unmarshal(response.Body.Bytes(), &reply)

		if response.Code != http.StatusAccepted {
			t.Error("Test failed for route authenticate, bad status code")
		}

		if reply.Message != "Authenticated !" {
			t.Errorf("Test failed for route authenticate, reply %s", reply.Message)
		}
	})

	t.Run("handle logger", func(t *testing.T) {
		payload := "{\"action\":\"log\",\"log\":{\"name\":\"event\",\"data\":\"hello world\"}}"

		request, _ := http.NewRequest(http.MethodPost, "/handle", strings.NewReader(payload))
		response := httptest.NewRecorder()

		mux := s.Routes()
		mux.ServeHTTP(response, request)

		var reply replyPayload
		json.Unmarshal(response.Body.Bytes(), &reply)

		if response.Code != http.StatusAccepted {
			t.Error("Test failed for route logger, bad status code")
		}

		if reply.Message != "Log handled for:event" {
			t.Errorf("Test failed for route logger, reply %s", reply.Message)
		}
	})

	t.Run("handle mail", func(t *testing.T) {
		payload := "{\"action\":\"mail\",\"mail\":{\"from\":\"homer@gmail.com\",\"to\":\"simpson@gmail.com\"}}"

		request, _ := http.NewRequest(http.MethodPost, "/handle", strings.NewReader(payload))
		response := httptest.NewRecorder()

		mux := s.Routes()
		mux.ServeHTTP(response, request)

		var reply replyPayload
		json.Unmarshal(response.Body.Bytes(), &reply)

		if response.Code != http.StatusAccepted {
			t.Error("Test failed for route mail, bad status code")
		}

		if reply.Message != "Message sent to simpson@gmail.com" {
			t.Errorf("Test failed for route mail, reply %s", reply.Message)
		}
	})
}
