package server_test

import (
	"log-service/repositories"
	"log-service/server"
	"testing"
)

func TestLogInfo(t *testing.T) {
	s := server.NewServer(repositories.NewLogTestRepository())

	payload := server.Payload{
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
