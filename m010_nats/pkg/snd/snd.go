package snd

import (
	"github.com/nats-io/nats-server/v2/server"
)

func Server() (*server.Server, error) {
	opts := &server.Options{Port: 12222, JetStream: true}
	ns, err := server.NewServer(opts)

	if err != nil {
		return nil, err
	}
	return ns, nil
}
