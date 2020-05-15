package main

import (
	"encoding/base64"
	"encoding/json"
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
func (c *Client) PatternSearch() (interface{}, error) {
	paramerters := "?page_size=1&availability=free&sort=recently-popular"
	data, _ := c.doRequest("https://api.ravelry.com/patterns/search.json" + paramerters)

	// fmt.Println(string(data))

	var result map[string]interface{}
	json.Unmarshal(data, &result)
	patterns := result["patterns"].([]interface{})

	// fmt.Println(patterns)

	return patterns[0], nil
}

type Pattern struct {
	name string
}

type PatternList struct {
	patterns []string
}

func (c *Client) doRequest(url string) ([]byte, error) {
	req, _ := http.NewRequest("GET", url, nil)
	req.Header.Add("Authorization", c.authHeader)
	res, _ := http.DefaultClient.Do(req)
	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)
	return body, nil
}
