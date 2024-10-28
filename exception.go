package database

import (
	"fmt"
	. "github.com/gopi-frame/contract/exception"
	"github.com/gopi-frame/exception"
)

type NotConfiguredException struct {
	Throwable
}

func NewNotConfiguredException(name string) *NotConfiguredException {
	return &NotConfiguredException{
		exception.New(fmt.Sprintf("connection [%s] not configured", name)),
	}
}

type UnregisteredDriverException struct {
	Throwable
}

func NewUnregisteredDriverException(name string) *UnregisteredDriverException {
	return &UnregisteredDriverException{
		exception.New(fmt.Sprintf("unregistered driver [%s]", name)),
	}
}

type DuplicateConnectionException struct {
	Throwable
}

func NewDuplicateDriverException(name string) *DuplicateConnectionException {
	return &DuplicateConnectionException{
		exception.New(fmt.Sprintf("duplicate driver [%s]", name)),
	}
}
