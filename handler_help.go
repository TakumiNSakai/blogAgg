package main

import (
	"fmt"
)

func handlerHelp(s *state, cmd command) error {
	fmt.Println("For list of commands, refer to https://github.com/TakuumiS/blogAgg")
	return nil
}
