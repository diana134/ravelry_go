// Twitter bot written in golang
// Grabs a random interesting fact from the Raverly api and tweets every few hours

package main

import (
	"fmt"

	"github.com/spf13/viper"
)

func main() {
	// read the config file
	viper.SetConfigName("config")
	viper.SetConfigType("json")
	viper.AddConfigPath(".")

	err := viper.ReadInConfig() // Find and read the config file
	if err != nil {             // Handle errors reading the config file
		panic(fmt.Errorf("fatal error config file %s", err))
	}

	// set up the Ravelry client
	ravelryCredentials := RavelryCredentials{
		AuthUsername: viper.GetString("authUsername"),
		AuthPassword: viper.GetString("authPassword"),
	}
	ravelryClient := GetRavelryClient(&ravelryCredentials)

	// set up the Twitter client
	// twitterCredentials := Credentials{
	// 	ConsumerKey:       viper.GetString("apiKey"),
	// 	ConsumerSecret:    viper.GetString("apiSecretKey"),
	// 	AccessToken:       viper.GetString("accessToken"),
	// 	AccessTokenSecret: viper.GetString("accessTokenSecret"),
	// }
	// twitterClient, err := GetClient(&twitterCredentials)
	// if err != nil {
	// 	log.Println("Error getting Twitter Client")
	// 	log.Println(err)
	// }

	// make Ravelry request
	patternData, _ := ravelryClient.PatternSearch()
	fmt.Println("the hottest pattern right now is", patternData["name"])

	// tweet, resp, err := twitterClient.Statuses.Update("Hello World!", nil)
	// if err != nil {
	// 	log.Println(err)
	// }
	// log.Printf("%+v\n", resp)
	// log.Printf("%+v\n", tweet)
}
