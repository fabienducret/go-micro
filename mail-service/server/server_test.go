package server_test

import (
	"mailer-service/adapters/tests"
	"mailer-service/server"
	"testing"
)

func TestSendMail(t *testing.T) {
	// Given
	s := server.NewServer(tests.NewMailRepositoryStub())
	payload := server.Payload{
		From:    "from@example.com",
		To:      "homer@example.com",
		Subject: "Subject",
		Message: "Hello Homer !",
	}

	// When
	var reply string
	err := s.SendMail(payload, &reply)

	// Then
	if err != nil {
		t.Errorf("Test failed with error %s", err)
	}

	if reply != "Message sent to homer@example.com" {
		t.Errorf("Test failed with reply %s", reply)
	}
}
