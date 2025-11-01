package app

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/alexeevayaan/portfolio-react-golang-todo/api/config"
	"github.com/alexeevayaan/portfolio-react-golang-todo/api/internal/adapter/postgres"
	"github.com/alexeevayaan/portfolio-react-golang-todo/api/internal/adapter/redis"
	"github.com/alexeevayaan/portfolio-react-golang-todo/api/internal/controller/http"
	"github.com/alexeevayaan/portfolio-react-golang-todo/api/internal/usecase"
	"github.com/alexeevayaan/portfolio-react-golang-todo/api/pkg/httpserver"
	db "github.com/alexeevayaan/portfolio-react-golang-todo/api/pkg/postgres"
	redisLib "github.com/alexeevayaan/portfolio-react-golang-todo/api/pkg/redis"
	"github.com/alexeevayaan/portfolio-react-golang-todo/api/pkg/router"
	"github.com/alexeevayaan/portfolio-react-golang-todo/api/pkg/transaction"
	"github.com/rs/zerolog/log"
)


func Run(ctx context.Context, c config.Config) error{

	pgPool, err:= db.New(ctx, c.Postgres)
	if err !=nil{
		return fmt.Errorf("postgres.New: %w",err)
	}

	transaction.Init(pgPool)

	redisClient, err:= redisLib.New(c.Redis)
		if err !=nil{
		return fmt.Errorf("redis.New: %w",err)
	}

	uc :=usecase.New(
		postgres.New(), 
		redis.New(redisClient),
	)
	
	// HTTP
	r:=router.New()
	
	http.Router(r, uc)

	httpServer:= httpserver.New(r, c.HTTP)

	log.Info().Msg("App started")

	sig:= make(chan os.Signal, 1)

	signal.Notify(sig, syscall.SIGINT, syscall.SIGTERM)

	<-sig

	log.Info().Msg("App got signal to stop")

	httpServer.Close()

	redisClient.Close()
	pgPool.Close()

	log.Info().Msg("App stopped!")

	return nil
}