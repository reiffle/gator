package main

import (
	"context"
	"fmt"
	"os"
	"time"

	"github.com/google/uuid"
	"github.com/reiffle/gator/internal/database"
)

func handlerAddFeed(s *state, cmd command) error {
	if len(cmd.args) != 2 {
		fmt.Println("must include feed name and url")
		os.Exit(1)
	}
	curr_name := s.cfg.Current_user_name
	if len(curr_name) == 0 {
		fmt.Println("no current user")
		os.Exit(1)
	}
	curr_user, err := s.db.GetUser(context.Background(), curr_name)
	if err != nil {
		fmt.Println("current user not in database")
		os.Exit(1)
		return err
	}

	curr_user_id := curr_user.ID
	name := cmd.args[0]
	url := cmd.args[1]
	feed_params := database.CreateFeedParams{
		ID:        uuid.New(),
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		Name:      name,
		Url:       url,
		UserID:    curr_user_id,
	}

	feed, err := s.db.CreateFeed(context.Background(), feed_params)
	if err != nil {
		fmt.Println("Error creating feed", err)
		os.Exit(1)
		return err
	}

	printFeedInfo(feed)

	follow_feed_params := database.CreateFeedFollowParams{
		ID:        uuid.New(),
		CreatedAt: feed.CreatedAt,
		UpdatedAt: feed.UpdatedAt,
		UserID:    feed.UserID,
		FeedID:    feed.ID,
	}

	follow_feed, err := s.db.CreateFeedFollow(context.Background(), follow_feed_params)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
		return err
	}

	fmt.Println("========================")
	fmt.Println("Feed successfully followed")
	fmt.Println("Feed Name: ", follow_feed.FeedName)
	fmt.Println("User Name: ", follow_feed.UserName)

	return nil
}
