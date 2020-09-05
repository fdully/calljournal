package serverenv

import (
	"context"

	"github.com/fdully/calljournal/internal/database"
	"github.com/fdully/calljournal/internal/queue"
	"github.com/fdully/calljournal/internal/storage"
)

type ServerEnv struct {
	blobstore storage.Blobstore
	database  *database.DB
	publisher queue.Publisher
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

func (s *ServerEnv) Blobstore() storage.Blobstore {
	return s.blobstore
}

func (s *ServerEnv) Database() *database.DB {
	return s.database
}

func (s *ServerEnv) Publisher() queue.Publisher {
	return s.publisher
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

	return err
}
