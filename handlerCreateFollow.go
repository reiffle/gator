package main

import (
	"context"
	"fmt"
	"os"
	"time"

	"github.com/google/uuid"
	"github.com/reiffle/gator/internal/database"
)

func handlerFeedFollow(s *state, cmd command, user database.User) error {

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
	feed_params := database.CreateFeedFollowParams{
		ID:        uuid.New(),
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		UserID:    curr_user_id,
		FeedID:    feed_id,
	}

	feed, err := s.db.CreateFeedFollow(context.Background(), feed_params)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
		return err
	}

	fmt.Println("Feed successfully followed")
	fmt.Println("Feed Name: ", feed.FeedName)
	fmt.Println("User Name: ", feed.UserName)
	return nil
}
