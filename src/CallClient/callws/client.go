package callws

import (
	"CallClient/model"
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/url"
)

type Client struct {
	BaseURL    *url.URL
	httpClient *http.Client
}

func NewClient(host, port, scheme string) *Client {
	return &Client{
		BaseURL:    &url.URL{Scheme: scheme, Host: fmt.Sprintf("%s:%s", host, port)},
		httpClient: http.DefaultClient,
	}
}

func (c *Client) newRequest(method, path string, body interface{}, responseContent interface{},
	queryParameters map[string]string) (*http.Response, error) {

	relativeUrl := &url.URL{Path: path}

	for paramKey, paramValue := range queryParameters {
		relativeUrl.Query().Add(paramKey, paramValue)
	}

	absUrl := c.BaseURL.ResolveReference(relativeUrl)
	var buf io.ReadWriter

	if body != nil {
		buf = new(bytes.Buffer)
		err := json.NewEncoder(buf).Encode(body)
		if err != nil {
			return nil, err
		}
	}
	request, err := http.NewRequest(method, absUrl.String(), buf)

	if err != nil {
		return nil, err
	}

	if body != nil {
		request.Header.Set("Content-Type", "application/json")
	}

	response, err := c.httpClient.Do(request)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()
	if responseContent != nil {
		err = json.NewDecoder(response.Body).Decode(responseContent)
	}
	return response, err
}

func (c *Client) AddCalls(callList []model.Call) (bool, error) {
	response, err := c.newRequest("PUT", "/call", callList, nil, map[string]string{})
	if err != nil {
		return false, err
	} else if response.StatusCode >= 200 && response.StatusCode < 300 {
		return response.StatusCode == http.StatusCreated, nil
	} else {
		return false, errors.New(fmt.Sprintf("Request Failed: Code: %d", response.StatusCode))
	}
}

func (c *Client) GetCalls(page int, pageSize int, filter model.Filter) (model.CallQueryResult, error) {
	var queryResult model.CallQueryResult
	paramMap := filter.ParamMap
	paramMap["page"] = fmt.Sprint(page)
	paramMap["pageSize"] = fmt.Sprint(pageSize)
	response, err := c.newRequest("GET", "/call", nil, &queryResult, paramMap)
	if err != nil {
		return queryResult, err
	} else if response.StatusCode >= 200 && response.StatusCode < 300 {
		return queryResult, nil
	} else {
		return queryResult, errors.New(fmt.Sprintf("Request Failed: Code: %d", response.StatusCode))
	}
}

func (c *Client) RemoveCalls(filter model.Filter) (bool, error) {
	response, err := c.newRequest("DELETE", "/call", nil, nil, filter.ParamMap)
	if err != nil {
		return false, err
	} else if response.StatusCode >= 200 && response.StatusCode < 300 {
		return response.StatusCode == http.StatusOK, nil
	} else {
		return false, errors.New(fmt.Sprintf("Request Failed: Code: %d", response.StatusCode))
	}
}

func (c *Client) GetMetadata() ([]model.CallMetadata, error) {
	var callMetaData []model.CallMetadata
	response, err := c.newRequest("GET", "/metadata", nil, &callMetaData, map[string]string{})
	if err != nil {
		return callMetaData, err
	} else if response.StatusCode >= 200 && response.StatusCode < 300 {
		return callMetaData, nil
	} else {
		return callMetaData, errors.New(fmt.Sprintf("Request Failed: Code: %d", response.StatusCode))
	}
}
