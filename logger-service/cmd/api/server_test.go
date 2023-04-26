package main

import (
	"log-service/repositories"
	"testing"
)

func TestLogInfo(t *testing.T) {
	s := Server{
		LogRepository: repositories.NewLogTestRepository(),
	}

	payload := Payload{
		Name: "testevent",
		Data: "data to log",
	}

	var reply string
	err := s.LogInfo(payload, &reply)
	if err != nil {
		t.Errorf("Test failed with error %s", err)
	}

	if reply != "Log handled for:testevent" {
		t.Errorf("Test failed with reply %s", reply)
	}
}
