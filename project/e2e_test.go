package project

import (
	"context"
	"encoding/json"
	"io"
	"net/http"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	tc "github.com/testcontainers/testcontainers-go/modules/compose"
)

type scenario struct {
	desc            string
	inRequest       func() *http.Request
	expectedCode    int
	expectedMessage string
}

const authPayloadWithSuccess = "{\"action\":\"auth\",\"auth\":{\"email\":\"admin@example.com\",\"password\":\"verysecret\"}}"
const authPayloadWithError = "{\"action\":\"auth\",\"auth\":{\"email\":\"admin@example.com\",\"password\":\"badpassword\"}}"
const logPayload = "{\"action\":\"log\",\"log\":{\"name\":\"event\",\"data\":\"hello world\"}}"
const mailPayload = "{\"action\":\"mail\",\"mail\":{\"from\":\"homer@gmail.com\",\"to\":\"simpson@gmail.com\"}}"
const url = "http://localhost:8080/message"

func TestE2E(t *testing.T) {
	prepareCompose(t)

	scenarios := []scenario{
		{
			desc: "should send a log event with success",
			inRequest: func() *http.Request {
				request, _ := http.NewRequest(http.MethodPost, url, strings.NewReader(logPayload))
				return request
			},
			expectedCode:    http.StatusAccepted,
			expectedMessage: "Log handled for:event",
		},
		{
			desc: "should send a mail with success",
			inRequest: func() *http.Request {
				request, _ := http.NewRequest(http.MethodPost, url, strings.NewReader(mailPayload))
				return request
			},
			expectedCode:    http.StatusAccepted,
			expectedMessage: "Message sent to simpson@gmail.com",
		},
		{
			desc: "should authenticate a user with success",
			inRequest: func() *http.Request {
				request, _ := http.NewRequest(http.MethodPost, url, strings.NewReader(authPayloadWithSuccess))
				return request
			},
			expectedCode:    http.StatusAccepted,
			expectedMessage: "Authenticated !",
		},
		{
			desc: "should return an error for bad authentication",
			inRequest: func() *http.Request {
				request, _ := http.NewRequest(http.MethodPost, url, strings.NewReader(authPayloadWithError))
				return request
			},
			expectedCode:    http.StatusUnauthorized,
			expectedMessage: "invalid credentials",
		},
	}

	for _, s := range scenarios {
		t.Run(s.desc, func(t *testing.T) {
			// Act
			client := http.Client{}
			resp, err := client.Do(s.inRequest())
			if err != nil {
				t.Fatal(err)
			}

			defer resp.Body.Close()
			message := messageFrom(resp.Body)

			// Assert
			assert.Equal(t, resp.StatusCode, s.expectedCode)
			assert.Equal(t, message, s.expectedMessage)
		})
	}
}

func prepareCompose(t *testing.T) {
	compose, err := tc.NewDockerCompose("docker-compose.yml")
	if err != nil {
		t.Fatalf("Error on docker compose %s", err)
	}

	t.Cleanup(func() {
		assert.NoError(t, compose.Down(context.Background(), tc.RemoveOrphans(true), tc.RemoveImagesLocal), "compose.Down()")
	})

	ctx, cancel := context.WithCancel(context.Background())
	t.Cleanup(cancel)

	assert.NoError(t, compose.Up(ctx, tc.Wait(true)), "compose.Up()")
}

func messageFrom(body io.ReadCloser) string {
	content, _ := io.ReadAll(body)

	var reply struct {
		Message string `json:"message"`
	}
	json.Unmarshal(content, &reply)

	return reply.Message
}
