package dockerapi

import (
	"fmt"
	"encoding/json"
)


const (
	RESOURCE_PATH_INSPECT = "/containers/%s/json"
)


func NewClient(socketPath string) *Client {

	c := new(Client)
	c.socketPath = socketPath
	return c
}

type Client struct {
	socketPath string
}

func (c *Client) Inspect(containerId string) (Container, error) {
	var err error

	var entity Container

	bytes, err := apiGet(c.socketPath, fmt.Sprintf(RESOURCE_PATH_INSPECT, containerId))
	if err != nil {
		return entity, err
	}

	if err = json.Unmarshal(bytes, &entity); err != nil {
		return entity, err
	}

	return entity, nil
}



