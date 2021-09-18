package mymock

type API interface {
	SendMessage(msg string) error
	ConsumeMessage() (string, error)
}

type Messenger struct {
	api API
}
