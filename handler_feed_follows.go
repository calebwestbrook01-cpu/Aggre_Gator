package main

import (
	"calebwestbrook01/aggre_gator/internal/database"
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/google/uuid"
)

func handlerFollow(s *State, cmd Command, user database.User) error {
	if len(cmd.Args) != 1 {
		return errors.New("Incorrect arguments")
	}
	feed, err := s.db.GetFeedByURL(context.Background(), cmd.Args[0])
	if err != nil {
		return err
	}
	feed_follow := database.CreateFeedFollowParams{ID: uuid.New(), CreatedAt: time.Now(), UpdatedAt: time.Now(), UserID: user.ID, FeedID: feed.ID}
	result, err := s.db.CreateFeedFollow(context.Background(), feed_follow)
	if err != nil {
		return err
	}
	fmt.Printf("%s\n%s", result.FeedName, result.UserName)
	return nil
}

func handlerFollowing(s *State, cmd Command, user database.User) error {
	if len(cmd.Args) > 0 {
		return errors.New("No arguments allowed")
	}
	rows, err := s.db.GetFeedFollowsForUser(context.Background(), user.ID)
	if err != nil {
		return err
	}
	for _, item := range rows {
		fmt.Printf("%s\n", item.FeedName)
	}
	return nil
}

func handlerUnfollow(s *State, cmd Command, user database.User) error {
	if len(cmd.Args) != 1 {
		return errors.New("Provide a single url")
	}
	feed, err := s.db.GetFeedByURL(context.Background(), cmd.Args[0])
	if err != nil {
		return err
	}
	feed_follow := database.DeleteFeedFollowParams{Url: feed.Url, Name: user.Name}
	err = s.db.DeleteFeedFollow(context.Background(), feed_follow)
	if err != nil {
		return err
	}
	return nil
}
