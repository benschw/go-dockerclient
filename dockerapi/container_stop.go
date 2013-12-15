package dockerapi

import (
	"errors"
	"fmt"
	"log"
	"net/http"
	"strconv"
)

var _ = log.Print // For debugging; delete when done.

func (c *Client) StopContainer(id string, timeout int) error {
	var err error

	bytes, status, err := c.post(fmt.Sprintf(RESOURCE_PATH_STOP_CONTAINER, id, strconv.Itoa(timeout)), nil)
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
