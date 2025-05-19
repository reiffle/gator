package main

import (
	"fmt"
	"os"

	"github.com/reiffle/gator/internal/config" //Need full path for main to access internal
)

func main() {
	//get current config file
	cfg, err := config.Read()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	new_state := state{cfg: &cfg}
	commands := commands{cmds: make(map[string]func(*state, command) error)}
	commands.register("login", handlerLogin)
	if len(os.Args) < 2 {
		fmt.Println("need command name")
		os.Exit(1)
	}
	new_command := command{command_name: os.Args[1], args: os.Args[2:]}
	err = commands.run(&new_state, new_command)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	os.Exit(0)
}
