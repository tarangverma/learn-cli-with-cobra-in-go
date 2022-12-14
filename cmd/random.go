/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/spf13/cobra"
)

// randomCmd represents the random command
var randomCmd = &cobra.Command{
	Use:   "random",
	Short: "command to run joke",
	Long:  "This command runs a joke every time triggered",
	Run: func(cmd *cobra.Command, args []string) {
		getRandomJoke()
	},
}

func init() {
	rootCmd.AddCommand(randomCmd)

}

type Joke struct {
	Id     string `json:"id"`
	Joke   string `json:"joke"`
	Status int    `json:"status"`
}

func getRandomJoke() {
	url := "https://icanhazdadjoke.com/"
	responseBytes := getJokeData(url)
	joke := Joke{}

	if err := json.Unmarshal(responseBytes, &joke); err != nil {
		log.Printf("Could not unmarshal response -%v", err)
	}
	fmt.Println(string(joke.Joke))
}

func getJokeData(baseAPI string) []byte {
	request, err := http.NewRequest(
		http.MethodGet,
		baseAPI,
		nil,
	)
	if err != nil {
		log.Printf("Could not access dad joke -%v", err)
	}

	request.Header.Add("Accept", "application/json")
	request.Header.Add("User-Agent", "DadJoke CLI (examples/go-cli)")

	response, err := http.DefaultClient.Do(request)
	if err != nil {
		log.Printf("Could not make a request -%v", err)
	}
	responseBytes, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Printf("Could not read response -%v", err)
	}
	return responseBytes
}
