package redis

import (
	"time"

	"github.com/alexeevayaan/portfolio-react-golang-todo/api/pkg/redis"
)


const (
	idempotencyPrefix = "api-idempotency:"
	ttl = time.Minute *10
)

type Redis struct{
	redis *redis.CLient
}
func New(client *redis.CLient) *Redis{
	return &Redis{
		redis: client,
	}
}