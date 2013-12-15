package main

import (
	"flag"
	"fmt"
	"github.com/benschw/go-dockerclient/dockerapi"
	"log"
	"os"
)

func main() {
	var (
		socketPath string
	)

	flag.StringVar(&socketPath, "s", "/var/run/docker.sock", "unix socket to connect to")
	//	flag.StringVar(&containerId, "c", "", "container id")
	flag.Parse()

	c := dockerapi.NewClient(socketPath)

	// create container
	createReq := dockerapi.CreateContainerRequest{Image: "benschw/etcd"}

	resp, err := c.CreateContainer(createReq)
	if err != nil {
		log.Fatal(err)
		return
	}
	fmt.Fprintf(os.Stdout, "new container: %s\n", resp.Id)

	// start container
	ports := []dockerapi.ContainerHostPort{dockerapi.ContainerHostPort{HostPort: "11022"}}
	portsWrapper := make(map[string][]dockerapi.ContainerHostPort)
	portsWrapper["4001/tcp"] = ports

	startReq := dockerapi.StartContainerRequest{PortBindings: portsWrapper}

	if err = c.StartContainer(resp.Id, startReq); err != nil {
		log.Fatal(err)
		return
	}

	// inspect container
	container, err := c.InspectContainer(resp.Id)
	if err != nil {
		log.Fatal(err)
		return
	}

	fmt.Fprintf(os.Stdout, "%v", container)
}
