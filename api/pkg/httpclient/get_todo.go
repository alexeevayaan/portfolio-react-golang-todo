package httpclient

import (
	"context"
	"fmt"
	"net/http"
	"time"

	http_client "github.com/alexeevayaan/portfolio-react-golang-todo/api/gen/http/todo_v1/client"
	"github.com/google/uuid"
)

type Todo struct {
	ID          uuid.UUID `json:"id"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
	DeletedAt   time.Time `json:"deleted_at"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	Completed   bool      `json:"completed"`
}

func (c *Client) Get(ctx context.Context, id string) (*http_client.GetTodoOutput, error) {
	output, err := c.client.GetTodoByIDWithResponse(ctx, uuid.MustParse(id))
	if err != nil {
		return nil, fmt.Errorf("GetTodoByIDWithResponse: %w", err)
	}

	if output.StatusCode() == http.StatusNotFound {
		return nil, ErrNotFound
	}

	if output.StatusCode() != http.StatusOK {
		return nil, fmt.Errorf("request failed: status: %s, body:%s", output.Status(), output.Body)
	}

	return output.JSON200, nil
}
