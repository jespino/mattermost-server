package systembus

import (
	"errors"
	"time"

	"github.com/nats-io/nats-server/v2/server"
	"github.com/nats-io/nats.go"
)

const (
	PostCreatedSubject = "create_post"
	PostUpdatedSubject = "create_updated"
)

type Event struct {
	Subject string
	Data    []byte
}

type SystemBus struct {
	natsServer *server.Server
	natsConn   *nats.Conn
}

func New() (*SystemBus, error) {
	opts := &server.Options{}
	ns, err := server.NewServer(opts)
	if err != nil {
		return nil, err
	}
	go ns.Start()
	return &SystemBus{natsServer: ns}, nil
}

func (sb *SystemBus) Start() error {
	if !sb.natsServer.ReadyForConnections(4 * time.Second) {
		return errors.New("not ready for connection")
	}
	nc, err := nats.Connect(sb.natsServer.ClientURL())
	if err != nil {
		return err
	}
	sb.natsConn = nc

	return nil
}

func (sb *SystemBus) Stop() {
	sb.natsServer.Shutdown()
}

func (sb *SystemBus) Subscribe(subject string, handler func(data []byte)) {
	sb.natsConn.Subscribe(subject, func(msg *nats.Msg) {
		handler(msg.Data)
	})
}

func (sb *SystemBus) Publish(subject string, data []byte) {
	sb.natsConn.Publish(subject, data)
}
