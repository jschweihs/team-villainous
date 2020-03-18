package v1

import (
	apictx "vil/api/context"
	"vil/api/v1/handlers/blog"
	"vil/api/v1/handlers/file"
	"vil/api/v1/handlers/login"
	"vil/api/v1/handlers/public"
	"vil/api/v1/handlers/roles"
	"vil/api/v1/handlers/users"

	"github.com/beeker1121/httprouter"
)

// New creates a new API v1 application. All of the necessary routes for
// v1 of the API will be created on the given router, which should then be
// used to create the web server. The root domain should be api.maildb.io
// or something similar.

func New(ac *apictx.Context, router *httprouter.Router) {
	public.New(ac, router)
	// Create all of the API v1 routes
	blog.New(ac, router)
	file.New(ac, router)
	login.New(ac, router)
	roles.New(ac, router)
	users.New(ac, router)
}
