package gowitness

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"strconv"
	"wsw/backend/domain/url"
	"wsw/backend/ent"
)

type (
	Client interface {
		AddURL(string) (int, error)
		Details(*ent.Url) (*DetailsURL, error)
	}
	DetailsURL struct {
		ID     int
		Image  string
		Status url.Status
	}
	addURLResponse  struct{ ID int }
	detailsResponse struct {
		ID       int
		Title    string
		Filename string
	}

	clientImpl struct {
		baseURL   string
		imageHost string
		client    http.Client
	}
)

// Details implements Client.
func (c *clientImpl) Details(url *ent.Url) (*DetailsURL, error) {
	if url.APIURLID != nil {
		response, err := http.Get(c.baseURL + "/detail/" + strconv.FormatInt(int64(*url.APIURLID), 10))
		if err != nil {
			return nil, err
		}
		if response.Body != nil {
			defer response.Body.Close()
		}
		result, readErr := io.ReadAll(response.Body)
		if readErr != nil {
			return nil, readErr
		}

		responseType := detailsResponse{}
		jsonErr := json.Unmarshal(result, &responseType)
		if jsonErr != nil {
			return nil, jsonErr
		}

		return c.createDetailsWithScreenshot(url.ID, responseType), nil
	}
	return c.createDetailsForNewUrl(url.ID), nil
}

// AddURL implements Client.
func (c *clientImpl) AddURL(url string) (int, error) {
	postBody, _ := json.Marshal(map[string]string{
		"url":        url,
		"oneshot":    "false",
		"foreground": "true",
	})
	response, err := http.Post(c.baseURL+"/screenshot", "application/json", bytes.NewBuffer(postBody))
	if err != nil {
		return 0, err
	}
	if response.Body != nil {
		defer response.Body.Close()
	}
	result, readErr := io.ReadAll(response.Body)
	if readErr != nil {
		return 0, readErr
	}
	responseType := addURLResponse{}
	jsonErr := json.Unmarshal(result, &responseType)
	if jsonErr != nil {
		return 0, jsonErr
	}
	return responseType.ID, nil
}

func (c *clientImpl) createDetailsWithScreenshot(id int, response detailsResponse) *DetailsURL {
	return &DetailsURL{
		ID:     id,
		Image:  c.imageHost + "/screenshots/" + response.Filename,
		Status: url.Success,
	}
}

func (c *clientImpl) createDetailsForNewUrl(id int) *DetailsURL {
	return &DetailsURL{ID: id, Image: c.imageHost + "/assets/loader-200px-200px.gif", Status: url.Pending}
}
func NewClient(client http.Client, baseURL string, imageHost string) Client {
	return &clientImpl{client: client, baseURL: baseURL, imageHost: imageHost}
}
