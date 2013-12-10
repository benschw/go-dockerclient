package main

import (
	"flag"
	"fmt"
	"os"
	"github.com/benschw/go-dockerclient/dockerapi"
	"log"
)

func main() {
	var (
		socketPath  string
		containerId string
	)

	flag.StringVar(&socketPath, "s", "/var/run/docker.sock", "unix socket to connect to")
	flag.StringVar(&containerId, "c", "", "container id")
	flag.Parse()

	c := dockerapi.NewClient(socketPath)

	container, err := c.Inspect(containerId)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Fprintf(os.Stdout, "%v", container)

}
