package main

import (
	"errors"
	"fmt"
)

// this will be used to add and build functions
type command struct {
	command_name string
	args         []string
}

// holds a map (dictionary) of all user commands to run functions
type commands struct {
	//key is [command name] value is a command func
	cmds map[string]func(*state, command) error
}

// create a new command
func (c *commands) register(name string, f func(*state, command) error) {
	if name == "" {
		fmt.Println("no function name given")
		return
	}
	_, exists := c.cmds[name]
	if exists {
		fmt.Println("function name already exists")
		return
	}
	c.cmds[name] = f
}

// run an existing command or return error if it doesn't exist
func (c *commands) run(s *state, cmd command) error {
	new_func, ok := c.cmds[cmd.command_name]
	if !ok {
		return errors.New("command does not exist")
	}
	//return the error code after running the function
	return new_func(s, cmd)
}

// End commands struct and methods
