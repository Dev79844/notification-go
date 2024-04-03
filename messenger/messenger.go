package messenger

type Messenger interface{
	SendMessage(message string) error
}