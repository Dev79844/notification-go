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
	auth		smtp.Auth
}

func NewSMTPSender(config config.EmailSenderConfig) *smtpSender {
	auth := smtp.PlainAuth("", config.Username, config.Password, config.Password)
	return &smtpSender{
		config: config,
		auth: auth,
	}
}
 
func (s *smtpSender) SendMessage() error {
	return smtp.SendMail(s.config.Host+":"+strconv.Itoa(s.config.Port), s.auth, s.config.From, s.to, []byte(s.body))
}
