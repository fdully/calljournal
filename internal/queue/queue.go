package queue

import (
	"io"
)

type Publisher interface {
	Publish(topic string, msg []byte) error
	io.Closer
}
