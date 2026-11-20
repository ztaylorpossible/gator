package main

import (
	"fmt"
	"os"
	"database/sql"
	"github.com/ztaylorpossible/gator/internal/config"
	"github.com/ztaylorpossible/gator/internal/database"
)

import _ "github.com/lib/pq"

func main() {
	conf, err := config.Read();

	if err != nil {
		fmt.Printf("Error reading config: %v\n", err);
		return;
	}

	state := State{conf: &conf};

	db, err := sql.Open("postgres", state.conf.Db_url);
	state.db = database.New(db);

	commands := Commands{commands: make(map[string]func(*State, Command) error)};

	commands.register("login", handlerLogin);
	commands.register("register", handlerRegister);
	commands.register("reset", handlerReset);
	commands.register("users", handlerUsers);
	commands.register("agg", handlerAgg);
	commands.register("addfeed", handlerAddFeed);
	commands.register("feeds", handlerFeeds);

	if len(os.Args) < 2 {
		fmt.Println("No command provided.");
		os.Exit(1);
	}

	command := Command{name: os.Args[1], args: os.Args[2:]}
	err = commands.run(&state, command);

	if err != nil {
		fmt.Printf("Error running command %s: %v\n", command.name, err);
		os.Exit(1);
	}
}

type State struct {
	conf *config.Config
	db *database.Queries
}
