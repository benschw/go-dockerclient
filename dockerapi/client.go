package dockerapi

import (
	"bytes"
	"encoding/json"
	"io"
	"io/ioutil"
	"log"
	"net"
	"net/http"
	"net/http/httputil"
)

const (
	RESOURCE_PATH_INSPECT          = "/containers/%s/json"
	RESOURCE_PATH_CREATE_CONTAINER = "/containers/create"
	RESOURCE_PATH_START_CONTAINER  = "/containers/%s/start"
)

func NewClient(socketPath string) *Client {

	c := new(Client)
	c.socketPath = socketPath
	return c
}

type Client struct {
	socketPath string
}

func (c *Client) get(path string) ([]byte, int, error) {
	return c.apiCall("GET", path, nil)
}

func (c *Client) post(path string, data interface{}) ([]byte, int, error) {
	return c.apiCall("POST", path, data)
}

func (c *Client) apiCall(method string, path string, data interface{}) ([]byte, int, error) {
	status := 0

	// setup request
	var params io.Reader

	if data != nil {
		buf, err := json.Marshal(data)
		if err != nil {
			return nil, status, err
		}
		params = bytes.NewBuffer(buf)
		log.Print(string(buf[:]))

	}

	req, err := http.NewRequest(method, path, params)
	if err != nil {
		return nil, status, err
	}

	if data != nil {
		req.Header.Set("Content-Type", "application/json")
	} else if method == "POST" {
		req.Header.Set("Content-Type", "plain/text")
	}

	// setup connection
	dial, err := net.Dial("unix", c.socketPath)
	if err != nil {
		return nil, status, err
	}

	// make request
	var resp *http.Response
	clientconn := httputil.NewClientConn(dial, nil)
	resp, err = clientconn.Do(req)
	defer clientconn.Close()

	// if resp.Header.Get("Content-Type") != "application/json" {
	// 	return nil, errors.New("expected application/json")
	// }

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Print(err)
		return nil, resp.StatusCode, err
	}

	return body, resp.StatusCode, nil
}
