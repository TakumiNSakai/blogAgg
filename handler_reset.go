package main

import (
	"context"
	"fmt"
)

// handlerReset deletes all users from the database.
func handlerReset(s *state, cmd command) error {
	err := s.db.DeleteUsers(context.Background())
	if err != nil {
		return fmt.Errorf("couldn't delete users: %w", err)
	}
	fmt.Println("Database reset successfully!")
	return nil
}
