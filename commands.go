package main

import (
	"fmt"
)

type Command struct {
	name string
	args []string
}

type Commands struct {
	commands map[string]func(*State, Command) error
}

func (c *Commands) run(s *State, cmd Command) error {
	action, ok := c.commands[cmd.name];
	if !ok {
		return fmt.Errorf("Command %s isn't registered.", cmd.name);
	}

	err := action(s, cmd);
	if err != nil {
		return err;
	}

	return nil;
}

func (c *Commands) register(name string, f func(*State, Command) error) {
	c.commands[name] = f;
}
