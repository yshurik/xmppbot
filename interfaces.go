package xmppbot

type Bot interface {
	Name() string
	FullName() string
	Send(msg string)
	Connect() error
	Listen() chan Message
}

//*************************************************
type Message interface {
	Body() string
	From() string
}
