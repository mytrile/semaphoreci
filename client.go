package semaphoreci

import (
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
)

const (
	api_base = "https://semaphoreapp.com/api/v1"
)

type Client struct {
	auth_token string
	client     *http.Client
}

func NewClient(auth_token string) *Client {
	return &Client{auth_token, new(http.Client)}
}

func (c *Client) GetRequest(urlString string) ([]byte, error) {
	url := fmt.Sprintf("%s/%s?auth_token=%v", api_base, urlString, c.auth_token)
	req, err := http.NewRequest("GET", url, nil)
	resp, err := c.client.Do(req)
	if err != nil {
		return make([]byte, 0), err
	}
	if resp.StatusCode != 200 {
		return make([]byte, 0), errors.New(fmt.Sprintf("Got a %v status code on fetch of %v", resp.StatusCode, urlString))
	}
	body, err := ioutil.ReadAll(resp.Body)
	return body, err
}
