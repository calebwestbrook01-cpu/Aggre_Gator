package main

import (
	"calebwestbrook01/aggre_gator/internal/config"
	"calebwestbrook01/aggre_gator/internal/database"
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/lib/pq"
)

func main() {
	cfg, err := config.Read()
	if err != nil {
		log.Fatal(err)
	}
	dbURL := cfg.URL
	db, err := sql.Open("postgres", dbURL)
	if err != nil {
		log.Fatal(err)
	}
	dbQueries := database.New(db)
	state := State{Config: &cfg, db: dbQueries}
	cmds := Commands{Handlers: map[string]func(*State, Command) error{}}
	cmds.register("login", handlerLogin)
	cmds.register("register", handlerRegister)
	cmds.register("reset", handlerReset)
	cmds.register("users", handlerUsers)
	cmds.register("agg", handlerAgg)
	cmds.register("addfeed", handlerAddFeed)
	cmds.register("feeds", handlerFeeds)
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
	db     *database.Queries
}
