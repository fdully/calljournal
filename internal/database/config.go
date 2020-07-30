package database

import (
	"time"
)

type Config struct {
	Name               string        `env:"CJ_DB_NAME"`
	User               string        `env:"CJ_DB_USER"`
	Host               string        `env:"CJ_DB_HOST, default=localhost"`
	Port               string        `env:"CJ_DB_PORT, default=5432"`
	SSLMode            string        `env:"CJ_DB_SSLMODE, default=disable"`
	ConnectionTimeout  int           `env:"CJ_DB_CONNECT_TIMEOUT"`
	Password           string        `env:"CJ_DB_PASSWORD"`
	SSLCertPath        string        `env:"CJ_DB_SSLCERT"`
	SSLKeyPath         string        `env:"CJ_DB_SSLKEY"`
	SSLRootCertPath    string        `env:"CJ_DB_SSLROOTCERT"`
	PoolMinConnections string        `env:"CJ_DB_POOL_MIN_CONNS"`
	PoolMaxConnections string        `env:"CJ_DB_POOL_MAX_CONNS"`
	PoolMaxConnLife    time.Duration `env:"CJ_DB_POOL_MAX_CONN_LIFETIME"`
	PoolMaxConnIdle    time.Duration `env:"CJ_DB_POOL_MAX_CONN_IDLE_TIME"`
	PoolHealthCheck    time.Duration `env:"CJ_DB_POOL_HEALTH_CHECK_PERIOD"`
}
