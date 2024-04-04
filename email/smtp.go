package email

import (
	"net/smtp"
	"strconv"

	"github.com/Dev79844/email-fcm-package/config"
)

type smtpSender struct {
	config 		config.EmailSenderConfig
	to 	 		[]string
	body 		string
}

func NewSMTPSender(config config.EmailSenderConfig) *smtpSender {
	return &smtpSender{config: config}
}
 
func (s *smtpSender) SendMessage() error {

	auth := smtp.PlainAuth("", s.config.Username, s.config.Password, s.config.Password)

	return smtp.SendMail(s.config.Host+":"+strconv.Itoa(s.config.Port), auth, s.config.From, s.to, []byte(s.body))
}
