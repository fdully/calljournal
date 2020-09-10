package queue

import "github.com/nsqio/go-nsq"

func NewSubscriber(config *Config) Subscribe {
	nsqConfig := nsq.NewConfig()
	nsqConfig.MaxAttempts = config.NsqMaxAttempts
	nsqConfig.MsgTimeout = config.NsqMsgTimeout

	return func(topic, channel string, handler nsq.HandlerFunc) (Stop, error) {
		c, err := nsq.NewConsumer(topic, channel, nsqConfig)
		if err != nil {
			return nil, err
		}

		c.AddHandler(handler)

		return c.Stop, c.ConnectToNSQD(config.NSQDAddr)
	}
}
