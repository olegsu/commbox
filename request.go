package commbox

import (
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"path"
)

const (
	baseurl = "api.commbox.io"
	scheme  = "https"
)

func (c *cb) buildRequest(verb string, endpoint string, data io.Reader) (*http.Request, error) {
	u, err := url.Parse(path.Join(baseurl, endpoint))
	u.Scheme = scheme
	if err != nil {
		return nil, err
	}
	q := url.Values{}
	q.Add("access_token", c.token)
	u.RawQuery = q.Encode()
	req, err := http.NewRequest(verb, u.String(), data)
	req.Header.Add("Content-Type", "application/json")
	if err != nil {
		return nil, err
	}
	return req, nil
}

func (c *cb) call(req *http.Request) ([]byte, error) {
	resp, err := c.http.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	return data, nil
}
