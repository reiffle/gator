package main

import (
	"context"
	"fmt"
)

func handlerFetchFeed(s *state, cmd command) error {

	//Get aggregation
	URL := "https://www.wagslane.dev/index.xml"
	feed, err := fetchFeed(context.Background(), URL)
	if err != nil {
		fmt.Println("couldn't fetch aggregations")
		return nil
	}
	fmt.Printf("%+v\n", feed) //%+v\n prints structs nicely
	return nil
}
