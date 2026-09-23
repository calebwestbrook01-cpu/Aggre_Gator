package main

import (
	"calebwestbrook01/aggre_gator/internal/database"
	"context"
	"errors"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/google/uuid"
)

func handlerLogin(s *State, cmd Command) error {
	if len(cmd.Args) == 0 {
		return errors.New("Missing username")
	}
	_, err := s.db.GetUser(context.Background(), cmd.Args[0])
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	err = s.Config.SetUser(cmd.Args[0])
	if err != nil {
		return err
	}
	fmt.Println("User set!")
	return nil
}

func handlerRegister(s *State, cmd Command) error {
	if len(cmd.Args) == 0 {
		return errors.New("Missing name")
	}
	params := database.CreateUserParams{ID: uuid.New(), CreatedAt: time.Now(), UpdatedAt: time.Now(), Name: cmd.Args[0]}
	user, err := s.db.CreateUser(context.Background(), params)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	err = s.Config.SetUser(user.Name)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	log.Printf("%+v", user)
	fmt.Println("User set!")
	return nil
}

func handlerReset(s *State, cmd Command) error {
	err := s.db.DeleteUsers(context.Background())
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Println("Reset successful!")
	return nil
}

func handlerUsers(s *State, cmd Command) error {
	rows, err := s.db.GetUsers(context.Background())
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	for _, user := range rows {
		if user.Name == s.Config.Username {
			fmt.Printf("%s (current)", user.Name)
		} else {
			fmt.Println(user.Name)
		}
	}
	return nil
}
