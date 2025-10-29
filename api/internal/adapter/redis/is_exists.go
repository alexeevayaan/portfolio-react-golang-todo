package redis

import (
	"context"
	"errors"

	"github.com/redis/go-redis/v9"
	"github.com/rs/zerolog/log"
)


func (r *Redis) IsExists(ctx context.Context, idempotencyKey string) bool{

	key:= idempotencyPrefix + idempotencyKey

	err:= r.redis.Get(ctx,key).Err()
	if err != nil{
		return true
	}

	if !errors.Is(err, redis.Nil){
		log.Error().Err(err).Msg("redis: IsExists: r.redis.Get")
	}

	err = r.redis.Set(ctx, key, []byte{}, ttl).Err()
	if err !=nil{
		log.Error().Err(err).Msg("redis: IsExists: r.redis.Set")
	}

	return false
}