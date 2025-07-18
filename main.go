package main

import (
	"context"
	"fmt"
	"os"

	"github.com/reiffle/gator/internal/config"

	/*
		..................................
		begin imports we will always need for working with databases
		..................................
	*/
	"database/sql" // standard Go commands to use with databases

	_ "github.com/lib/pq"                        // specific initialization for DB we're using (postGres, mysql, etc), only needed indirectly
	"github.com/reiffle/gator/internal/database" //the actual database path I'm working with
	/*
		..................................
		end universal imports
		..................................
	*/)

type state struct {
	cfg *config.Config //cfg has Current_user_name and DbURL fields
	db  *database.Queries
}

func main() {
	//get current config file
	cfg, err := config.Read()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	//get database by passing in the type of database and it's location
	db, err := sql.Open("postgres", cfg.DBURL)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	defer db.Close()              //don't forget to close db
	dbQueries := database.New(db) //returns a pointer to a new instance of db that is typesafe and correctly uses the functions I defined by sqlc from by sql files

	new_state := &state{cfg: &cfg, db: dbQueries} //groups the updated cfg and db structs, plus whatever we add later

	commands := commands{cmds: make(map[string]func(*state, command) error)} //create a new commands instance
	//populate the new commands variable
	commands.register("login", handlerLogin)
	commands.register("register", handlerRegister)
	commands.register("reset", handlerReset)
	commands.register("users", handlerPrintUsers)
	commands.register("agg", scrapeFeeds)
	commands.register("addfeed", middlewareLoggedIn(handlerAddFeed))
	commands.register("feeds", handlerFeeds)
	commands.register("follow", middlewareLoggedIn(handlerFeedFollow))
	commands.register("following", middlewareLoggedIn(handlerFollowing))
	commands.register("unfollow", middlewareLoggedIn(handlerUnfollow))
	commands.register("browse", middlewareLoggedIn(handlerBrowse))

	if len(os.Args) < 2 { //check to make sure that the user put in a command name
		fmt.Println("need command name")
		os.Exit(1)
	}

	new_command := command{command_name: os.Args[1], args: os.Args[2:]}
	err = commands.run(new_state, new_command)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	os.Exit(0)
}

func middlewareLoggedIn(handler func(s *state, cmd command, user database.User) error) func(*state, command) error {
	return func(s *state, cmd command) error { //we are actually returning middlewareLoggedIn, not the inner function anymore
		curr_user, err := s.db.GetUser(context.Background(), s.cfg.Current_user_name)
		if err != nil {
			return err
		}

		return handler(s, cmd, curr_user) //we are running a command
	}
}
