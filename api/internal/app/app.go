package app

import (
	"context"

	"github.com/alexeevayaan/portfolio-react-golang-todo/api/config"
	"github.com/alexeevayaan/portfolio-react-golang-todo/api/pkg/router"
)


func Run(ctx context.Context, c config.Config) error{

	r:=router.New()
	

	return nil
}