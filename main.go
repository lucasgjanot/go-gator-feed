package main

import (
	"database/sql"
	"errors"
	"fmt"
	"log"
	"os"

	_ "github.com/lib/pq"
	"github.com/lucasgjanot/go-gator-feed/internal/cli"
	"github.com/lucasgjanot/go-gator-feed/internal/commands"
	"github.com/lucasgjanot/go-gator-feed/internal/config"
	"github.com/lucasgjanot/go-gator-feed/internal/database"
	"github.com/lucasgjanot/go-gator-feed/internal/runtime"
)

func main() {
	cfg, err := config.Read()
	if err != nil {
		log.Fatalf("error reading config: %v", err)
	}

	db, err := sql.Open("postgres", cfg.DBURL)
	if err != nil {
		log.Fatalf("database connection error: %v", err)
	}
	defer db.Close()


	queries := database.New(db)

	state := &runtime.State{
		Config:   &cfg,
		Database: runtime.Database{
			User: queries,
			Feed: queries,
		},
		Output:   cli.CLIOutput{},
	}

	cmds := commands.Commands{
		RegisteredCommands: make(map[string]func(*runtime.State, commands.Command) error),
	}

	cmds.Register("login", commands.CommandLogin)
	cmds.Register("register", commands.CommandRegister)
	cmds.Register("reset", commands.CommandReset)
	cmds.Register("users", commands.CommandUsers)
	cmds.Register("agg", commands.CommandAgg)
	cmds.Register("addfeed", commands.CommandAddFeed)

	if len(os.Args) < 2 {
		log.Fatal("Usage: cli <command> [args...]")
	}

	cmd := commands.Command{
		Name: os.Args[1],
		Args: os.Args[2:],
	}

	if err := cmds.Run(state, cmd); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

}

func handleError(err error) {
	switch {
	case errors.Is(err, runtime.ErrUserNotFound):
		fmt.Println("User not found")
	case errors.Is(err, runtime.ErrUserExists):
		fmt.Println("User already exists")
	case errors.Is(err, runtime.ErrNoUsers):
		fmt.Println("no users registered")
	case errors.Is(err, runtime.ErrFeedNotFound):
		fmt.Println("Feed not found")
	case errors.Is(err, runtime.ErrFeedExists):
		fmt.Println("Feed already exists")
	case errors.Is(err, runtime.ErrNoFeed):
		fmt.Println("no Feeds registered")
	default:
		// erro técnico, útil para debug
		fmt.Println("Erro:", err)
	}
}
