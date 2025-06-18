package main

import "errors"

// command represents a command that can be executed by the application.
type command struct {
	Name string
	Args []string
}

// commands represents a collection of commands that can be executed by the application.
type commands struct {
	registeredCommands map[string]func(*state, command) error
}

// register adds a new command to the collection.
func (c *commands) register(name string, f func(*state, command) error) {
	c.registeredCommands[name] = f
}

// run executes the given command.
func (c *commands) run(s *state, cmd command) error {
	f, ok := c.registeredCommands[cmd.Name]
	if !ok {
		return errors.New("command not found")
	}
	return f(s, cmd)
}
