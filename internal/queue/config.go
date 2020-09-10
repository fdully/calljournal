package queue

import "time"

type Config struct {
	NSQDAddr       string        `env:"CJ_NSQD_ADDR, default=localhost:4150"`
	NsqMaxAttempts uint16        `env:"CJ_NSQ_MAX_ATTEMPTS, default=1000"`
	NsqMsgTimeout  time.Duration `env:"CJ_NSQ_MSG_TIMEOUT, default=30m"`
}
