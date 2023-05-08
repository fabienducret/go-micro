package server_test

import (
	"authentication/adapters/tests"
	"authentication/server"
	"testing"
)

func TestAuthenticate(t *testing.T) {
	s := server.NewServer(
		tests.NewUserRepositoryStub(),
		tests.NewLoggerStub(),
	)

	t.Run("valid_credentials", func(t *testing.T) {
		// Given
		payload := server.Payload{
			Email:    "test@gmail.com",
			Password: "password",
		}

		// When
		var reply server.Identity
		err := s.Authenticate(payload, &reply)
		if err != nil {
			t.Errorf("Test failed with error %s", err)
		}

		// Then
		if reply.Email != "test@gmail.com" {
			t.Errorf("Test failed, bad email received %s", reply.Email)
		}
	})

	t.Run("invalid_credentials", func(t *testing.T) {
		// Given
		payload := server.Payload{
			Email:    "test@gmail.com",
			Password: "toto",
		}

		// When
		var reply server.Identity
		err := s.Authenticate(payload, &reply)

		// Then
		if err == nil {
			t.Errorf("Error must be defined")
		}
	})
}
