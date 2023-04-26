package main

import (
	"authentication/repositories"
	"testing"
)

func TestAuthenticate(t *testing.T) {
	s := Server{
		UserRepository:   repositories.NewPostgresTestRepository(),
		LoggerRepository: repositories.NewLoggerTestRepository(),
	}

	t.Run("valid_credentials", func(t *testing.T) {
		payload := Payload{
			Email:    "test@gmail.com",
			Password: "password",
		}

		var reply Identity
		err := s.Authenticate(payload, &reply)
		if err != nil {
			t.Errorf("Test failed with error %s", err)
		}

		if reply.Email != "test@gmail.com" {
			t.Errorf("Test failed with error %s", err)
		}
	})
	t.Run("invalid_credentials", func(t *testing.T) {
		payload := Payload{
			Email:    "test@gmail.com",
			Password: "toto",
		}

		var reply Identity
		err := s.Authenticate(payload, &reply)
		if err == nil {
			t.Errorf("Error must be defined")
		}
	})
}
