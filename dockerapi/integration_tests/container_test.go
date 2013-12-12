package test

import (
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

func Test_CreateContainer_etcd(t *testing.T) {

	c := dockerapi.NewClient(socketPath)

	// create container
	createReq := dockerapi.CreateContainerRequest{Image: "benschw/etcd"}

	resp, err := c.CreateContainer(createReq)
	if err != nil {
		t.Error(err)
	}

	if resp.Id == "" {
		t.Errorf("response doesn't look right %v", resp)
	}
}
