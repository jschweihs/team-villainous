package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"time"
	"vil/api"
	"vil/api/config"
	"vil/database"
	"vil/fileserver"

	"github.com/beeker1121/creek"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jschweihs/httprouter"
)

func main() {
	// Parse the API configuration file
	cfg, err := config.ParseConfigFile("config.json")
	if err != nil {
		panic(err)
	}

	// Get the configuration environment variables
	// cfg.DBHost = os.Getenv("DB_HOST")
	// cfg.DBPort = os.Getenv("DB_PORT")
	// cfg.DBName = os.Getenv("DB_NAME")
	// cfg.DBUser = os.Getenv("DB_USER")
	// cfg.DBPass = os.Getenv("DB_PASS")
	// cfg.APIHost = os.Getenv("API_HOST")
	// cfg.APIPort = os.Getenv("API_PORT")
	// cfg.JWTSecret = os.Getenv("JWT_SECRET")

	// Create new creek logger with 10 MB max file size.
	logger := log.New(creek.New(cfg.LogFile, 10), "Vil API: ", log.Llongfile|log.LstdFlags)
	logger.Printf("Starting Vil API server at %s/n", time.Now().UTC().Format(time.RFC3339))

	// Connect to the MYSQL database.
	conn := cfg.DBUser + ":" + cfg.DBPass + "@tcp(" + cfg.DBHost + ":" + cfg.DBPort + ")/" + cfg.DBName + "?parseTime=true"
	db, err := sql.Open("mysql", conn)
	if err != nil {
		logger.Fatal(err)
	}
	defer db.Close()

	// Test database connection
	if err := db.Ping(); err != nil {
		log.Fatal(err)
	}

	// Create a new Vil database
	vdb := database.New(db)

	// Create a new API
	router := httprouter.New()

	// Add servers
	fileserver.New(router)
	api.New(cfg, logger, vdb, router)

	// Create a new HTTP server
	server := &http.Server{
		Addr:           ":" + cfg.APIPort,
		Handler:        router,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	fmt.Printf("\nRunning server...\n")

	// Start the HTTP server
	if err := server.ListenAndServe(); err != nil {
		logger.Fatal(err)
	}

}
