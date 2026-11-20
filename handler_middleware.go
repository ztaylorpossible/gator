package main

import (
	"context"
	"github.com/ztaylorpossible/gator/internal/database"
)

func middlewareLoggedIn(handler func(s *State, cmd Command, user database.User) error) func(*State, Command) error {
	return func(s *State, cmd Command) error {
		user, err := s.db.GetUser(context.Background(), s.conf.Current_user_name);
		if err != nil {
			return err;
		}

		err = handler(s, cmd, user);
		if err != nil {
			return err;
		}

		return nil;
	};
}
