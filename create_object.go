package commbox

import (
	"bytes"
	"encoding/json"
	"fmt"
)

type CreateObjectOptions struct {
	Data CreateObjectData `json:"data"`
}

type CreateObjectData struct {
	Type                   int64                `json:"Type"`
	UserStreamProviderID   string               `json:"UserStreamProviderId"`
	UserStreamProviderType int64                `json:"UserStreamProviderType"`
	Content                *CreateObjectContent `json:"Content,omitempty"`
	Message                string               `json:"Message"`
	User                   *CreateObjectUser    `json:"User,omitempty"`
}

type CreateObjectContent struct {
	Subject string `json:"subject"`
}

type CreateObjectUser struct {
	UniqueID  string `json:"UniqueId"`
	LastName  string `json:"LastName"`
	FirstName string `json:"FirstName"`
	Phone1    string `json:"Phone1"`
	Email     string `json:"Email"`
}

func (c *cb) CreateObject(streamID int64, opt CreateObjectOptions) (string, error) {
	data, err := json.Marshal(opt)
	if err != nil {
		return "", err
	}
	req, err := c.buildRequest("POST", fmt.Sprintf(createStreamObject, streamID), bytes.NewBuffer(data))
	if err != nil {
		return "", err
	}
	res, err := c.call(req)
	if err != nil {
		return "", err
	}
	return string(res), nil
}
