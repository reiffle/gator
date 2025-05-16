package main

import (
	"fmt"

	"github.com/reiffle/gator/internal/config" //Need full path for main to access internal
)

func main() {
	cfg, err := config.Read()
	if err != nil {
		fmt.Print(err)
	}
	err = cfg.SetUser("Peder")
	if err != nil {
		fmt.Print(err)
	}
	cfg, err = config.Read()
	if err != nil {
		fmt.Print(err)
	}
	fmt.Printf("%+v\n", cfg) //%+v prints struct key and value, %v just prints struct values
}
