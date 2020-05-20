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

// PatternSearch returns the top free Hot Right Now pattern (for now) TODO
func (c *Client) PatternSearch() (map[string]interface{}, error) {
	paramerters := "?page_size=1&availability=free&sort=recently-popular"
	data, _ := c.doRequest("https://api.ravelry.com/patterns/search.json" + paramerters)

	// result is an empty "dictionary" of "stuff"
	// unmarshal the json data into it
	// now we can get the "patterns" part of the response, which is currently in the form of a slice of "stuff"
	// we get the first "stuff" and make it into the format of an empty "dictionary" of "stuff"
	// and now we're returning a nice key/value set of data for a single pattern from Ravelry
	var result map[string]interface{}
	json.Unmarshal(data, &result)
	patterns := result["patterns"].([]interface{})
	patternsContents := patterns[0].(map[string]interface{})

	return patternsContents, nil
}

func (c *Client) doRequest(url string) ([]byte, error) {
	req, _ := http.NewRequest("GET", url, nil)
	req.Header.Add("Authorization", c.authHeader)
	res, _ := http.DefaultClient.Do(req)
	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)
	return body, nil
}
