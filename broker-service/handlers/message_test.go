package handlers_test

import (
	"broker/adapters/tests"
	"broker/handlers"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

type scenario struct {
	desc            string
	inRequest       func() *http.Request
	logger          handlers.Logger
	mailer          handlers.Mailer
	expectedCode    int
	expectedMessage string
}

const authPayloadWithValidPassword = "{\"action\":\"auth\",\"auth\":{\"email\":\"admin@example.com\",\"password\":\"verysecret\"}}"
const authPayloadWithInvalidPassword = "{\"action\":\"auth\",\"auth\":{\"email\":\"admin@example.com\",\"password\":\"badpassword\"}}"
const logPayload = "{\"action\":\"log\",\"log\":{\"name\":\"event\",\"data\":\"hello world\"}}"
const mailPayload = "{\"action\":\"mail\",\"mail\":{\"from\":\"homer@gmail.com\",\"to\":\"simpson@gmail.com\"}}"
const url = "/message"

func TestMessage(t *testing.T) {
	scenarios := []scenario{
		{
			desc: "handle authenticate with success",
			inRequest: func() *http.Request {
				request, _ := http.NewRequest(http.MethodPost, url, strings.NewReader(authPayloadWithValidPassword))
				return request
			},
			expectedCode:    http.StatusAccepted,
			expectedMessage: "Authenticated !",
		},
		{
			desc: "handle authenticate with error",
			inRequest: func() *http.Request {
				request, _ := http.NewRequest(http.MethodPost, url, strings.NewReader(authPayloadWithInvalidPassword))
				return request
			},
			expectedCode:    http.StatusUnauthorized,
			expectedMessage: "invalid password",
		},
		{
			desc: "handle logger with success",
			inRequest: func() *http.Request {
				request, _ := http.NewRequest(http.MethodPost, url, strings.NewReader(logPayload))
				return request
			},
			expectedCode:    http.StatusAccepted,
			expectedMessage: "Log handled for:event",
		},
		{
			desc: "handle logger with error",
			inRequest: func() *http.Request {
				request, _ := http.NewRequest(http.MethodPost, url, strings.NewReader(logPayload))
				return request
			},
			logger:          tests.LoggerStub{WithError: true},
			expectedCode:    http.StatusInternalServerError,
			expectedMessage: "server error on logger",
		},
		{
			desc: "handle mail with success",
			inRequest: func() *http.Request {
				request, _ := http.NewRequest(http.MethodPost, url, strings.NewReader(mailPayload))
				return request
			},
			expectedCode:    http.StatusAccepted,
			expectedMessage: "Message sent to simpson@gmail.com",
		},
		{
			desc: "handle mail with error",
			inRequest: func() *http.Request {
				request, _ := http.NewRequest(http.MethodPost, url, strings.NewReader(mailPayload))
				return request
			},
			mailer:          tests.MailerStub{WithError: true},
			expectedCode:    http.StatusInternalServerError,
			expectedMessage: "server error on mail",
		},
	}

	for _, s := range scenarios {
		t.Run(s.desc, func(t *testing.T) {
			// Arrange
			h := http.HandlerFunc(
				handlers.MessageFactory(
					tests.AuthenticationStub{},
					loggerStubFrom(s),
					mailerStubFrom(s),
				))

			// Act
			rr := httptest.NewRecorder()
			h.ServeHTTP(rr, s.inRequest())
			var reply struct {
				Message string `json:"message"`
			}
			json.Unmarshal(rr.Body.Bytes(), &reply)

			// Assert
			assert.Equal(t, rr.Code, s.expectedCode)
			assert.Equal(t, reply.Message, s.expectedMessage)
		})
	}
}

func loggerStubFrom(scenario scenario) handlers.Logger {
	if scenario.logger != nil {
		return scenario.logger
	}

	return tests.LoggerStub{}
}

func mailerStubFrom(scenario scenario) handlers.Mailer {
	if scenario.mailer != nil {
		return scenario.mailer
	}

	return tests.MailerStub{}
}
