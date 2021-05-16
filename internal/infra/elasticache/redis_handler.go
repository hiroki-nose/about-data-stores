package elasticache

import (
	"context"
	"github.com/go-redis/redis/v8"
)

var ctx = context.Background()

type ElastiCacheHandler struct {
	DbService *redis.Client
}

func NewElastiCacheHandler() (*ElastiCacheHandler, error) {
	rdb := redis.NewClient(&redis.Options{
		Addr:     "{hostname}:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	handler := ElastiCacheHandler{
		DbService: rdb,
	}
	return &handler, nil
}
func (e ElastiCacheHandler) Get(key string) (string, error) {
	res, err := e.DbService.Get(ctx, key).Result()
	return res, err
}
