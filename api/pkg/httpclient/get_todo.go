package httpclient

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"

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

func (c *Client) Get(ctx context.Context, id string) (Todo, error) {
	const getTodo = "v1/todo"

	path := fmt.Sprintf("http://%s/%s/%s", c.host, getTodo, id)

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, path, http.NoBody)
	if err != nil {
		return Todo{}, fmt.Errorf("http.NewRequestWithContext: %w", err)
	}

	resp, err := c.client.Do(req)
	if err != nil {
		return Todo{}, fmt.Errorf("client.Do: %w", err)
	}

	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return Todo{}, fmt.Errorf("io.ReadAll: %w", err)
	}

	if resp.StatusCode == http.StatusNotFound {
		return Todo{}, ErrNotFound
	}

	if resp.StatusCode != http.StatusOK {
		return Todo{}, fmt.Errorf("request failed with status: %s, body: %s", resp.Status, body)
	}

	var todo Todo

	if err = json.Unmarshal(body, &todo); err != nil {
		return Todo{}, fmt.Errorf("json.Unmarshal: %w", err)
	}

	return todo, nil
}
