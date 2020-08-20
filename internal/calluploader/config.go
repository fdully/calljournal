package calluploader

import "time"

type Config struct {
	BaseCallDir    string        `env:"CJ_BASECALL_DIR, required"`
	ReadDirPeriod  time.Duration `env:"CJ_READ_DIR_PERIOD, default=20s"`
	GrpcServerAddr string        `env:"CJ_GRPC_SERVER_ADDR, default=localhost:9111"`
	StorageAddr    string        `env:"CJ_STORAGE_ADDR, required"`
}
