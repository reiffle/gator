package main

import (
	"context"
	"fmt"
)

func handlerFetchFeed(s *state, cmd command, URL string) error {

	//Get feed
	feed, err := fetchFeed(context.Background(), URL)
	if err != nil {
		fmt.Println("couldn't fetch feed")
		return nil
	}
	fmt.Println(feed.Channel.Title)
	fmt.Printf("\n")
	fmt.Println("---------------------------")
	for _, x := range feed.Channel.Item {
		fmt.Println("* ", x.Title) //Print rss titles
	}
	fmt.Printf("\n")
	return nil
}
