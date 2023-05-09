package server_test

import (
	"mailer-service/adapters/tests"
	"mailer-service/server"
	"testing"
)

func TestSendMail(t *testing.T) {
	payload := server.Payload{
		From:    "from@example.com",
		To:      "homer@example.com",
		Subject: "Subject",
		Message: "Hello Homer !",
	}

	t.Run("send mail with success", func(t *testing.T) {
		// Given
		s := server.NewServer(tests.MailRepositoryStub{})

		// When
		var reply string
		err := s.SendMail(payload, &reply)

		// Then
		assertErrorIsNil(t, err)
		assertEqual(t, reply, "Message sent to homer@example.com")
	})

	t.Run("send mail with error", func(t *testing.T) {
		// Given
		s := server.NewServer(tests.MailRepositoryStub{WithError: true})

		// When
		err := s.SendMail(payload, nil)

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
