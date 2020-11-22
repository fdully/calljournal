package cdrclient

import (
	"time"

	"google.golang.org/grpc"
)

const (
	GRPCMaxMsgSize   = 5 * 1024 * 1024 * 1024 * 1024
	GRPCMsgByteChunk = 64 * 1024
)

type Config struct {
	BaseCallDir    string        `env:"CJ_BASECALL_DIR, required"`
	ReadDirPeriod  time.Duration `env:"CJ_READ_DIR_PERIOD, default=60s"`
	GrpcServerAddr string        `env:"CJ_GRPC_SERVER_ADDR, default=localhost:9111"`
	AllCDR         bool          `env:"CJ_ALL_CDR, default=true"`
	NumWorkers     int           `env:"CJ_NUM_WORKERS, default=5"`
}

func (c *Config) GRPCPoolConfig() (string, []grpc.DialOption, int) {
	var dialOpt []grpc.DialOption
	dialOpt = append(dialOpt, grpc.WithDefaultCallOptions(grpc.MaxCallRecvMsgSize(GRPCMaxMsgSize),
		grpc.MaxCallSendMsgSize(GRPCMaxMsgSize)), grpc.WithInsecure())

	return c.GrpcServerAddr, dialOpt, c.NumWorkers
}
