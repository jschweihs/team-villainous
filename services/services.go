package services

import (
	"vil/database"
	"vil/services/blog"
	"vil/services/roles"
	"vil/services/users"
)

// Services defines the services
type Services struct {
	Blog  *blog.Service
	Roles *roles.Service
	Users *users.Service
}

// New returns a new set of services
func New(db *database.Database) *Services {
	return &Services{
		Blog:  blog.New(db),
		Roles: roles.New(db),
		Users: users.New(db),
	}
}
