package server_test

import (
	"log-service/adapters/tests"
	"log-service/server"
	"testing"
)

func TestLogInfo(t *testing.T) {
	payload := server.Payload{
		Name: "testevent",
		Data: "data to log",
	}

	t.Run("log with success", func(t *testing.T) {
		// Given
		s := server.New(tests.LogRepositoryStub{})

		// When
		var reply string
		err := s.LogInfo(payload, &reply)

		// Then
		assertErrorIsNil(t, err)
		assertEqual(t, reply, "Log handled for:testevent")
	})

	t.Run("log with error", func(t *testing.T) {
		// Given
		s := server.New(tests.LogRepositoryStub{WithError: true})

		// When
		err := s.LogInfo(payload, nil)

		// Then
		assertErrorIsDefined(t, err)
	})
}

func assertErrorIsNil(t *testing.T, err error) {
	if err != nil {
		t.Errorf("Test failed with error %s", err)
	}
}

func assertErrorIsDefined(t *testing.T, err error) {
	if err == nil {
		t.Error("Error must be defined")
	}
}

func assertEqual(t *testing.T, got, want string) {
	if got != want {
		t.Errorf("Test failed got %s, want %s", got, want)
	}
}
