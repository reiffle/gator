package main

import (
	"context"
	"fmt"
	"os"
	"time"

	"github.com/reiffle/gator/internal/database"
)

func scrapeFeeds(s *state, cmd command) error {
	if len(cmd.args) != 1 {
		fmt.Printf("command needs a unit of time")
		os.Exit(1)
		return nil
	}
	name := s.cfg.Current_user_name
	user, err := s.db.GetUser(context.Background(), name)
	if err != nil {
		fmt.Println("couldn't get current user")
		return err
	}
	user_id := user.ID
	time_between_requests, err := time.ParseDuration(cmd.args[0])
	if err != nil {
		fmt.Println("Incorrect time given")
		fmt.Println(err)
		os.Exit(1)
		return nil
	}

	ticker := time.NewTicker(time_between_requests)
	for ; ; <-ticker.C {
		fmt.Println("Collecting Feeds every", time_between_requests)
		next_feed, err := s.db.GetNextFeedToFetch(context.Background(), user_id)
		if err != nil {
			fmt.Println("Couldn't get feed")
			fmt.Println(err)
			os.Exit(1)
			return nil
		}
		update_feed := database.MarkFeedsFetchedParams{
			UpdatedAt: time.Now().UTC(),
			ID:        next_feed.ID,
		}

		err = s.db.MarkFeedsFetched(context.Background(), update_feed)
		if err != nil {
			fmt.Println("Couldn't update feed")
			fmt.Println(err)
			os.Exit(1)
			return nil
		}

		err = handlerFetchFeed(s, cmd, next_feed.Url)
		if err != nil {
			fmt.Printf("Couldn't fetch feed from %s\n", next_feed.Url)
			os.Exit(1)
			return nil
		}
	}
}
