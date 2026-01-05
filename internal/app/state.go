package app

import (
	"github.com/lucasgjanot/go-gator-feed/internal/config"
	"github.com/lucasgjanot/go-gator-feed/internal/database"
)

type State struct {
	Config *config.Config
	DB *database.Queries
}