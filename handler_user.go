package main

import (
	"fmt"
	"errors"
)

func handlerLogin(s *State, cmd Command) error {
	if len(cmd.args) <= 0 {
		return errors.New("login expects a username");
	}

	err := s.conf.SetUser(cmd.args[0]);
	if err != nil {
		return err;
	}

	fmt.Printf("User has been set to %s\n", cmd.args[0]);

	return nil;
}
