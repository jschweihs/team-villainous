package api

import (
	"log"

	"vil/api/config"
	apictx "vil/api/context"
	v1 "vil/api/v1"
	"vil/database"
	"vil/services"

	"github.com/beeker1121/httprouter"
)

// New creates a new API application. All of the necessary routes for the
// API will be created on the given router, which should then be used to
// create the web server
func New(config *config.Config, logger *log.Logger, vdb *database.Database, router *httprouter.Router) {
	// Create the services
	services := services.New(vdb)

	// Create a new API context
	ac := apictx.New(config, logger, services)

	// Create a new API v1
	v1.New(ac, router)

	// Handle not found
	// router.NotFound = http.HandlerFunc(handleNotFound(ac))
}

// func handleNotFound(ac *apictx.Context) http.HandlerFunc {
// 	return func(w http.ResponseWriter, r *http.Request) {
// 		// Send 404 response
// 		errors.Default(ac.Logger, w, errors.ErrNotFound)
// 	}
// }
