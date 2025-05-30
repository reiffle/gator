package main

import (
	"context"
	"fmt"

	"github.com/reiffle/gator/internal/database"
)

func handlerPrintUsers(s *state, cmd command) error {

	//print users in database
	name := s.cfg.Current_user_name
	users, err := s.db.GetUsers(context.Background(), name)
	if err != nil {
		fmt.Println("couldn't print users")
		return err
	}
	for _, user := range users {
		fmt.Println(user)
	}
	return nil
}

func printUser(user database.User) {
	fmt.Printf("User ID:	%v\n", user.ID)
	fmt.Printf("User Name:	%v\n", user.Name)
}

func printFeedInfo(feed database.Feed) {
	fmt.Printf("ID:		%s\n", feed.ID)
	fmt.Printf("Created At:	%v\n", feed.CreatedAt)
	fmt.Printf("Updated At:	%v\n", feed.UpdatedAt)
	fmt.Printf("Name:		%s\n", feed.Name)
	fmt.Printf("URL:		%s\n", feed.Url)
	fmt.Printf("UserID:		%s\n", feed.UserID)
}
