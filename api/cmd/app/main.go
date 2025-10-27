package main

import (
	"context"

	"github.com/alexeevayaan/portfolio-react-golang-todo/api/config"
	"github.com/rs/zerolog/log"
)


func main(){

	c,err:= config.New()

	if err != nil{
		log.Fatal().Err(err).Msg("config.New")
	}

	ctx := context.Background()


		

}