package dockerapi

import (
	"errors"
	"fmt"
	"log"
	"net/http"
	"strconv"
)

var _ = log.Print // For debugging; delete when done.

/* Start Container
 */

type ContainerHostPort struct {
	HostPort string
}
type StartContainerRequest struct {
	PortBindings map[string][]ContainerHostPort
}

// StartContainer will start up a previously created container
func (c *Client) StartContainer(id string, req StartContainerRequest) error {
	var err error

	bytes, status, err := c.post(fmt.Sprintf(RESOURCE_PATH_START_CONTAINER, id), req)
	if status == http.StatusNotFound {
		return errors.New("Container not found")
	}
	if err != nil {
		return err
	}

	if status != http.StatusNoContent {
		return errors.New("status: " + strconv.Itoa(status) + " - " + string(bytes[:]))
	}
	return nil

}
