package server_test

import (
	"mailer-service/repositories"
	"mailer-service/server"
	"testing"
)

func TestSendMail(t *testing.T) {
	s := server.NewServer(repositories.NewMailhogTestRepository())

	payload := server.Payload{
		From:    "from@example.com",
		To:      "homer@example.com",
		Subject: "Subject",
		Message: "Hello Homer !",
	}

	var reply string
	err := s.SendMail(payload, &reply)
	if err != nil {
		t.Errorf("Test failed with error %s", err)
	}

	if reply != "Message sent to homer@example.com" {
		t.Errorf("Test failed with reply %s", reply)
	}
}
