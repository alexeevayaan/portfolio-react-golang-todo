package httpclient

import (
	"errors"
	"fmt"

	http_client "github.com/alexeevayaan/portfolio-react-golang-todo/api/gen/http/todo_v1/client"
)

var ErrNotFound = errors.New("not found")

type Config struct {
	Host string `default:"localhost" envconfig:"HTTP_CLIENT_HOST"`
	Port string `default:"8080" envconfig:"HTTP_CLIENT_PORT"`
}

type Client struct {
	client *http_client.ClientWithResponses
}

func New(c Config) (*Client, error) {
	baseUrl := fmt.Sprintf("http://%s:%s/v1", c.Host, c.Port)

	client, err := http_client.NewClientWithResponses(baseUrl)
	if err != nil {
		return nil, fmt.Errorf("http_client.NewClientWithResponses: %w", err)
	}
	return &Client{
		client: client,
	}, nil
}
