package model

import (
	"errors"
)

var (
	connectSMTPError = errors.New("cannot connect to smtp server")
)

type Sender interface {
	Send(data interface{}) bool
}
