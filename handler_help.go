package main

import (
	"fmt"
)

// handlerReset deletes all users from the database.
func handlerHelp(s *state, cmd command) error {
	fmt.Println("For list of commands, refer to https://github.com/TakuumiS/blogAgg")
	return nil
}
