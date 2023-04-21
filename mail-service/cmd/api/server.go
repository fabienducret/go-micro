package main

type Server struct {
	Mailer Mail
}

type Payload struct {
	From    string
	To      string
	Subject string
	Message string
}

func (r *Server) SendMail(payload Payload, resp *string) error {
	msg := Message{
		From:    payload.From,
		To:      payload.To,
		Subject: payload.Subject,
		Data:    payload.Message,
	}

	err := r.Mailer.SendSMTPMessage(msg)
	if err != nil {
		return err
	}

	*resp = "Message sent to " + payload.To

	return nil
}
