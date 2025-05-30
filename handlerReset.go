package main

import (
	"context"
	"fmt"
)

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
