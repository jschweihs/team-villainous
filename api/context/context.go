package context

import (
	"log"
	"vil/api/config"
	"vil/services"
)

type Context struct {
	Config   *config.Config
	Logger   *log.Logger
	Services *services.Services
}

// New returns a new API context
func New(config *config.Config, logger *log.Logger, services *services.Services) *Context {
	return &Context{
		Config:   config,
		Logger:   logger,
		Services: services,
	}
}
