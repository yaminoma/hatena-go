package main

import (
	"fmt"
	"log"

	"github.com/8pockets/hatena-go"
)

func main() {
	res, err := hatena_go.EntryInfo("https://github.com/")

	if err != nil {
		log.Fatal(err)
	}

	if res.Bookmarks != nil {
		fmt.Println("Related Bookmark Entry:")
		for _, entry := range res.RelatedEntries {
			fmt.Println(entry.Title, entry.Url)
		}
	}
}
