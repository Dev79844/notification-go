package email

import (
	"net/smtp"
	"strconv"

	"github.com/Dev79844/email-fcm-package/config"
)

type smtpSender struct {
	config config.EmailSenderConfig
}

func NewSMTPSender(config config.EmailSenderConfig) Sender {
	return &smtpSender{config: config}
}
 
func (s *smtpSender) SendEmail(to []string, subject, body string) error {

	auth := smtp.PlainAuth("", s.config.Username, s.config.Password, s.config.Password)
	msg := []byte("Hello there!!!")

	return smtp.SendMail(s.config.Host+":"+strconv.Itoa(s.config.Port), auth, s.config.From, to, msg)
}
