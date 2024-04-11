package main

import(
	"fmt"
	"github.com/Dev79844/email-fcm-package/config"
	"github.com/Dev79844/email-fcm-package/email"
	"github.com/Dev79844/email-fcm-package/fcm"
)

func main(){
	emailConfig := config.EmailSenderConfig{
		Host:     "smtp.example.com",
		Port:     587,
		Username: "username",
		Password: "password",
		From:     "sender@example.com",
		To:       []string{"recipient@example.com"},
		Body:     "This is a test email message.",
	}

	emailSender := email.NewSMTPSender(emailConfig)

	err := emailSender.SendMessage()
	if err != nil {
		fmt.Println("Error sending email:", err)
		return
	}
	fmt.Println("Email sent successfully")

	fcmConfig := &config.FCMSenderConfig{
		ServerKey: "serviceAccountKey.json",
		Topic:     "topic",
		Message: "FCM notification",
	}

	fcmSender, err := fcm.NewFCMMessenger(*fcmConfig)
	if err!=nil{
		fmt.Println("Error initialising fcm")
		return
	}

	err = fcmSender.SendMessage()
	if err != nil {
		fmt.Println("Error sending FCM notification:", err)
		return
	}
	fmt.Println("FCM notification sent successfully")
}