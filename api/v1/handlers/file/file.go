package file

import (
	"io"
	"mime/multipart"
	"net/http"
	"os"
	"strings"

	apictx "vil/api/context"
	"vil/api/errors"
	"vil/api/middleware/auth"
	"vil/api/render"

	"github.com/jschweihs/httprouter"
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
// If no file is provided, this will attempt to use the placeholder
// file in the given folder instead
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
		name := strings.ReplaceAll(r.FormValue("name"), " ", "")
		if name == "" {
			errs.Add(errors.New(http.StatusBadRequest, "name", ErrNameInvalid.Error()))
		}

		// Get folder name
		folder := r.FormValue("folder")
		if folder == "" {
			errs.Add(errors.New(http.StatusBadRequest, "folder", ErrFolderInvalid.Error()))
		}

		// Return if there were errors
		if errs.Length() > 0 {
			errors.Multiple(ac.Logger, w, http.StatusBadRequest, errs)
			return
		}

		// Get the file if provided
		file, handler, err := r.FormFile("image")

		// Create new variable for file extension
		var ext string

		// No file was provided, use placeholders
		if file == nil || handler == nil {
			file, err = os.Open("./../../app/public/images/" + folder + "/placeholder.jpg")
			if err != nil {
				errors.Default(ac.Logger, w, errors.ErrInternalServerError)
			}
			ext = ".jpg"
		} else if err != nil {
			errors.Default(ac.Logger, w, errors.ErrInternalServerError)
			return
		} else {
			ext = "." + strings.Split(handler.Filename, ".")[1]
		}
		defer file.Close()

		// Create file at source location
		appPath := "./../../app/public/images/" + folder + "/" + name + ext
		appFile, err := os.Create(appPath)
		if err != nil {
			errors.Default(ac.Logger, w, errors.ErrInternalServerError)
			return
		}
		defer appFile.Close()

		// Move uploaded file to appFile
		size, err := io.Copy(appFile, file)
		if size == 0 || err != nil {
			errors.Default(ac.Logger, w, errors.ErrInternalServerError)
			return
		}

		// Open file placed in app
		in, err := os.Open(appPath)
		if err != nil {
			errors.Default(ac.Logger, w, errors.ErrInternalServerError)
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
		if size == 0 || err != nil {
			errors.Default(ac.Logger, w, errors.ErrInternalServerError)
			return
		}

		// Create a new result
		result := ResultUpload{
			Data: true,
		}

		// Render output
		if err := render.JSON(w, true, result); err != nil {
			ac.Logger.Printf("render.JSON() error: %s\n", err)
			errors.Default(ac.Logger, w, errors.ErrInternalServerError)
			return
		}

	}
}
