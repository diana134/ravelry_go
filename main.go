// Twitter bot written in golang
// Grabs a random interesting fact from the Raverly api and tweets every few hours

package main

import (
	"encoding/base64"
	"fmt"
	"io/ioutil"
	"net/http"

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
	authUsername := viper.GetString("authUsername")
	authPassword := viper.GetString("authPassword")

	url := "https://api.ravelry.com/projects/wool-rat/list.json"

	req, _ := http.NewRequest("GET", url, nil)

	authString := authUsername + ":" + authPassword
	req.Header.Add("Authorization", "Basic "+base64.StdEncoding.EncodeToString([]byte(authString)))

	res, _ := http.DefaultClient.Do(req)

	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)

	fmt.Println(string(body))
}
