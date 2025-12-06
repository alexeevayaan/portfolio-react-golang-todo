package httpclient

import (
	"context"
	"errors"
	"fmt"
	"net/http"

	http_client "github.com/alexeevayaan/portfolio-react-golang-todo/api/gen/http/todo_v1/client"
	"github.com/google/uuid"
)

func (c *Client) Create(ctx context.Context, title string, description string) (uuid.UUID, error) {
	input := http_client.CreateTodoInput{
		Title:       title,
		Description: description,
	}

	output, err := c.client.CreateTodoWithResponse(ctx, input)
	if err != nil {
		return uuid.Nil, fmt.Errorf("create todo: %w", err)
	}

	if output.StatusCode() != http.StatusOK {
		return uuid.Nil, fmt.Errorf("create todo: %w", errors.New(output.JSON400.Error))
	}

	return output.JSON200.ID, nil
}
