package server_test

import (
	"broker/repositories"
	"broker/server"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestServer(t *testing.T) {
	s := server.NewServer(
		repositories.NewAuthenticationTestRepository(),
		repositories.NewLoggerTestRepository(),
		repositories.NewMailerTestRepository(),
	)

	t.Run("handle authenticate with success", func(t *testing.T) {
		payload := "{\"action\":\"auth\",\"auth\":{\"email\":\"admin@example.com\",\"password\":\"verysecret\"}}"

		request, _ := http.NewRequest(http.MethodPost, "/handle", strings.NewReader(payload))
		response := httptest.NewRecorder()

		mux := s.Routes()
		mux.ServeHTTP(response, request)

		if response.Code != http.StatusAccepted {
			t.Error("Test failed for route authenticate")
		}
	})

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

	t.Run("handle logger", func(t *testing.T) {
		payload := "{\"action\":\"log\",\"log\":{\"name\":\"event\",\"data\":\"hello world\"}}"

		request, _ := http.NewRequest(http.MethodPost, "/handle", strings.NewReader(payload))
		response := httptest.NewRecorder()

		mux := s.Routes()
		mux.ServeHTTP(response, request)

		if response.Code != http.StatusAccepted {
			t.Error("Test failed for route logger")
		}
	})

	t.Run("handle mail", func(t *testing.T) {
		payload := "{\"action\":\"mail\",\"mail\":{\"from\":\"toto@gmail.com\",\"to\":\"example@gmail.com\"}}"

		request, _ := http.NewRequest(http.MethodPost, "/handle", strings.NewReader(payload))
		response := httptest.NewRecorder()

		mux := s.Routes()
		mux.ServeHTTP(response, request)

		if response.Code != http.StatusAccepted {
			t.Error("Test failed for route mail")
		}
	})
}
