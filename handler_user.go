package main

import (
	"fmt"
	"errors"
	"time"
	"context"
	"github.com/google/uuid"
	"github.com/ztaylorpossible/gator/internal/database"
)

func handlerLogin(s *State, cmd Command) error {
	if len(cmd.args) <= 0 {
		return errors.New("login expects a username");
	}

	user, err := s.db.GetUser(context.Background(), cmd.args[0]);
	if err != nil {
		return err;
	}

	err = s.conf.SetUser(user.Name);
	if err != nil {
		return err;
	}

	fmt.Printf("User has been set to %s\n", cmd.args[0]);

	return nil;
}

 func handlerRegister(s *State, cmd Command) error {
	if len(cmd.args) <= 0 {
		return errors.New("register expects a username");
	}

	userParams := database.CreateUserParams{
		ID: uuid.New(),
		Name: cmd.args[0],
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	};

	user, err := s.db.CreateUser(context.Background(), userParams);
	if err != nil {
		return err;
	}

	err = s.conf.SetUser(user.Name);
	if err != nil {
		return err;
	}

	fmt.Printf("New user registered: %s at %v\n", user.Name, user.CreatedAt);

	return nil;
 }

 func handlerReset(s *State, cmd Command) error {
	 err := s.db.DeleteUsers(context.Background());
	 if err != nil {
		 return err;
	 }

	 fmt.Println("Users were successfully deleted.");

	 return nil;
 }
