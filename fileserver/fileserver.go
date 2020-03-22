package fileserver

import (
	"net/http"

	"github.com/jschweihs/httprouter"
)

// New creates a new fileserver application. The fileserver is
// responsible for serving the index.html file as well as
// public assets such as css, js, etc
func New(router *httprouter.Router) {

	// Serve index for existing pages
	router.GET("/", HandleIndex())
	router.GET("/team", HandleIndex())
	router.GET("/partners", HandleIndex())
	router.GET("/events", HandleIndex())
	router.GET("/shop", HandleIndex())
	router.GET("/contact", HandleIndex())
	// Admin pages still serve index
	router.GET("/login", HandleIndex())
	router.GET("/admin", HandleIndex())
	router.GET("/admin/users", HandleIndex())
	router.GET("/admin/users/:id", HandleIndex())
	router.GET("/admin/roles", HandleIndex())
	router.GET("/admin/blog", HandleIndex())
	router.GET("/admin/blog/:id", HandleIndex())

	router.GET("/favicon.ico", HandleFavIcon())
	router.GET("/admin/favicon.ico", HandleFavIcon())

	// Serve public assets
	router.ServeFiles("/css/*filepath", http.Dir("./../../public/css"))
	router.ServeFiles("/js/*filepath", http.Dir("./../../public/js"))
	router.ServeFiles("/images/*filepath", http.Dir("./../../public/images"))
	router.ServeFiles("/fonts/*filepath", http.Dir("./../../public/fonts"))
}

// HandleIndex is responsible for returning the index.html file
func HandleIndex() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "./../../public/index.html")
	}
}

// HandleFavIcon is responsible for returning the favicon.ico file
func HandleFavIcon() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "image/x-icon")
		http.ServeFile(w, r, "./../../public/favicon.ico")
	}
}
