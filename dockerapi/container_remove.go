package dockerapi

import (
	"errors"
	"fmt"
	"log"
	"net/http"
	"strconv"
)

var _ = log.Print // For debugging; delete when done.

func (c *Client) RemoveContainer(id string) error {
	var err error

	bytes, status, err := c.delete(fmt.Sprintf(RESOURCE_PATH_REMOVE_CONTAINER, id))
	if err != nil {
		return err
	}
	if status == http.StatusNotFound {
		return errors.New("Container not found")
	}
	if status != http.StatusNoContent {
		return errors.New("status: " + strconv.Itoa(status) + " - " + string(bytes[:]))
	}
	return nil
}
