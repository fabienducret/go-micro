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
	payload         string
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
			desc:            "authenticate message with success",
			payload:         authPayloadWithValidPassword,
			expectedCode:    http.StatusAccepted,
			expectedMessage: "Authenticated !",
		},
		{
			desc:            "authenticate message with invalid password",
			payload:         authPayloadWithInvalidPassword,
			expectedCode:    http.StatusUnauthorized,
			expectedMessage: "invalid password",
		},
		{
			desc:            "log message with success",
			payload:         logPayload,
			expectedCode:    http.StatusAccepted,
			expectedMessage: "Log handled for:event",
		},
		{
			desc:            "logger message with error",
			payload:         logPayload,
			logger:          tests.LoggerStub{WithError: true},
			expectedCode:    http.StatusInternalServerError,
			expectedMessage: "server error on logger",
		},
		{
			desc:            "mail message with success",
			payload:         mailPayload,
			expectedCode:    http.StatusAccepted,
			expectedMessage: "Message sent to simpson@gmail.com",
		},
		{
			desc:            "mail message with error",
			payload:         mailPayload,
			mailer:          tests.MailerStub{WithError: true},
			expectedCode:    http.StatusInternalServerError,
			expectedMessage: "server error on mail",
		},
	}

	for _, s := range scenarios {
		t.Run(s.desc, func(t *testing.T) {
			// Arrange
			sut := handlerFrom(s)
			rr := httptest.NewRecorder()
			request, _ := http.NewRequest(http.MethodPost, url, strings.NewReader(s.payload))

			// Act
			sut.ServeHTTP(rr, request)

			// Assert
			assertReply(t, rr, s)
		})
	}
}

func handlerFrom(s scenario) http.HandlerFunc {
	return http.HandlerFunc(
		handlers.MessageFactory(
			tests.AuthenticationStub{},
			loggerStubFrom(s),
			mailerStubFrom(s),
		))
}

func loggerStubFrom(s scenario) handlers.Logger {
	if s.logger != nil {
		return s.logger
	}

	return tests.LoggerStub{}
}

func mailerStubFrom(s scenario) handlers.Mailer {
	if s.mailer != nil {
		return s.mailer
	}

	return tests.MailerStub{}
}

func assertReply(t *testing.T, rr *httptest.ResponseRecorder, s scenario) {
	var reply struct {
		Message string `json:"message"`
	}
	json.Unmarshal(rr.Body.Bytes(), &reply)

	assert.Equal(t, rr.Code, s.expectedCode)
	assert.Equal(t, reply.Message, s.expectedMessage)
}
