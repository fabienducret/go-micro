package server_test

import (
	"authentication/adapters/tests"
	"authentication/server"
	"testing"
)

func TestAuthenticate(t *testing.T) {
	payloadWithValidPassword := server.Payload{
		Email:    "test@gmail.com",
		Password: "password",
	}

	payloadWithInvalidPassword := server.Payload{
		Email:    "test@gmail.com",
		Password: "toto",
	}

	t.Run("with valid credentials", func(t *testing.T) {
		// Given
		s := server.New(tests.UserRepositoryStub{}, tests.LoggerStub{})

		// When
		var reply server.Identity
		err := s.Authenticate(payloadWithValidPassword, &reply)

		// Then
		assertErrorIsNil(t, err)
		assertEqual(t, reply.Email, "test@gmail.com")
	})

	t.Run("with invalid credentials", func(t *testing.T) {
		// Given
		s := server.New(tests.UserRepositoryStub{}, tests.LoggerStub{})

		// When
		err := s.Authenticate(payloadWithInvalidPassword, nil)

		// Then
		assertErrorIsDefined(t, err)
	})

	t.Run("error in logger call", func(t *testing.T) {
		// Given
		s := server.New(
			tests.UserRepositoryStub{},
			tests.LoggerStub{WithError: true},
		)

		// When
		err := s.Authenticate(payloadWithValidPassword, nil)

		// Then
		assertErrorIsDefined(t, err)
	})
}

func assertErrorIsNil(t *testing.T, err error) {
	if err != nil {
		t.Errorf("Test failed with error %s", err)
	}
}

func assertEqual(t *testing.T, got, want string) {
	if got != want {
		t.Errorf("Test failed got %s, want %s", got, want)
	}
}

func assertErrorIsDefined(t *testing.T, err error) {
	if err == nil {
		t.Errorf("Error must be defined")
	}
}
