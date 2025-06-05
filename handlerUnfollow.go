package main

import (
	"context"
	"fmt"
	"os"

	"github.com/reiffle/gator/internal/database"
)

func handlerUnfollow(s *state, cmd command, user database.User) error {

	if len(cmd.args) != 1 {
		fmt.Println("must include url")
		os.Exit(1)
	}

	url := cmd.args[0]
	curr_feed, err := s.db.FindFeed(context.Background(), url)
	if err != nil {
		fmt.Printf("%v+\n", curr_feed)
		fmt.Println("can't find feed for that URL")
		fmt.Println(err)
		os.Exit(1)
		return nil
	}

	curr_user_id := user.ID
	feed_id := curr_feed.ID
	unfollow_params := database.UnfollowFeedParams{
		UserID: curr_user_id,
		FeedID: feed_id,
	}

	err = s.db.UnfollowFeed(context.Background(), unfollow_params)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
		return err
	}

	fmt.Println("Feed successfully unfollowed")
	fmt.Println("Feed Name: ", url)
	fmt.Println("User Name: ", user.Name)
	return nil
}
