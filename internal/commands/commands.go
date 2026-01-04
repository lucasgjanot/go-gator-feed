package commands

import (
	"errors"

	"github.com/lucasgjanot/go-gator-feed/internal/app"
)

type Command struct {
	Name string
	Args []string
}

type Commands struct {
	RegisteredCommands map[string]func(*app.State, Command) error
}

func (c *Commands) Register(name string, f func(*app.State, Command) error) {
	c.RegisteredCommands[name] = f
}

func (c *Commands) Run(s *app.State, cmd Command) error {
	f, ok := c.RegisteredCommands[cmd.Name]
	if !ok {
		return errors.New("command not found")
	}
	
	return f(s, cmd)
}
