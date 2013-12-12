package dockerapi

import (
	"testing"
)

func TestInspect(t *testing.T) {
	bytes := []byte(expected)

	var entity Container

	if err := containerFromJson(bytes, &entity); err != nil {
		t.Errorf("Problem mapping json to Container structure")
	}

	if entity.ID != "e8488a1d0b294c9e1221181e2187ce2d3589a781d3ff67b3da2f3e501425d138" {
		t.Errorf("ID field didn't get mapped correctly")
	}
}

var expected = `{
    "ID": "e8488a1d0b294c9e1221181e2187ce2d3589a781d3ff67b3da2f3e501425d138",
    "Created": "2013-12-10T17:34:04.767878929Z",
    "Path": "/opt/inspector",
    "Args": [
        "-s",
        "/docker/docker.sock"
    ],
    "Config": {
        "Hostname": "e8488a1d0b29",
        "Domainname": "",
        "User": "",
        "Memory": 0,
        "MemorySwap": 0,
        "CpuShares": 0,
        "AttachStdin": false,
        "AttachStdout": false,
        "AttachStderr": false,
        "PortSpecs": null,
        "ExposedPorts": {
            "8080/tcp": {}
        },
        "Tty": false,
        "OpenStdin": false,
        "StdinOnce": false,
        "Env": [
            "HOME=/",
            "PATH=/usr/local/sbin:/usr/local/bin:/usr/sbin:/usr/bin:/sbin:/bin"
        ],
        "Cmd": [],
        "Dns": null,
        "Image": "benschw/inspector",
        "Volumes": {
            "/docker": {}
        },
        "VolumesFrom": "",
        "WorkingDir": "",
        "Entrypoint": [
            "/opt/inspector",
            "-s",
            "/docker/docker.sock"
        ],
        "NetworkDisabled": false
    },
    "State": {
        "Running": true,
        "Pid": 31638,
        "ExitCode": 0,
        "StartedAt": "2013-12-10T17:34:04.783907206Z",
        "FinishedAt": "0001-01-01T00:00:00Z",
        "Ghost": true
    },
    "Image": "fb641c41b508466693b64a376e7ba11443fa26277960340dc5186eeba68bfb55",
    "NetworkSettings": {
        "IPAddress": "172.17.0.25",
        "IPPrefixLen": 16,
        "Gateway": "172.17.42.1",
        "Bridge": "docker0",
        "PortMapping": null,
        "Ports": {
            "8080/tcp": null
        }
    },
    "SysInitPath": "/usr/bin/docker",
    "ResolvConfPath": "/var/lib/docker/containers/e8488a1d0b294c9e1221181e2187ce2d3589a781d3ff67b3da2f3e501425d138/resolv.conf",
    "HostnamePath": "/var/lib/docker/containers/e8488a1d0b294c9e1221181e2187ce2d3589a781d3ff67b3da2f3e501425d138/hostname",
    "HostsPath": "/var/lib/docker/containers/e8488a1d0b294c9e1221181e2187ce2d3589a781d3ff67b3da2f3e501425d138/hosts",
    "Name": "/sad_turing",
    "Driver": "aufs",
    "Volumes": {
        "/docker": "/var/run/"
    },
    "VolumesRW": {
        "/docker": true
    }
}`
