package rpc

import (
	"context"
	"github.com/raf924/connector-sdk/domain"
)

var connectorRelayBuilders = map[string]ConnectorRelayBuilder{}

type ConnectorRelayBuilder func(config interface{}) ConnectorRelay

func RegisterConnectorRelay(name string, relayBuilder ConnectorRelayBuilder) {
	connectorRelayBuilders[name] = relayBuilder
}

var _ = RegisterConnectorRelay

func GetConnectorRelay(relayKey string) ConnectorRelayBuilder {
	if builder, ok := connectorRelayBuilders[relayKey]; ok {
		return builder
	}
	return nil
}

var _ = GetConnectorRelay

type ConnectorRelay interface {
	Start(ctx context.Context, botUser *domain.User, onlineUsers domain.UserList, trigger string) error
	Accept() (Dispatcher, error)
	Recv() (*domain.ClientMessage, error)
	Done() <-chan struct{}
	Err() error
}
