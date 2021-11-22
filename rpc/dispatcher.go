package rpc

import "github.com/raf924/connector-sdk/domain"

type Dispatcher interface {
	Dispatch(message domain.ServerMessage) error
	Commands() domain.CommandList
	Done() <-chan struct{}
	Err() error
}
