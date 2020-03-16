package database

import (
	"database/sql"

	"vil/database/blog"
	"vil/database/roles"
	"vil/database/users"
)

// Database defines the database
type Database struct {
	Blog  *blog.Database
	Roles *roles.Database
	Users *users.Database
}

// New returns a new database.
func New(db *sql.DB) *Database {
	return &Database{
		Blog:  blog.New(db),
		Roles: roles.New(db),
		Users: users.New(db),
	}
}
