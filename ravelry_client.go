package main

import (
	"encoding/base64"
	"io/ioutil"
	"net/http"
)

// RavelryCredentials contains the username and password for authentication
type RavelryCredentials struct {
	AuthUsername string
	AuthPassword string
}

// Client is the Ravelry http client
type Client struct {
	authString string
	authHeader string
}

// GetRavelryClient returns a Ravelry client
func GetRavelryClient(creds *RavelryCredentials) (client *Client) {
	authString := creds.AuthUsername + ":" + creds.AuthPassword
	return &Client{
		authString: authString,
		authHeader: "Basic " + base64.StdEncoding.EncodeToString([]byte(authString)),
	}
}

// PatternSearch returns the top Hot Right Now pattern (for now) TODO
func (c *Client) PatternSearch() (string, error) {
	data, _ := c.doRequest("https://api.ravelry.com/patterns/search.json")

	return data, nil
}

func (c *Client) doRequest(url string) (string, error) {
	req, _ := http.NewRequest("GET", url, nil)
	req.Header.Add("Authorization", c.authHeader)
	res, _ := http.DefaultClient.Do(req)
	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)
	return string(body), nil
}
