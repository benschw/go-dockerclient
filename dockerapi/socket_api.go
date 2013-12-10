package dockerapi

import (
	"errors"
	"io/ioutil"
	"net"
	"net/http"
	"net/http/httputil"
)


func apiGet(socketPath string, path string) ([]byte, error) {
	req, err := http.NewRequest("GET", path, nil)
	if err != nil {
		return nil, err
	}
	dial, err := net.Dial("unix", socketPath)
	if err != nil {
		return nil, err
	}

	var resp *http.Response
	clientconn := httputil.NewClientConn(dial, nil)
	resp, err = clientconn.Do(req)
	defer clientconn.Close()

	if resp.StatusCode < 200 || resp.StatusCode >= 400 {
		return nil, errors.New("bad status code")
	}


	if resp.Header.Get("Content-Type") != "application/json" {
		return nil, errors.New("expected application/json")
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return body, nil
}


