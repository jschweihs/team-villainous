package file

import (
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"net/url"
	"os"
	"strings"

	apictx "vil/api/context"
	"vil/api/errors"
	"vil/api/middleware/auth"
	"vil/api/render"

	"github.com/beeker1121/httprouter"
)

// Note: Right now this package is only designed for handling images

// TODO: Implement type for upload and create it's own service
type UploadParams struct {
	Name   *string
	Folder *string
	File   *multipart.File
}

// ResultUpload defines the response data for HandlePost
type ResultUpload struct {
	Data bool `json:"data"`
}

// New creates the routes for the blog entries endpoints of the API
func New(ac *apictx.Context, router *httprouter.Router) {
	// Handle the routes
	router.POST("/api/v1/upload", auth.AuthenticateEndpoint(ac, HandlePost(ac)))
}

// HandlePost handles the /api/v1/upload POST route of the API
func HandlePost(ac *apictx.Context) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Get this user from the request context
		// This validates that this user exists and has their token
		_, err := auth.GetUserFromRequest(r)
		if err != nil {
			errors.Default(ac.Logger, w, errors.ErrInternalServerError)
			return
		}

		// Create a new API Errors
		errs := &errors.Errors{}

		// Parse multiform
		if err := r.ParseMultipartForm(32 << 20); err != nil {
			errors.Default(ac.Logger, w, errors.ErrBadRequest)
			return
		}

		// Get new image name
		name := strings.Replace(url.QueryEscape(r.FormValue("name")), "+", "%20", -1)
		if name == "" {
			errs.Add(errors.New(http.StatusBadRequest, "name", ErrNameInvalid.Error()))
		}

		// Get folder name
		folder := r.FormValue("folder")
		if folder == "" {
			errs.Add(errors.New(http.StatusBadRequest, "folder", ErrFolderInvalid.Error()))
		}

		// Get the file
		file, handler, err := r.FormFile("image")
		if err != nil {
			errors.Default(ac.Logger, w, errors.ErrInternalServerError)
			return
		}
		defer file.Close()

		// Get file extension
		ext := "." + strings.Split(handler.Filename, ".")[1]

		// Create file at source location
		// os.Create takes a string and returns *File, error
		// The File type DOES implement the Read/Write interface
		appPath := "./../../app/public/images/" + folder + "/" + name + ext
		appFile, err := os.Create(appPath)
		if err != nil {
			errors.Default(ac.Logger, w, errors.ErrInternalServerError)
			return
		}
		defer appFile.Close()

		// Move uploaded file to appFile
		// io.Copy takes a Writer and Reader and returns int64, error
		size, err := io.Copy(appFile, file)
		if size == 0 {
			errors.Default(ac.Logger, w, errors.ErrInternalServerError)
			return
		} else if err != nil {
			errors.Default(ac.Logger, w, errors.ErrInternalServerError)
			return
		}

		// Open file placed in app
		in, err := os.Open(appPath)
		if err != nil {
			return
		}
		defer in.Close()

		// Create file at public location
		publicPath := "./../../public/images/" + folder + "/" + name + ext
		publicFile, err := os.Create(publicPath)
		if err != nil {
			errors.Default(ac.Logger, w, errors.ErrInternalServerError)
			return
		}
		defer publicFile.Close()

		// Move uploaded file to publicFile
		size, err = io.Copy(publicFile, in)
		if err != nil {
			errors.Default(ac.Logger, w, errors.ErrInternalServerError)
			return
		} else if size == 0 {
			errors.Default(ac.Logger, w, errors.ErrInternalServerError)
			return
		}

		// Create a new result
		result := ResultUpload{
			Data: true,
		}

		// Render output
		if err := render.JSON(w, true, result); err != nil {
			fmt.Println(err)
			ac.Logger.Printf("render.JSON() error: %s\n", err)
			errors.Default(ac.Logger, w, errors.ErrInternalServerError)
			return
		}

	}
}
