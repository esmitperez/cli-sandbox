package cmd

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"net/http"
	"sort"

	"github.com/spf13/cobra"
)

var BaseUrl string = "https://esmit.me"

type websiteFunc func(cmd *cobra.Command, args []string)

var websiteCmd = &cobra.Command{
	Use:   "website",
	Short: "Displays Esmit's personal site link",
	Run:   websiteFuncF(),
}

type Item struct {
	Title string `xml:"title"`
	Link  string `xml:"link"`
}

type Channel struct {
	Title string `xml:"title"`
	Link  string `xml:"link"`
	Item  []Item `xml:"item"`
}

type RSS struct {
	XMLName xml.Name `xml:"rss"`
	// Xml     string `xml:",innerxml"`
	Channel Channel `xml:"channel"`
}

func websiteFuncF() websiteFunc {
	return func(cmd *cobra.Command, args []string) {
		fmt.Printf("\nRetrieving %s\n\n", BaseUrl)

		client := &http.Client{}

		res, err := client.Get(fmt.Sprintf("%s/tags/index.xml", BaseUrl))

		if err != nil {
			fmt.Printf("Error connecting to %s\n", BaseUrl)
		}
		defer res.Body.Close()

		bodyBytes, err := ioutil.ReadAll(res.Body)
		if err != nil {
			fmt.Println("Error reading body")

		}

		var document RSS
		if err = xml.Unmarshal(bodyBytes, &document); err != nil {
			fmt.Println("Error unmarshalling")
		}

		sortedTags := make([]string, len(document.Channel.Item))

		for i, tag := range document.Channel.Item {
			sortedTags[i] = tag.Title
		}
		sort.Strings(sortedTags)

		for _, tag := range sortedTags {
			fmt.Printf("%s\n", tag)
		}
	}
}

func init() {
	// make our program aware about this command
	rootCmd.AddCommand(websiteCmd)
}
