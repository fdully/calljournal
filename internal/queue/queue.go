package queue

import (
	"io"

	"github.com/nsqio/go-nsq"
)

type Publisher interface {
	Publish(topic string, msg []byte) error
	io.Closer
}

type Stop func()

type Subscribe func(topic, channel string, handler nsq.HandlerFunc) (Stop, error)
