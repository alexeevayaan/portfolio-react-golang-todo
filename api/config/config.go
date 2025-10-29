package config

import (
	"fmt"

	"github.com/alexeevayaan/portfolio-react-golang-todo/api/pkg/httpserver"
	"github.com/alexeevayaan/portfolio-react-golang-todo/api/pkg/postgres"
	"github.com/alexeevayaan/portfolio-react-golang-todo/api/pkg/redis"
	"github.com/joho/godotenv"
	"github.com/kelseyhightower/envconfig"
	"github.com/rs/zerolog/log"
)

type App struct{
	Name string `envconfig:"APP_NAME" required:"true"`
	Version string `envconfig:"APP_VERSION" required:"true"`
}

type Config struct{
	App App
	HTTP httpserver.Config
	Postgres postgres.Config
	Redis redis.Config
}

func New() (Config, error){
	var config Config

	err:= godotenv.Load("../.env")
	if err != nil{
		return config, fmt.Errorf("godotenv.Load: %w", err)
	}

	err = envconfig.Process("", &config)
	if err !=nil{
		return config, fmt.Errorf("envconfig.Process: %w", err)
	}


	path := fmt.Sprintf("config loaded: %s", config)

	log.Info().Msg(path)

	return config, nil
}