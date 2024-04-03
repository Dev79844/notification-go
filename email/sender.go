package email

type Sender interface {
	SendEmail(to []string, subject, body string) error
}
