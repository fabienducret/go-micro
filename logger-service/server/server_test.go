package server_test

import (
	"log-service/adapters/tests"
	"log-service/server"
	"testing"
)

func TestLogInfo(t *testing.T) {
	// Given
	s := server.NewServer(tests.NewLogRepositoryStub())
	payload := server.Payload{
		Name: "testevent",
		Data: "data to log",
	}

	// When
	var reply string
	err := s.LogInfo(payload, &reply)

	// Then
	if err != nil {
		t.Errorf("Test failed with error %s", err)
	}

	if reply != "Log handled for:testevent" {
		t.Errorf("Test failed with reply %s", reply)
	}
}
