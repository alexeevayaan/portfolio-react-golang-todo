package redis

import (
	"github.com/redis/go-redis/v9"
	"github.com/rs/zerolog/log"
)



type Config struct{
	Addr string `envconfig:"REDIS_ADDR" required:"true"`
	Password string `envconfig:"REDIS_PASSWORD"`
	DB int `envconfig:"REDIS_DB" default:"0"`
}

type CLient struct{
	*redis.Client
}

func New(c Config) (*CLient, error){
	client:=redis.NewClient(&redis.Options{
		Addr: c.Addr,
		Password: c.Password,
		DB: c.DB,
	})

	return &CLient{Client: client},nil
}


func (c *CLient) Close(){

	err:=c.Client.Close()
	if err != nil{
		log.Error().Err(err).Msg("redis: close")
	}

	log.Info().Msg("redis: closed")
}