package server_test

import (
	"broker/adapters/tests"
	"broker/server"
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/go-chi/chi/v5"
)

type replyPayload struct {
	Error   bool   `json:"error"`
	Message string `json:"message"`
}

func TestServer(t *testing.T) {
	s := server.NewServer(
		tests.AuthenticationStub{},
		tests.LoggerStub{},
		tests.MailerStub{},
	)

	t.Run("handle hit", func(t *testing.T) {
		// Given
		request, _ := http.NewRequest(http.MethodPost, "/", nil)

		// When
		httpResponse := reponseFrom(s.Routes(), request)
		reply := replyFrom(httpResponse)

		// Then
		assertStatusCode(t, httpResponse.Code, http.StatusOK)
		assertMessage(t, reply.Message, "Hit the broker")
	})

	t.Run("handle authenticate with success", func(t *testing.T) {
		// Given
		payload := "{\"action\":\"auth\",\"auth\":{\"email\":\"admin@example.com\",\"password\":\"verysecret\"}}"
		request, _ := http.NewRequest(http.MethodPost, "/handle", strings.NewReader(payload))

		// When
		httpResponse := reponseFrom(s.Routes(), request)
		reply := replyFrom(httpResponse)

		// Then
		assertStatusCode(t, httpResponse.Code, http.StatusAccepted)
		assertMessage(t, reply.Message, "Authenticated !")
	})

	t.Run("handle authenticate with error", func(t *testing.T) {
		// Given
		payload := "{\"action\":\"auth\",\"auth\":{\"email\":\"admin@example.com\",\"password\":\"badpassword\"}}"
		request, _ := http.NewRequest(http.MethodPost, "/handle", strings.NewReader(payload))

		// When
		httpResponse := reponseFrom(s.Routes(), request)

		// Then
		assertStatusCode(t, httpResponse.Code, http.StatusUnauthorized)
	})

	t.Run("handle logger with success", func(t *testing.T) {
		// Given
		payload := "{\"action\":\"log\",\"log\":{\"name\":\"event\",\"data\":\"hello world\"}}"
		request, _ := http.NewRequest(http.MethodPost, "/handle", strings.NewReader(payload))

		// When
		httpResponse := reponseFrom(s.Routes(), request)
		reply := replyFrom(httpResponse)

		// Then
		assertStatusCode(t, httpResponse.Code, http.StatusAccepted)
		assertMessage(t, reply.Message, "Log handled for:event")
	})

	t.Run("handle logger with error", func(t *testing.T) {
		// Given
		payload := "{\"action\":\"log\",\"log\":{\"name\":\"event\",\"data\":\"hello world\"}}"
		request, _ := http.NewRequest(http.MethodPost, "/handle", strings.NewReader(payload))
		s := server.NewServer(
			tests.AuthenticationStub{},
			tests.LoggerStub{Error: errors.New("")},
			tests.MailerStub{},
		)

		// When
		httpResponse := reponseFrom(s.Routes(), request)
		reply := replyFrom(httpResponse)

		// Then
		assertStatusCode(t, httpResponse.Code, http.StatusInternalServerError)
		assertMessage(t, reply.Message, "server error on logger")
	})

	t.Run("handle mail with success", func(t *testing.T) {
		// Given
		payload := "{\"action\":\"mail\",\"mail\":{\"from\":\"homer@gmail.com\",\"to\":\"simpson@gmail.com\"}}"
		request, _ := http.NewRequest(http.MethodPost, "/handle", strings.NewReader(payload))

		// When
		httpResponse := reponseFrom(s.Routes(), request)
		reply := replyFrom(httpResponse)

		// Then
		assertStatusCode(t, httpResponse.Code, http.StatusAccepted)
		assertMessage(t, reply.Message, "Message sent to simpson@gmail.com")
	})

	t.Run("handle mail with error", func(t *testing.T) {
		// Given
		payload := "{\"action\":\"mail\",\"mail\":{\"from\":\"homer@gmail.com\",\"to\":\"simpson@gmail.com\"}}"
		request, _ := http.NewRequest(http.MethodPost, "/handle", strings.NewReader(payload))
		s := server.NewServer(
			tests.AuthenticationStub{},
			tests.LoggerStub{},
			tests.MailerStub{Error: errors.New("")},
		)

		// When
		httpResponse := reponseFrom(s.Routes(), request)
		reply := replyFrom(httpResponse)

		// Then
		assertStatusCode(t, httpResponse.Code, http.StatusInternalServerError)
		assertMessage(t, reply.Message, "server error on mail")
	})
}

func reponseFrom(mux *chi.Mux, request *http.Request) *httptest.ResponseRecorder {
	response := httptest.NewRecorder()
	mux.ServeHTTP(response, request)

	return response
}

func replyFrom(response *httptest.ResponseRecorder) replyPayload {
	var reply replyPayload
	json.Unmarshal(response.Body.Bytes(), &reply)

	return reply
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
