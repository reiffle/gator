package main

import (
	"context"
	"fmt"
	"os"
	"strconv"

	"github.com/reiffle/gator/internal/database"
)

func handlerBrowse(s *state, cmd command, user database.User) error {
	post_returns := int32(2)
	if len(cmd.args) != 0 {
		n, err := strconv.Atoi(cmd.args[0])
		if err != nil {
			fmt.Println("Argument should be an int")
			os.Exit(1)
			return nil
		}
		post_returns = int32(n)
	}

	get_posts := database.GetPostsForUserParams{
		UserID: user.ID,
		Limit:  post_returns,
	}
	posts, err := s.db.GetPostsForUser(context.Background(), get_posts)
	if err != nil {
		fmt.Println("Couldn't fetch posts")
		os.Exit(1)
		return err
	}

	for _, post := range posts {
		fmt.Println(post.Title)
		fmt.Println(post.Url)
		fmt.Println(post.PublishedAt.Time.Format("Mon Jan 2"))
		fmt.Printf("\n")
		fmt.Printf("%s", post.Description.String)
	}
	return nil
}
