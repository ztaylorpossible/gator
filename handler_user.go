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

func handlerUsers(s *State, cmd Command) error {
	users, err := s.db.GetUsers(context.Background());
	if err != nil {
		return err;
	}

	for _, user := range users {
		var name string;
		if user.Name == s.conf.Current_user_name {
			name = fmt.Sprintf("* %s (current)\n", user.Name);
		} else {
			name = fmt.Sprintf("* %s\n", user.Name);
		}
		fmt.Printf("%s", name);
	}
	
	return nil;
}

func handlerFollow(s *State, cmd Command) error {
	if len(cmd.args) <= 0 {
		return errors.New("follow expects a url");
	}

	user, err := s.db.GetUser(context.Background(), s.conf.Current_user_name);
	if err != nil {
		return err;
	}

	feed, err := s.db.GetFeedByUrl(context.Background(), cmd.args[0]);
	if err != nil {
		return err;
	}

	params := database.CreateFeedFollowParams{
		ID: uuid.New(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		UserID: user.ID,
		FeedID: feed.ID,
	};

	follow, err := s.db.CreateFeedFollow(context.Background(), params);
	if err != nil {
		return err;
	}

	fmt.Printf("Feed \"%s\" followed by %s.\n", follow.FeedName, follow.UserName);

	return nil;
}

func handlerFollowing(s *State, cmd Command) error {
	user, err := s.db.GetUser(context.Background(), s.conf.Current_user_name);
	if err != nil {
		return err;
	}

	feeds, err := s.db.GetFeedFollowsForUser(context.Background(), user.ID);
	if err != nil {
		return err;
	}

	fmt.Printf("%s is following:\n", user.Name);
	for _, feed := range feeds {
		fmt.Printf("\"%s\"\n", feed.FeedName);
	}

	return nil;
}
