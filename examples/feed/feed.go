package main

import (
	"fmt"
	"log"

	"github.com/yaminoma/hatena-go"
)

func main() {

	// Get IT-category feed
	newOpts := map[string]string{
		"sort":      "popular",
		"threshold": "10",
	}
	new, err := hatena.NewEntry("it", newOpts)
	if err != nil {
		log.Fatal(err)
	}

	if new.Entry != nil {
		for _, entry := range new.Entry {
			fmt.Println(entry.Subject, entry.Title)
		}
	}

	// Search entry from matching url
	searchOpts := map[string]string{
		"sort": "count",
	}
	search, err := hatena.SearchUrlEntry("hatena.ne.jp", searchOpts)
	if err != nil {
		log.Fatal(err)
	}

	if search.Entry != nil {
		for _, entry := range search.Entry {
			fmt.Println(entry.Subject, entry.Title)
		}
	}
}
