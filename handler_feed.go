package main

import (
	"calebwestbrook01/aggre_gator/internal/database"
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/google/uuid"
)

func handlerAddFeed(s *State, cmd Command) error {
	if len(cmd.Args) != 2 {
		return errors.New("Requires two arguments")
	}
	name := s.Config.Username
	user, err := s.db.GetUser(context.Background(), name)
	if err != nil {
		return err
	}
	params := database.CreateFeedParams{ID: uuid.New(), CreatedAt: time.Now(), UpdatedAt: time.Now(), Name: cmd.Args[0], Url: cmd.Args[1], UserID: user.ID}
	feed, err := s.db.CreateFeed(context.Background(), params)
	if err != nil {
		return err
	}
	fmt.Printf("%+v\n", feed)
	return nil
}

func handlerFeeds(s *State, cmd Command) error {
	if len(cmd.Args) > 0 {
		return errors.New("No arguments allowed")
	}
	rows, err := s.db.GetFeeds(context.Background())
	if err != nil {
		return err
	}
	for _, item := range rows {
		fmt.Printf("-%s:%s:%s\n", item.Name, item.Url, item.Username)
	}
	return nil
}
