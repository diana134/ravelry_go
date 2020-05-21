// Twitter bot written in golang
// Grabs a random interesting fact from the Raverly api and tweets every few hours

package main

import (
	"fmt"
	"log"
	"math/rand"
	"time"

	"github.com/dghubble/go-twitter/twitter"
	"github.com/spf13/viper"
)

func main() {
	err := readConfigFile()
	if err != nil { // Handle errors reading the config file
		panic(fmt.Errorf("fatal error config file %s", err))
	}

	ravelryClient := setUpRavelryClient()

	twitterClient, err := setUpTwitterClient()
	if err != nil {
		panic(fmt.Errorf("Error getting Twitter Client %s", err))
	}

	// the part that loops

	for {

		// choose what query to run
		availabilityType, sortType := chooseQuery()

		// make Ravelry request
		patternData, err := ravelryClient.PatternSearch(availabilityType, sortType)
		if err != nil {
			fmt.Errorf("Error making Ravelry request %s", err)
		}

		// generate the text for the tweet
		text := "The " + sortType.tweetText + " " + availabilityType.tweetText + " pattern right now is " + patternData["name"].(string) + ": " + "ravelry.com/patterns/library/" + patternData["permalink"].(string)
		fmt.Println(text)

		sendTweet(twitterClient, text)

		// wait for an hour
		time.Sleep(time.Hour)
	}
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

func chooseQuery() (availabilityType Parameter, sortType Parameter) {
	rand.Seed(time.Now().Unix()) // initialize global pseudo random generator

	randomIndex := rand.Intn(len(AvailabilityTypes))
	availabilityType = AvailabilityTypes[randomIndex]

	randomIndex = rand.Intn(len(SortTypes))
	sortType = SortTypes[randomIndex]

	return availabilityType, sortType
}

func sendTweet(client *twitter.Client, text string) {
	tweet, resp, err := client.Statuses.Update(text, nil)
	if err != nil {
		log.Println(err)
	}
	log.Printf("%+v\n", resp)
	log.Printf("%+v\n", tweet)
}
