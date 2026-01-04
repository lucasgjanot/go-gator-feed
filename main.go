package main

import (
	"log"
	"os"

	"github.com/lucasgjanot/go-gator-feed/internal/app"
	"github.com/lucasgjanot/go-gator-feed/internal/commands"
	"github.com/lucasgjanot/go-gator-feed/internal/config"
	"github.com/lucasgjanot/go-gator-feed/internal/handlers"
)





func main() {
	cfg, err := config.Read()
	if err != nil {
		log.Fatalf("error reading config: %v", err)
	}

	programState := &app.State{
		Config: &cfg,
	}

	cmds := commands.Commands{
		RegisteredCommands: make(map[string]func(*app.State, commands.Command) error),
	}

	cmds.Register("login", handlers.HandlerLogin)

	if len(os.Args) < 2 {
		log.Fatal("Usage: cli <command> [args...]")
	}

	cmdName := os.Args[1]
	cmdArgs := os.Args[2:]

	err = cmds.Run(programState, commands.Command{Name: cmdName, Args: cmdArgs})
	if err != nil {
		log.Fatal(err)
	}
}
