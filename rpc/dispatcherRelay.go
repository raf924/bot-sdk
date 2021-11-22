package rpc

import (
	"github.com/raf924/connector-sdk/domain"
)

var dispatcherRelayBuilders = map[string]DispatcherRelayBuilder{}

type DispatcherRelayBuilder func(config interface{}) DispatcherRelay

func RegisterDispatcherRelay(key string, relayBuilder DispatcherRelayBuilder) {
	dispatcherRelayBuilders[key] = relayBuilder
}

var _ = RegisterDispatcherRelay

func GetDispatcherRelay(relayKey string) DispatcherRelayBuilder {
	if builder, ok := dispatcherRelayBuilders[relayKey]; ok {
		return builder
	}
	return nil
}

var _ = GetDispatcherRelay

type DispatcherRelay interface {
	Connect(registration *domain.RegistrationMessage) (*domain.ConfirmationMessage, error)
	Send(packet *domain.ClientMessage) error
	Recv() (domain.ServerMessage, error)
	Done() <-chan struct{}
	Err() error
}
