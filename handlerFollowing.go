package main

import (
	"context"
	"fmt"
	"os"

	"github.com/reiffle/gator/internal/database"
)

func handlerFollowing(s *state, cmd command, user database.User) error {
	if len(cmd.args) != 0 {
		fmt.Println("This function doesn't take any parameters")
		os.Exit(1)
		return nil
	}

	feed_follows, err := s.db.GetFeedFollowsForUser(context.Background(), user.Name)
	if err != nil {
		fmt.Println("No follows for current user")
		os.Exit(1)
		return nil
	}
	fmt.Printf("%s is following\n\n", user.Name)
	if len(feed_follows) == 0 {
		fmt.Println("NONE")
		return nil
	}
	for _, record := range feed_follows {
		fmt.Println(record.FeedName)
	}
	return nil
}
