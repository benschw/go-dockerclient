package dockerapi

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"
	"strconv"
)

func ContainerFromJson(bytes []byte, entity *Container) error {
	if err := json.Unmarshal(bytes, &entity); err != nil {
		return err
	}
	return nil
}

// Inspect Container
type Container struct {
	ID              string
	Created         string
	Path            string
	Args            []string
	Name            string
	Image           string
	NetworkSettings NetworkSettings
	Volumes         map[string]string
}

type NetworkSettings struct {
	IPAddress string
	Gateway   string
	Ports     map[string][]Address
}

type Address struct {
	HostIp   string
	HostPort string
}

// Inspect will get details about a running container
func (c *Client) Inspect(id string) (Container, error) {
	var err error

	var entity Container

	bytes, err := apiGet(c.socketPath, fmt.Sprintf(RESOURCE_PATH_INSPECT, id))
	if err != nil {
		return entity, err
	}

	if err = ContainerFromJson(bytes, &entity); err != nil {
		return entity, err
	}

	return entity, nil
}

/* Create Container */

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
func (c *Client) CreateContainer(data interface{}) (CreateContainerResponse, error) {
	var err error
	var entity CreateContainerResponse

	bytes, status, err := apiPost(c.socketPath, RESOURCE_PATH_CREATE_CONTAINER, data)
	if status == http.StatusNotFound {
		return entity, errors.New("Image not found")
	}
	if err != nil {
		return entity, err
	}
	if err = CreateContainerResponseFromJson(bytes, &entity); err != nil {
		log.Print(entity)
		return entity, err
	}

	return entity, nil
}

type ContainerHostPort struct {
	HostPort string
}
type StartContainerRequest struct {
	PortBindings map[string][]ContainerHostPort
}

// StartContainer will start up a previously created container
func (c *Client) StartContainer(id string, data StartContainerRequest) error {
	var err error

	bytes, status, err := apiPost(c.socketPath, fmt.Sprintf(RESOURCE_PATH_START_CONTAINER, id), data)
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
