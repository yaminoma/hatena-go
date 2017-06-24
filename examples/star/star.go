package main

import (
	"fmt"
	"log"

	"github.com/8pockets/hatena-go"
)

func main() {

	//
	// Count User's star
	//
	res1, err := hatena.UserStar("jkondo")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(res1.StarCount)

	//
	// Count hatena comment star
	//
	res2, err := hatena.CommentStar("http://b.hatena.ne.jp/jkondo/20150210#bookmark-241358174")

	if err != nil {
		log.Fatal(err)
	}

	if res2.Entries != nil {
		for _, star := range res2.Entries[0].Stars {
			fmt.Println(star.Quote, star.Name)
		}
	}
}
