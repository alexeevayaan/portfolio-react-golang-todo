package main

import (
	"context"

	"github.com/alexeevayaan/portfolio-react-golang-todo/api/config"
	"github.com/alexeevayaan/portfolio-react-golang-todo/api/internal/app"
	"github.com/rs/zerolog/log"
)


func main(){

	c,err:= config.New()
	if err != nil{
		log.Fatal().Err(err).Msg("config.New")
	}

	ctx := context.Background()

	err = app.Run(ctx,c)
	if err !=nil{
		log.Error().Err(err).Msg("app.Run")
	}
}