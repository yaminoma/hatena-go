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
	//res, err := hatena_go.UserStar("eightpockets")
	//fmt.Println(res.StarCount)

	//
	// Count hatena comment star
	//
	res, err := hatena_go.CommentStar("http://b.hatena.ne.jp/jkondo/20150210#bookmark-241358174")

	if err != nil {
		log.Fatal(err)
	}

	if res.Entries != nil {
		fmt.Println("Star:")
		for _, star := range res.Entries[0].Stars {
			fmt.Println(star.Quote, star.Name)
		}
	}
}
