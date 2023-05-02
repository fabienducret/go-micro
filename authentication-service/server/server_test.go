package server_test

import (
	"authentication/adapters"
	"authentication/server"
	"testing"
)

func TestAuthenticate(t *testing.T) {
	s := server.NewServer(
		adapters.NewUserTestRepository(),
		adapters.NewTestLogger(),
	)

	t.Run("valid_credentials", func(t *testing.T) {
		payload := server.Payload{
			Email:    "test@gmail.com",
			Password: "password",
		}

		var reply server.Identity
		err := s.Authenticate(payload, &reply)
		if err != nil {
			t.Errorf("Test failed with error %s", err)
		}

		if reply.Email != "test@gmail.com" {
			t.Errorf("Test failed, bad email received %s", reply.Email)
		}
	})

	t.Run("invalid_credentials", func(t *testing.T) {
		payload := server.Payload{
			Email:    "test@gmail.com",
			Password: "toto",
		}

		var reply server.Identity
		err := s.Authenticate(payload, &reply)

		if err == nil {
			t.Errorf("Error must be defined")
		}
	})
}
