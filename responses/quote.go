package responses

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/spf13/cobra"
)

var baseURL string = "https://programming-quotes-api.herokuapp.com"

// Quote needs a comment so you don't annoy me
type Quote struct {
	Text   string `json:"en"`
	Author string `json:"author"`
}

// GetRandomQuote gets a quote
var GetRandomQuote = &cobra.Command{
	Use:   "random",
	Short: "Get a random tech quote",
	Long:  `Using random returns an English tech quote and the person who said it`,
	Run: func(cmd *cobra.Command, args []string) {
		client := &http.Client{}

		req, err := http.NewRequest("GET", baseURL+"/quotes/random", nil)
		if err != nil {
			fmt.Println(err)
		}

		req.Header.Add("content-type", "application/json")
		res, err := client.Do(req)
		if err != nil {
			fmt.Println(err)
		}
		defer res.Body.Close()

		bytes, _ := ioutil.ReadAll(res.Body)

		q := Quote{}

		err = json.Unmarshal(bytes, &q)
		if err != nil {
			fmt.Println(err)
		}
		fmt.Printf("%v\n- %v\n", q.Text, q.Author)
	},
}
