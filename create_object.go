package commbox

import (
	"bytes"
	"encoding/json"
	"fmt"
)

type (
	CreateObjectRequest struct {
		Data CreateObjectData `json:"data"`
	}
	CreateObjectData struct {
		Type                   int64                `json:"Type"`
		UserStreamProviderID   string               `json:"UserStreamProviderId"`
		UserStreamProviderType int64                `json:"UserStreamProviderType"`
		Content                *CreateObjectContent `json:"Content,omitempty"`
		Message                string               `json:"Message"`
		User                   *CreateObjectUser    `json:"User,omitempty"`
	}
	CreateObjectContent struct {
		Subject string `json:"subject"`
	}
	CreateObjectUser struct {
		UniqueID  string `json:"UniqueId"`
		LastName  string `json:"LastName"`
		FirstName string `json:"FirstName"`
		Phone     string `json:"Phone1"`
		Email     string `json:"Email"`
	}
	CreateObjectResponse struct {
		Status       int64                    `json:"status"`
		Description  string                   `json:"description"`
		ResponseTime string                   `json:"response_time"`
		Data         CreateObjectResponseData `json:"data"`
	}
	CreateObjectResponseData struct {
		ID int64 `json:"Id"`
	}
)

func (c *cb) CreateObject(streamID int64, opt CreateObjectRequest) (CreateObjectResponse, error) {
	response := CreateObjectResponse{}
	data, err := json.Marshal(opt)
	if err != nil {
		return response, err
	}
	req, err := c.buildRequest("POST", fmt.Sprintf(createStreamObject, streamID), bytes.NewBuffer(data))
	if err != nil {
		return response, err
	}
	res, err := c.call(req)
	if err != nil {
		return response, err
	}
	if err := json.Unmarshal(res, &response); err != nil {
		return response, err
	}
	return response, nil
}
