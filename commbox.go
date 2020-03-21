package commbox

import (
	"bytes"
	"io"
	"io/ioutil"
	"net/http"
)

type (
	// Commbox expose the inteface to talk to Commbox.io api
	Commbox interface {
		Status() (string, error)
		Request(string, string, io.Reader) (string, error)
		CreateObject(int64, CreateObjectOptions) (string, error)
	}

	// Optons that are required to create Commbox
	Optons struct {
		Token string
	}

	cb struct {
		http  *http.Client
		token string
	}
)

// New build Commbox client
func New(opt *Optons) Commbox {
	return &cb{
		http:  &http.Client{},
		token: opt.Token,
	}
}

func (c *cb) Status() (string, error) {
	req, err := c.buildRequest("GET", systemStatus, nil)
	if err != nil {
		return "", err
	}
	data, err := c.call(req)
	if err != nil {
		return "", err
	}
	return string(data), nil
}

func (c *cb) Request(verb string, endpoint string, reader io.Reader) (string, error) {
	if reader == nil {
		reader = bytes.NewReader(nil)
	}
	data, err := ioutil.ReadAll(reader)
	if err != nil {
		return "", err
	}
	req, err := c.buildRequest(verb, endpoint, bytes.NewBuffer(data))
	if err != nil {
		return "", err
	}
	res, err := c.call(req)
	if err != nil {
		return "", err
	}
	return string(res), nil
}
