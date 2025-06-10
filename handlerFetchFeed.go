package main

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/reiffle/gator/internal/database"
)

func handlerFetchFeed(s *state, cmd command, URL string, ID uuid.UUID) error {

	//Get feed
	feed, err := fetchFeed(context.Background(), URL)
	if err != nil {
		fmt.Println("couldn't fetch feed")
		return nil
	}
	fmt.Println(feed.Channel.Title)
	fmt.Printf("\n")
	fmt.Println("---------------------------")

	post_Params := database.CreatePostParams{}

	for _, x := range feed.Channel.Item {

		post_Params.ID = uuid.New()
		post_Params.CreatedAt = time.Now().UTC()
		post_Params.UpdatedAt = time.Now().UTC()
		post_Params.Title = x.Title
		post_Params.Url = x.Link
		if x.Description != "" {
			post_Params.Description = sql.NullString{String: x.Description, Valid: true} //little complicated
		} else {
			post_Params.Description = sql.NullString{Valid: false}
		}
		publishedAt, err := parsePublishedAt(x.PubDate)
		if err == nil {
			post_Params.PublishedAt = sql.NullTime{Time: publishedAt, Valid: true}
		} else {
			post_Params.PublishedAt = sql.NullTime{Valid: false}
			fmt.Printf("Error parsing published date '%s': %v\n", x.PubDate, err)
		}
		post_Params.FeedID = ID
		_, err = s.db.CreatePost(context.Background(), post_Params)
		if err != nil {
			return err
		}
	}

	return nil
}

func parsePublishedAt(pubDate string) (time.Time, error) {
	layouts := []string{
		time.RFC1123Z,
		time.RFC1123,
		time.RFC822Z,
		time.RFC822,
		time.RFC3339,
	}
	for _, layout := range layouts {
		t, err := time.Parse(layout, pubDate)
		if err == nil {
			return t, nil
		}
	}
	return time.Time{}, fmt.Errorf("unable to parse published date '%s' with known layouts", pubDate)
}
