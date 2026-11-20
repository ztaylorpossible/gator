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

func handlerAddFeed(s *State, cmd Command, user database.User) error {
	if len(cmd.args) < 2 {
		return errors.New("addfeed expects a feed name and url");
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

	followParams := database.CreateFeedFollowParams{
		ID: uuid.New(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		UserID: user.ID,
		FeedID: feed.ID,
	};

	_, err = s.db.CreateFeedFollow(context.Background(), followParams);
	if err != nil {
		return err;
	}

	fmt.Printf("New feed created: %v\n", feed);
	return nil;
}

func handlerFeeds(s *State, cmd Command) error {
	feeds, err := s.db.GetFeeds(context.Background());
	if err != nil {
		return err;
	}

	for _, feed := range feeds {
		user, err := s.db.GetUserById(context.Background(), feed.UserID);
		if err != nil {
			return nil;
		}

		fmt.Printf("Feed \"%s\": %s | Created by %s\n", feed.Name, feed.Url, user.Name);
	}

	return nil;
}
