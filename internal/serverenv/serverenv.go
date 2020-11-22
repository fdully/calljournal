package serverenv

import (
	"context"

	"github.com/fdully/calljournal/internal/database"
	"github.com/fdully/calljournal/internal/queue"
	"github.com/fdully/calljournal/internal/storage"
	gpool "github.com/processout/grpc-go-pool"
)

type ServerEnv struct {
	blobstore  storage.Blobstore
	database   *database.DB
	publisher  queue.Publisher
	subscriber queue.Subscribe
	grpcpool   *gpool.Pool
}

type Option func(*ServerEnv) *ServerEnv

func NewServerEnv(ctx context.Context, opts ...Option) *ServerEnv {
	env := &ServerEnv{}

	for _, f := range opts {
		env = f(env)
	}

	return env
}

func WithDatabase(db *database.DB) Option {
	return func(env *ServerEnv) *ServerEnv {
		env.database = db

		return env
	}
}

func WithBlobstore(s storage.Blobstore) Option {
	return func(env *ServerEnv) *ServerEnv {
		env.blobstore = s

		return env
	}
}

func WithPublisher(p queue.Publisher) Option {
	return func(env *ServerEnv) *ServerEnv {
		env.publisher = p

		return env
	}
}

func WithSubscriber(s queue.Subscribe) Option {
	return func(env *ServerEnv) *ServerEnv {
		env.subscriber = s

		return env
	}
}

func WithGRPCPool(p *gpool.Pool) Option {
	return func(env *ServerEnv) *ServerEnv {
		env.grpcpool = p

		return env
	}
}

func (s *ServerEnv) GRPCPool() *gpool.Pool {
	return s.grpcpool
}

func (s *ServerEnv) Blobstore() storage.Blobstore {
	return s.blobstore
}

func (s *ServerEnv) Database() *database.DB {
	return s.database
}

func (s *ServerEnv) Publisher() queue.Publisher {
	return s.publisher
}

func (s *ServerEnv) Subscriber() queue.Subscribe {
	return s.subscriber
}

func (s *ServerEnv) Close(ctx context.Context) error {
	if s == nil {
		return nil
	}

	var err error

	if s.database != nil {
		s.database.Close(ctx)
	}

	if s.publisher != nil {
		_ = s.publisher.Close()
	}

	if s.grpcpool != nil {
		s.grpcpool.Close()
	}

	return err
}
