package queue

type Config struct {
	Addr string `env:"CJ_NSQD_ADDR, default=localhost:4150"`
}
