package calljournal

type Config struct {
	Bucket string `env:"CJ_STORAGE_BUCKET, required"`
}
