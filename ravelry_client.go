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

// Parameter adjusts the query and the tweet text
type Parameter struct {
	urlKey    string
	urlValue  string
	tweetText string
}

// SortTypes contains all the different types of sort for queries
var SortTypes = [15]Parameter{
	Parameter{
		urlKey:    "sort",
		urlValue:  "recently-popular",
		tweetText: "hottest",
	},
	Parameter{
		urlKey:    "sort",
		urlValue:  "created",
		tweetText: "newest",
	},
	Parameter{
		urlKey:    "sort",
		urlValue:  "created_",
		tweetText: "oldest",
	},
	Parameter{
		urlKey:    "sort",
		urlValue:  "popularity",
		tweetText: "most popular",
	},
	Parameter{
		urlKey:    "sort",
		urlValue:  "popularity_",
		tweetText: "least popular",
	},
	Parameter{
		urlKey:    "sort",
		urlValue:  "projects",
		tweetText: "most made",
	},
	Parameter{
		urlKey:    "sort",
		urlValue:  "projects_",
		tweetText: "least made",
	},
	Parameter{
		urlKey:    "sort",
		urlValue:  "favorites",
		tweetText: "most favorited",
	},
	Parameter{
		urlKey:    "sort",
		urlValue:  "favorites_",
		tweetText: "least favorited",
	},
	Parameter{
		urlKey:    "sort",
		urlValue:  "queues",
		tweetText: "most queued",
	},
	Parameter{
		urlKey:    "sort",
		urlValue:  "queues_",
		tweetText: "least queued",
	},
	Parameter{
		urlKey:    "sort",
		urlValue:  "date",
		tweetText: "most recently published",
	},
	Parameter{
		urlKey:    "sort",
		urlValue:  "date_",
		tweetText: "oldest published",
	},
	Parameter{
		urlKey:    "sort",
		urlValue:  "rating",
		tweetText: "highest rated",
	},
	Parameter{
		urlKey:    "sort",
		urlValue:  "rating_",
		tweetText: "lowest rated",
	},
}

// AvailabilityTypes contains the different types of availability for queries
var AvailabilityTypes = [3]Parameter{
	Parameter{
		urlKey:    "availability",
		urlValue:  "free",
		tweetText: "free",
	},
	Parameter{
		urlKey:    "availability",
		urlValue:  "-free",
		tweetText: "paid",
	},
	Parameter{
		urlKey:    "availability",
		urlValue:  "discontinued",
		tweetText: "discontinued",
	},
}

// PatternBaseURL is the url that takes a suffix of query parameters
const PatternBaseURL = "ravelry.com/patterns/library/"

// GetRavelryClient returns a Ravelry client
func GetRavelryClient(creds *RavelryCredentials) (client *Client) {
	authString := creds.AuthUsername + ":" + creds.AuthPassword

	return &Client{
		authString: authString,
		authHeader: "Basic " + base64.StdEncoding.EncodeToString([]byte(authString)),
	}
}

// BuildParameterString builds the parameter string for the query
func BuildParameterString(availabilityType Parameter, sortType Parameter) string {
	parameterString := "?page_size=1"
	parameterString += "&" + availabilityType.urlKey + "=" + availabilityType.urlValue + "&" + sortType.urlKey + "=" + sortType.urlValue
	return parameterString
}

// PatternSearch returns the top free Hot Right Now pattern (for now) TODO
func (c *Client) PatternSearch(availabilityType Parameter, sortType Parameter) (map[string]interface{}, error) {
	parameters := BuildParameterString(availabilityType, sortType)
	data, err := c.doRequest("https://api.ravelry.com/patterns/search.json" + parameters)
	if err != nil {
		return nil, err
	}

	// result is an empty "dictionary" of "stuff"
	// unmarshal the json data into it
	// now we can get the "patterns" part of the response, which is currently in the form of a slice of "stuff"
	// we get the first "stuff" and make it into the format of an empty "dictionary" of "stuff"
	// and now we're returning a nice key/value set of data for a single pattern from Ravelry
	var result map[string]interface{}
	err = json.Unmarshal(data, &result)
	if err != nil {
		return nil, err
	}
	patterns := result["patterns"].([]interface{})
	patternsContents := patterns[0].(map[string]interface{})

	return patternsContents, nil
}

func (c *Client) doRequest(url string) ([]byte, error) {
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	req.Header.Add("Authorization", c.authHeader)
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}

	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	return body, nil
}
