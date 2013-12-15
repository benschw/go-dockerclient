package dockerapi

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"
)

var _ = log.Print // For debugging; delete when done.
var _ = fmt.Print // For debugging; delete when done.

/* Create Container
 */

type CreateContainerRequest struct {
	Image string
	// Hostname     string
	// User         string
	// Memory       string
	// MemorySwap   string
	// AttachStdin  bool
	// AttachStdout bool
	// AttachStderr bool
	// //     PortSpecs null,
	// Privileged   bool
	// Tty          bool
	// OpenStdin    bool
	// StdinOnce    bool
	// Env          []string
	// Cmd          []string
	// Dns          []string
	// // Volumes {}
	// // VolumesFrom,
	// WorkingDir   string
}

func CreateContainerResponseFromJson(bytes []byte, entity *CreateContainerResponse) error {
	if err := json.Unmarshal(bytes, &entity); err != nil {
		return err
	}
	return nil
}

type CreateContainerResponse struct {
	Id       string
	Warnings []string
}

// CreateContainer will create a container
func (c *Client) CreateContainer(req CreateContainerRequest) (CreateContainerResponse, error) {
	var err error
	var entity CreateContainerResponse

	bytes, status, err := c.post(RESOURCE_PATH_CREATE_CONTAINER, req)
	if status == http.StatusNotFound {
		return entity, errors.New("Image not found")
	}
	if err != nil {
		return entity, err
	}
	if err = CreateContainerResponseFromJson(bytes, &entity); err != nil {
		return entity, err
	}

	return entity, nil
}
