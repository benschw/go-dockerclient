package dockerapi

import (
	"log"
	"os"
	"testing"
)

var _ = log.Print // For debugging; delete when done.
var _ = os.Stdout // For debugging; delete when done.

var (
	socketPath = "/var/run/docker.sock"
)

func Test_NewClient(t *testing.T) {

	c := NewClient(socketPath)

	if c.socketPath != socketPath {
		t.Errorf("field 'socketPath' not setup correctly")
	}
}
