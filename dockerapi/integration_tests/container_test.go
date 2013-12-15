package test

import (
	"errors"
	"github.com/benschw/go-dockerclient/dockerapi"
	"log"
	"os"
	"testing"
)

var _ = log.Print // For debugging; delete when done.
var _ = os.Stdout // For debugging; delete when done.

var (
	socketPath = "/var/run/docker.sock"
)

func Test_CreateContainer(t *testing.T) {
	// given
	c := dockerapi.NewClient(socketPath)

	createReq := dockerapi.CreateContainerRequest{Image: "benschw/etcd"}

	// when
	resp, err := c.CreateContainer(createReq)

	// then
	if err != nil {
		t.Error(err)
	}

	if resp.Id == "" {
		t.Errorf("response doesn't look right %+v", resp)
	}
}

func Test_StartContainer(t *testing.T) {
	// given
	id, err := createContainer("benschw/etcd")
	if err != nil {
		t.Error(err)
	}

	c := dockerapi.NewClient(socketPath)

	ports := []dockerapi.ContainerHostPort{dockerapi.ContainerHostPort{HostPort: "11022"}}
	portsWrapper := make(map[string][]dockerapi.ContainerHostPort)
	portsWrapper["4001/tcp"] = ports

	startReq := dockerapi.StartContainerRequest{PortBindings: portsWrapper}

	// when
	err = c.StartContainer(id, startReq)
	defer stopContainer(id)

	// then
	if err != nil {
		t.Fatal(err)
	}
}

func Test_StopContainer(t *testing.T) {
	// given
	id, err := createContainer("benschw/etcd")
	if err != nil {
		t.Error(err)
	}

	if err := startEtcdContainer(id); err != nil {
		t.Error(err)
	}

	c := dockerapi.NewClient(socketPath)

	// when
	err = c.StopContainer(id, 1)

	// then
	if err != nil {
		t.Error(err)
	}
}

func Test_RemoveContainer(t *testing.T) {
	// given
	id, err := createContainer("benschw/etcd")
	if err != nil {
		t.Error(err)
	}

	c := dockerapi.NewClient(socketPath)

	// when
	err = c.RemoveContainer(id)

	// then
	if err != nil {
		t.Error(err)
	}

	if err := startEtcdContainer(id); err == nil {
		t.Error("container should be removed and not possible to start")
	}
}

func Test_InspectContainer(t *testing.T) {
	// given
	id, err := createContainer("benschw/etcd")
	if err != nil {
		t.Error(err)
	}

	if err := startEtcdContainer(id); err != nil {
		t.Error(err)
	}
	defer stopContainer(id)

	c := dockerapi.NewClient(socketPath)

	// when
	container, err := c.InspectContainer(id)

	// then
	if err != nil {
		t.Error(err)
	}
	if container.ID == "" {
		t.Error("field 'ID' is empty")
	}

	if container.Name == "" {
		t.Error("field 'Name' is empty")
	}

	// log.Printf("%+v", container)
}

/* Helper functions
 */

func createContainer(image string) (string, error) {
	id := ""

	c := dockerapi.NewClient(socketPath)

	// create container
	createReq := dockerapi.CreateContainerRequest{Image: image}

	resp, err := c.CreateContainer(createReq)
	if err != nil {
		return id, err
	}

	if resp.Id == "" {
		return id, errors.New("no id to return")
	}
	return resp.Id, nil
}

func startEtcdContainer(id string) error {
	c := dockerapi.NewClient(socketPath)

	// start container
	ports := []dockerapi.ContainerHostPort{dockerapi.ContainerHostPort{HostPort: "11022"}}
	portsWrapper := make(map[string][]dockerapi.ContainerHostPort)
	portsWrapper["4001/tcp"] = ports

	startReq := dockerapi.StartContainerRequest{PortBindings: portsWrapper}

	if err := c.StartContainer(id, startReq); err != nil {
		return err
	}
	return nil
}

func stopContainer(id string) error {
	c := dockerapi.NewClient(socketPath)

	if err := c.StopContainer(id, 1); err != nil {
		return err
	}
	return nil
}
