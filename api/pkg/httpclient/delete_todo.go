package httpclient

import (
	"context"
	"fmt"
	"net/http"

	"github.com/google/uuid"
)

func (c *Client) Delete(ctx context.Context, id string) error {
	output, err := c.client.DeleteTodoByIDWithResponse(ctx, uuid.MustParse(id))
	if err != nil {
		return fmt.Errorf("delete todo: %w", err)
	}

	if output.StatusCode() != http.StatusNoContent {
		return fmt.Errorf("request failed: status: %s, body:%s", output.Status(), output.Body)
	}

	return nil
}
