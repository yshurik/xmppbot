package xmppbot

import (
	"log"
)

type Bot interface {
	Name() string
	FullName() string
	Send(msg string)
	Connect() error
	Listen() chan Message
	SetLogger(*log.Logger)
	Log(msg string)
}

//*************************************************
type Message interface {
	Body() string
	From() string
}
