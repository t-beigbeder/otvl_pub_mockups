package snd

import (
	"context"
	"fmt"
	"github.com/nats-io/nats.go"
	"github.com/nats-io/nats.go/jetstream"
	"testing"
	"time"
)

func TestServerBasic(t *testing.T) {
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
	msgs, err := cons.Fetch(1, jetstream.FetchMaxWait(1*time.Second))
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

func TestServerTwoPub(t *testing.T) {
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
	_, err = jss.Publish("example-subject", []byte("Hello jss one!"))
	if err != nil {
		t.Fatal(err)
	}
	_, err = jss.Publish("example-subject", []byte("Hello jss two!"))
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
	msgs, err := cons.Fetch(2, jetstream.FetchMaxWait(1*time.Second))
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

	ncr2, err := nats.Connect(ns.ClientURL())
	if err != nil {
		t.Fatal(err)
	}
	jsr2, err := jetstream.New(ncr2)
	if err != nil {
		t.Fatal(err)
	}
	cons2, err := jsr2.OrderedConsumer(context.Background(), "example-stream", jetstream.OrderedConsumerConfig{
		FilterSubjects: []string{"example-subject"},
	})
	msgs2, err := cons2.Fetch(2, jetstream.FetchMaxWait(1*time.Second))
	if err != nil {
		t.Fatal(err)
	}
	for msg := range msgs2.Messages() {
		fmt.Printf("Received a JetStream message: %s\n", string(msg.Data()))
		err = msg.Ack()
		if err != nil {
			t.Fatal(err)
		}
	}

	ns.Shutdown()
}

func TestServerTwoOne(t *testing.T) {
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
	_, err = jss.Publish("example-subject", []byte("Hello jss one!"))
	if err != nil {
		t.Fatal(err)
	}
	_, err = jss.Publish("example-subject", []byte("Hello jss two!"))
	if err != nil {
		t.Fatal(err)
	}
	jssc, err := jetstream.New(ncs)
	if err != nil {
		t.Fatal(err)
	}
	_, err = jssc.OrderedConsumer(context.Background(), "example-stream", jetstream.OrderedConsumerConfig{
		FilterSubjects: []string{"example-subject"},
	})

	ncr, err := nats.Connect(ns.ClientURL())
	if err != nil {
		t.Fatal(err)
	}
	jsr, err := jetstream.New(ncr)
	if err != nil {
		t.Fatal(err)
	}
	sr, err := jsr.Stream(context.Background(), "example-stream")
	if err != nil {
		t.Fatal(err)
	}

	var cons jetstream.Consumer
	cis := sr.ListConsumers(context.Background())
	for ci := range cis.Info() {
		cons, err = sr.Consumer(context.Background(), ci.Name)
		if err != nil {
			t.Fatal(err)
		}
	}
	if cons == nil {
		t.Fatal("cons is nil")
	}
	msgs, err := cons.Fetch(2, jetstream.FetchMaxWait(1*time.Second))
	if err != nil {
		t.Fatal(err)
	}
	for msg := range msgs.Messages() {
		fmt.Printf("Received a JetStream message: %s\n", string(msg.Data()))
		err = msg.Ack()
		if err != nil {
			t.Fatal(err)
		}
	}

	var cons2 jetstream.Consumer
	cis = sr.ListConsumers(context.Background())
	for ci := range cis.Info() {
		cons2, err = sr.Consumer(context.Background(), ci.Name)
		if err != nil {
			t.Fatal(err)
		}
	}
	if cons2 == nil {
		t.Fatal("cons2 is nil")
	}
	msgs, err = cons2.Fetch(2, jetstream.FetchMaxWait(100*time.Millisecond))
	if err != nil {
		t.Fatal(err)
	}
	for msg := range msgs.Messages() {
		fmt.Printf("Received a JetStream message 2: %s\n", string(msg.Data()))
		err = msg.Ack()
		if err != nil {
			t.Fatal(err)
		}
	}

	ns.Shutdown()
}
