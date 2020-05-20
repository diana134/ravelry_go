// Twitter bot written in golang
// Grabs a random interesting fact from the Raverly api and tweets every few hours

package main

import (
	"fmt"
	"log"

	"github.com/dghubble/go-twitter/twitter"
	"github.com/spf13/viper"
)

func main() {
	err := readConfigFile()
	if err != nil { // Handle errors reading the config file
		panic(fmt.Errorf("fatal error config file %s", err))
	}

	ravelryClient := setUpRavelryClient()

	// twitterClient, err := setUpTwitterClient()
	if err != nil {
		panic(fmt.Errorf("Error getting Twitter Client %s", err))
	}

	// make Ravelry request
	patternData, err := ravelryClient.PatternSearch()
	if err != nil {
		panic(fmt.Errorf("Error making Ravelry request %s", err))
	}
	fmt.Println("the hottest free pattern right now is", patternData["name"], ": ", patternData["permalink"])

	// sendTweet(twitterClient, text)
}

func readConfigFile() error {
	viper.SetConfigName("config")
	viper.SetConfigType("json")
	viper.AddConfigPath(".")

	err := viper.ReadInConfig() // Find and read the config file

	return err
}

func setUpRavelryClient() *Client {
	ravelryCredentials := RavelryCredentials{
		AuthUsername: viper.GetString("authUsername"),
		AuthPassword: viper.GetString("authPassword"),
	}
	ravelryClient := GetRavelryClient(&ravelryCredentials)

	return ravelryClient
}

func setUpTwitterClient() (*twitter.Client, error) {
	twitterCredentials := Credentials{
		ConsumerKey:       viper.GetString("apiKey"),
		ConsumerSecret:    viper.GetString("apiSecretKey"),
		AccessToken:       viper.GetString("accessToken"),
		AccessTokenSecret: viper.GetString("accessTokenSecret"),
	}
	twitterClient, err := GetClient(&twitterCredentials)

	return twitterClient, err
}

func sendTweet(client *twitter.Client, text string) {
	tweet, resp, err := client.Statuses.Update(text, nil)
	if err != nil {
		log.Println(err)
	}
	log.Printf("%+v\n", resp)
	log.Printf("%+v\n", tweet)
}
