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

type scenario struct {
	desc            string
	inRequest       func() *http.Request
	logger          server.Logger
	mailer          server.Mailer
	expectedCode    int
	expectedMessage string
}

const authPayloadWithValidPassword = "{\"action\":\"auth\",\"auth\":{\"email\":\"admin@example.com\",\"password\":\"verysecret\"}}"
const authPayloadWithInvalidPassword = "{\"action\":\"auth\",\"auth\":{\"email\":\"admin@example.com\",\"password\":\"badpassword\"}}"
const logPayload = "{\"action\":\"log\",\"log\":{\"name\":\"event\",\"data\":\"hello world\"}}"
const mailPayload = "{\"action\":\"mail\",\"mail\":{\"from\":\"homer@gmail.com\",\"to\":\"simpson@gmail.com\"}}"

func TestServer(t *testing.T) {
	scenarios := []scenario{
		{
			desc: "handle hit",
			inRequest: func() *http.Request {
				request, _ := http.NewRequest(http.MethodPost, "/", nil)
				return request
			},
			expectedCode:    http.StatusOK,
			expectedMessage: "Hit the broker",
		},
		{
			desc: "handle authenticate with success",
			inRequest: func() *http.Request {
				request, _ := http.NewRequest(http.MethodPost, "/handle", strings.NewReader(authPayloadWithValidPassword))
				return request
			},
			expectedCode:    http.StatusAccepted,
			expectedMessage: "Authenticated !",
		},
		{
			desc: "handle authenticate with error",
			inRequest: func() *http.Request {
				request, _ := http.NewRequest(http.MethodPost, "/handle", strings.NewReader(authPayloadWithInvalidPassword))
				return request
			},
			expectedCode:    http.StatusUnauthorized,
			expectedMessage: "invalid password",
		},
		{
			desc: "handle logger with success",
			inRequest: func() *http.Request {
				request, _ := http.NewRequest(http.MethodPost, "/handle", strings.NewReader(logPayload))
				return request
			},
			expectedCode:    http.StatusAccepted,
			expectedMessage: "Log handled for:event",
		},
		{
			desc: "handle logger with error",
			inRequest: func() *http.Request {
				request, _ := http.NewRequest(http.MethodPost, "/handle", strings.NewReader(logPayload))
				return request
			},
			logger:          tests.LoggerStub{WithError: true},
			expectedCode:    http.StatusInternalServerError,
			expectedMessage: "server error on logger",
		},
		{
			desc: "handle mail with success",
			inRequest: func() *http.Request {
				request, _ := http.NewRequest(http.MethodPost, "/handle", strings.NewReader(mailPayload))
				return request
			},
			expectedCode:    http.StatusAccepted,
			expectedMessage: "Message sent to simpson@gmail.com",
		},
		{
			desc: "handle mail with error",
			inRequest: func() *http.Request {
				request, _ := http.NewRequest(http.MethodPost, "/handle", strings.NewReader(mailPayload))
				return request
			},
			mailer:          tests.MailerStub{WithError: true},
			expectedCode:    http.StatusInternalServerError,
			expectedMessage: "server error on mail",
		},
	}

	for _, scenario := range scenarios {
		// Given
		s := server.NewServer(
			tests.AuthenticationStub{},
			loggerStubFrom(scenario),
			mailerStubFrom(scenario),
		)

		// When
		response := reponseFrom(s.Routes(), scenario.inRequest())
		message := messageFrom(response)

		// Then
		assertStatusCode(t, response.Code, scenario.expectedCode)
		assertEqual(t, message, scenario.expectedMessage)
	}
}

func loggerStubFrom(scenario scenario) server.Logger {
	if scenario.logger != nil {
		return scenario.logger
	}

	return tests.LoggerStub{}
}

func mailerStubFrom(scenario scenario) server.Mailer {
	if scenario.mailer != nil {
		return scenario.mailer
	}

	return tests.MailerStub{}
}

func reponseFrom(mux *chi.Mux, request *http.Request) *httptest.ResponseRecorder {
	response := httptest.NewRecorder()
	mux.ServeHTTP(response, request)

	return response
}

func messageFrom(response *httptest.ResponseRecorder) string {
	var reply struct {
		Error   bool   `json:"error"`
		Message string `json:"message"`
	}
	json.Unmarshal(response.Body.Bytes(), &reply)

	return reply.Message
}

func assertStatusCode(t *testing.T, got, want int) {
	if got != want {
		t.Errorf("Test failed for route with status code, got %v, want %v", got, want)
	}
}

func assertEqual(t *testing.T, got, want string) {
	if got != want {
		t.Errorf("Test failed got %s, want %s", got, want)
	}
}
