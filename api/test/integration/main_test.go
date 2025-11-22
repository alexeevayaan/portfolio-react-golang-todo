//go:build integration

package test

import (
	"context"
	"testing"
	"time"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"

	"github.com/alexeevayaan/portfolio-react-golang-todo/api/config"
	"github.com/alexeevayaan/portfolio-react-golang-todo/api/internal/app"
	"github.com/alexeevayaan/portfolio-react-golang-todo/api/pkg/httpclient"
	"github.com/alexeevayaan/portfolio-react-golang-todo/api/pkg/httpserver"
	"github.com/alexeevayaan/portfolio-react-golang-todo/api/pkg/postgres"
	"github.com/alexeevayaan/portfolio-react-golang-todo/api/pkg/redis"
)

var (
	ctx         = context.Background()
	title       = "Test todo item title"
	description = "Test todo item description"
)

func Test_Integration(t *testing.T) {
	suite.Run(t, &Suite{})
}

type Suite struct {
	suite.Suite
	*require.Assertions

	todo *httpclient.Client
}

func (s *Suite) SetupSuite() {
	s.Assertions = s.Require()
	s.ResetMigration()

	c := config.Config{
		App: config.App{
			Name:    "my-app",
			Version: "test",
		},
		HTTP: httpserver.Config{
			Port: "8080",
		},
		Postgres: postgres.Config{
			Host:     "localhost",
			Port:     "5432",
			User:     "login",
			Password: "pass",
			DBName:   "postgres",
		},
		Redis: redis.Config{
			Addr: "localhost:6379",
		},
	}

	log.Logger = zerolog.Nop()

	go func() {
		err := app.Run(context.Background(), c)
		s.NoError(err)
	}()

	s.todo = httpclient.New(httpclient.Config{
		Host: "localhost",
		Port: "8080",
	})

	time.Sleep(1 * time.Second)
}

func (s *Suite) TearDownSuite() {}

func (s *Suite) SetupTest() {}

func (s *Suite) TearDownTest() {}
