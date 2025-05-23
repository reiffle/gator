package main

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/reiffle/gator/internal/database"
)

func handlerRegister(s *state, cmd command) error {
	if len(cmd.args) == 0 {
		return errors.New("the login handler expects a single argument, the username")
	}
	//create user
	ctxt := context.Background()
	name := cmd.args[0]
	params := database.CreateUserParams{
		ID:        uuid.New(),
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		Name:      name,
	}

	user, err := s.db.CreateUser(ctxt, params)
	if err != nil {
		return err
	}

	err = s.cfg.SetUser(user.Name)
	if err != nil {
		return err
	}

	fmt.Println("user has been created")
	printUser(user)
	return nil
}

func handlerLogin(s *state, cmd command) error {
	if len(cmd.args) == 0 {
		return errors.New("the login handler expects a single argument, the username")
	}
	//change the user
	name := cmd.args[0]
	_, err := s.db.GetUser(context.Background(), name)
	if err != nil {
		return err
	}
	err = s.cfg.SetUser(name)
	if err != nil {
		fmt.Printf("couldn't set user")
		return err
	}
	fmt.Printf("user '%s' has been set\n", name)
	return nil
}

func handlerReset(s *state, cmd command) error {

	//reset the database
	err := s.db.ResetUsers(context.Background())
	if err != nil {
		fmt.Println("couldn't reset table")
		return err
	}
	fmt.Println("table has been reset")
	return nil
}

func printUser(user database.User) {
	fmt.Printf("User ID:	%v\n", user.ID)
	fmt.Printf("User Name:	%v\n", user.Name)
}
