package httpclient

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/google/uuid"
)

func (c *Client) Create(ctx context.Context, title string, description string) (uuid.UUID, error) {
	const createTodo = "v1/todo"

	path := fmt.Sprintf("http://%s/%s", c.host, createTodo)

	request := struct {
		Title       string `json:"title"`
		Description string `json:"description"`
	}{
		Title:       title,
		Description: description,
	}
	body, err := json.Marshal(request)
	if err != nil {
		return uuid.Nil, fmt.Errorf("json.Marshal: %w", err)
	}

	req, err := http.NewRequestWithContext(ctx, http.MethodPost, path, bytes.NewReader(body))
	if err != nil {
		return uuid.Nil, fmt.Errorf("http.NewRequestWithContext: %w", err)
	}

	req.Header.Set("Content-Type", "application/json")

	resp, err := c.client.Do(req)
	if err != nil {
		return uuid.Nil, fmt.Errorf("client.Do: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return uuid.Nil, fmt.Errorf("request failed with status: %s, body: %s", resp.Status, body)
	}

	response := struct {
		ID uuid.UUID `json:"id"`
	}{}

	err = json.NewDecoder(resp.Body).Decode(&response)
	if err != nil {
		return uuid.Nil, fmt.Errorf("json.Decoder: %w", err)
	}

	return response.ID, nil
}
