package server_test

import (
	"broker/adapters/tests"
	"broker/server"
	"encoding/json"
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

const authPayloadWithValidPassword = "{\"action\":\"auth\",\"auth\":{\"email\":\"admin@example.com\",\"password\":\"verysecret\"}}"
const authPayloadWithInvalidPassword = "{\"action\":\"auth\",\"auth\":{\"email\":\"admin@example.com\",\"password\":\"badpassword\"}}"
const logPayload = "{\"action\":\"log\",\"log\":{\"name\":\"event\",\"data\":\"hello world\"}}"
const mailPayload = "{\"action\":\"mail\",\"mail\":{\"from\":\"homer@gmail.com\",\"to\":\"simpson@gmail.com\"}}"

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
		assertEqual(t, reply.Message, "Hit the broker")
	})

	t.Run("handle authenticate with success", func(t *testing.T) {
		// Given
		request, _ := http.NewRequest(http.MethodPost, "/handle", strings.NewReader(authPayloadWithValidPassword))

		// When
		httpResponse := reponseFrom(s.Routes(), request)
		reply := replyFrom(httpResponse)

		// Then
		assertStatusCode(t, httpResponse.Code, http.StatusAccepted)
		assertEqual(t, reply.Message, "Authenticated !")
	})

	t.Run("handle authenticate with error", func(t *testing.T) {
		// Given
		request, _ := http.NewRequest(http.MethodPost, "/handle", strings.NewReader(authPayloadWithInvalidPassword))

		// When
		httpResponse := reponseFrom(s.Routes(), request)

		// Then
		assertStatusCode(t, httpResponse.Code, http.StatusUnauthorized)
	})

	t.Run("handle logger with success", func(t *testing.T) {
		// Given
		request, _ := http.NewRequest(http.MethodPost, "/handle", strings.NewReader(logPayload))

		// When
		httpResponse := reponseFrom(s.Routes(), request)
		reply := replyFrom(httpResponse)

		// Then
		assertStatusCode(t, httpResponse.Code, http.StatusAccepted)
		assertEqual(t, reply.Message, "Log handled for:event")
	})

	t.Run("handle logger with error", func(t *testing.T) {
		// Given
		s := server.NewServer(
			tests.AuthenticationStub{},
			tests.LoggerStub{WithError: true},
			tests.MailerStub{},
		)
		request, _ := http.NewRequest(http.MethodPost, "/handle", strings.NewReader(logPayload))

		// When
		httpResponse := reponseFrom(s.Routes(), request)
		reply := replyFrom(httpResponse)

		// Then
		assertStatusCode(t, httpResponse.Code, http.StatusInternalServerError)
		assertEqual(t, reply.Message, "server error on logger")
	})

	t.Run("handle mail with success", func(t *testing.T) {
		// Given
		request, _ := http.NewRequest(http.MethodPost, "/handle", strings.NewReader(mailPayload))

		// When
		httpResponse := reponseFrom(s.Routes(), request)
		reply := replyFrom(httpResponse)

		// Then
		assertStatusCode(t, httpResponse.Code, http.StatusAccepted)
		assertEqual(t, reply.Message, "Message sent to simpson@gmail.com")
	})

	t.Run("handle mail with error", func(t *testing.T) {
		// Given
		s := server.NewServer(
			tests.AuthenticationStub{},
			tests.LoggerStub{},
			tests.MailerStub{WithError: true},
		)
		request, _ := http.NewRequest(http.MethodPost, "/handle", strings.NewReader(mailPayload))

		// When
		httpResponse := reponseFrom(s.Routes(), request)
		reply := replyFrom(httpResponse)

		// Then
		assertStatusCode(t, httpResponse.Code, http.StatusInternalServerError)
		assertEqual(t, reply.Message, "server error on mail")
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

func assertEqual(t *testing.T, got, expected string) {
	if got != expected {
		t.Errorf("Test failed got %s, expected %s", got, expected)
	}
}
