package main

import (
	"context"
	"fmt"
	"os"
)

func handlerFeeds(s *state, cmd command) error {
	feeds, err := s.db.PrintFeeds(context.Background())
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		os.Exit(1)
		return nil
	}

	for _, feed := range feeds {
		fmt.Println("=================")
		fmt.Println("Feed Name: ", feed.Name)
		fmt.Println("Feed URL: ", feed.Url)
		fmt.Println("User Name: ", feed.Name_2)
	}
	return nil
}
