package queue

import (
	"context"
	"fmt"

	"github.com/nsqio/go-nsq"
)

type pub struct {
	p *nsq.Producer
}

func NewPub(ctx context.Context, c *Config) (Publisher, error) {
	config := nsq.NewConfig()

	prod, err := nsq.NewProducer(c.NSQDAddr, config)
	if err != nil {
		return nil, fmt.Errorf("failed to create new nsq producer: %w", err)
	}

	return &pub{p: prod}, nil
}

func (p *pub) Publish(topic string, msg []byte) error {
	return p.p.Publish(topic, msg)
}

func (p *pub) Close() error {
	if p.p != nil {
		p.p.Stop()
	}

	return nil
}
