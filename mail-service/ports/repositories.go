package ports

type Message struct {
	From        string
	FromName    string
	To          string
	Subject     string
	Attachments []string
	Data        any
	DataMap     map[string]any
}

type MailRepository interface {
	SendSMTPMessage(Message) error
}
