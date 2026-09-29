package main

import (
	"calebwestbrook01/aggre_gator/internal/database"
	"context"
)

func middlewareLoggedIn(handler func(s *State, cmd Command, user database.User) error) func(*State, Command) error {
	return func(s *State, cmd Command) error {
		user, err := s.db.GetUser(context.Background(), s.Config.Username)
		if err != nil {
			return err
		}
		return handler(s, cmd, user)
	}
}
