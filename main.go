package main

import (
	"calebwestbrook01/aggre_gator/internal/config"
	"fmt"
	"log"
	"os"
)

func main() {
	cfg, err := config.Read()
	if err != nil {
		log.Fatal(err)
	}
	state := State{Config: &cfg}
	cmds := Commands{Handlers: map[string]func(*State, Command) error{}}
	cmds.register("login", handlerLogin)
	if len(os.Args) < 2 {
		fmt.Println("Missing arguments")
		os.Exit(1)
	}
	cmd := Command{Name: os.Args[1], Args: os.Args[2:]}
	err = cmds.run(&state, cmd)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

}

type State struct {
	Config *config.Config
}
