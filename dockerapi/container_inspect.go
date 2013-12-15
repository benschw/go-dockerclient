package dockerapi

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"
)

var _ = log.Print // For debugging; delete when done.

// InspectContainer will get details about a running container
func (c *Client) InspectContainer(id string) (Container, error) {
	var err error

	var entity Container

	bytes, status, err := c.get(fmt.Sprintf(RESOURCE_PATH_INSPECT, id))
	if status == http.StatusNotFound {
		return entity, errors.New("Container not found")
	}
	if err != nil {
		return entity, err
	}

	if err = containerFromJson(bytes, &entity); err != nil {
		return entity, err
	}

	return entity, nil
}

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

func containerFromJson(bytes []byte, entity *Container) error {
	if err := json.Unmarshal(bytes, &entity); err != nil {
		return err
	}
	return nil
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
