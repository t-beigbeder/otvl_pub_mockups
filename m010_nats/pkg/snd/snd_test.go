package snd

import (
	"context"
	"fmt"
	"github.com/nats-io/nats.go"
	"github.com/nats-io/nats.go/jetstream"
	"testing"
	"time"
)

func TestServer(t *testing.T) {
	ns, err := Server()
	if err != nil {
		t.Fatal(err)
	}
	go ns.Start()
	ns.ReadyForConnections(5 * time.Second)
	ncs, err := nats.Connect(ns.ClientURL())
	if err != nil {
		t.Fatal(err)
	}
	_ = ncs
	jss, err := ncs.JetStream()
	if err != nil {
		t.Fatal(err)
	}

	// Create a stream
	_, err = jss.AddStream(&nats.StreamConfig{
		Name:     "example-stream",
		Subjects: []string{"example-subject"},
		MaxBytes: 1024,
		Storage:  nats.MemoryStorage,
	})
	if err != nil {
		t.Fatal(err)
	}
	_, err = jss.Publish("example-subject", []byte("Hello jss!"))
	if err != nil {
		t.Fatal(err)
	}

	ncr, err := nats.Connect(ns.ClientURL())
	if err != nil {
		t.Fatal(err)
	}
	jsr, err := jetstream.New(ncr)
	if err != nil {
		t.Fatal(err)
	}
	cons, err := jsr.OrderedConsumer(context.Background(), "example-stream", jetstream.OrderedConsumerConfig{
		FilterSubjects: []string{"example-subject"},
	})
	msgs, err := cons.Fetch(10, jetstream.FetchMaxWait(1*time.Second))
	if err != nil {
		t.Fatal(err)
	}
	_ = msgs
	for msg := range msgs.Messages() {
		fmt.Printf("Received a JetStream message: %s\n", string(msg.Data()))
		err = msg.Ack()
		if err != nil {
			t.Fatal(err)
		}
	}
	ns.Shutdown()
}
