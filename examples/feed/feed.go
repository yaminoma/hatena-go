package main

import (
	"fmt"
	"log"

	"github.com/8pockets/hatena-go"
)

func main() {

	//
	// Get IT category feed
	//
	options := map[string]string{
		"sort":      "popular",
		"threshold": "10",
	}
	res, err := hatena_go.NewEntry("it", options)

	//
	// Search entry from matching url
	//

	//options := map[string]string{
	//	"sort": "count",
	//}
	//res, err := hatena_go.SearchUrlEntry("hatena.ne.jp", options)

	if err != nil {
		log.Fatal(err)
	}

	if res.Entry != nil {
		fmt.Println("Entries:")
		for _, entry := range res.Entry {
			fmt.Println(entry.Subject, entry.Title)
		}
	}
}
