package redis

import rdb "github.com/redis/go-redis/v9"

func (cfg *Config) Connect() *rdb.Client {
	return rdb.NewClient(&rdb.Options{
		Addr:     cfg.Addr,
		Password: cfg.Password,
	})
}
