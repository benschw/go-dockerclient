package dockerapi

import (
	"encoding/json"
)

func ContainerFromJson(bytes []byte, entity *Container) error {
	if err := json.Unmarshal(bytes, &entity); err != nil {
		return err
	}
	return nil
}

/* Container type
 */
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
