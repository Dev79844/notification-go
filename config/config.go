package config

type Config struct {
	EmailSender EmailSenderConfig
	FCMSender	FCMSenderConfig
}

type EmailSenderConfig struct {
	Host     string
	Port     int
	Username string
	Password string
	From     string
	To		 []string
	Body 	 string
}

type FCMSenderConfig struct {
    ServerKey string
	Topic     string
	Message	  string
}
