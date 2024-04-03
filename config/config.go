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
}

type FCMSenderConfig struct {
    ServerKey string
	Topic     string
}
