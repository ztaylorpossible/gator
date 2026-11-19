package main

import (
	"fmt"
	"errors"
	"time"
	"context"
	"github.com/google/uuid"
	"github.com/ztaylorpossible/gator/internal/database"
)

func handlerAgg(s *State, cmd Command) error {
	feed, err := fetchFeed(context.Background(), "https://www.wagslane.dev/index.xml")
	if err != nil {
		return err;
	}

	fmt.Printf("Feed:\n%v\n", feed);
	return nil;
}

func handlerAddFeed(s *State, cmd Command) error {
	if len(cmd.args) < 2 {
		return errors.New("addfeed expects a feed name and url");
	}

	user, err := s.db.GetUser(context.Background(), s.conf.Current_user_name);
	if err != nil {
		return err;
	}

	feedParams := database.CreateFeedParams{
		ID: uuid.New(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		Name: cmd.args[0],
		Url: cmd.args[1],
		UserID: user.ID,
	};

	feed, err := s.db.CreateFeed(context.Background(), feedParams);
	if err != nil {
		return err;
	}

	fmt.Printf("New feed created: %v\n", feed);
	return nil;
}
