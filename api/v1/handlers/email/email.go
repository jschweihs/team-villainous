package email

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/smtp"

	apictx "vil/api/context"
	"vil/api/errors"
	"vil/api/render"

	"github.com/jschweihs/httprouter"
)

// Note: Right now this package is only designed for handling images

// TODO: Implement type for upload and create it's own service
type ContactParams struct {
	Name     *string
	Email    *string
	Category *string
	Message  *string
}

// ResultSend defines the response data for HandlePost
type ResultSend struct {
	Data bool `json:"data"`
}

// New creates the routes for the blog entries endpoints of the API
func New(ac *apictx.Context, router *httprouter.Router) {
	// Handle the routes
	router.POST("/api/v1/email", HandlePost(ac))
}

type smtpServer struct {
	host string
	port string
}

func (s *smtpServer) Address() string {
	return s.host + ":" + s.port
}

// HandlePost handles the /api/v1/upload POST route of the API
// If no file is provided, this will attempt to use the placeholder
// file in the given folder instead
func HandlePost(ac *apictx.Context) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Get this user from the request context
		// This validates that this user exists and has their token
		// _, err := auth.GetUserFromRequest(r)
		// if err != nil {
		// 	errors.Default(ac.Logger, w, errors.ErrInternalServerError)
		// 	return
		// }

		// Parse the parameters from the request body
		var params ContactParams
		if err := json.NewDecoder(r.Body).Decode(&params); err != nil {
			errors.Default(ac.Logger, w, errors.ErrBadRequest)
			return
		}

		// Create a new API Errors
		// errs := &errors.Errors{}

		// Sender data
		from := "villainousteam2018@gmail.com"
		password := "vmxlkojmlxngjemm" // TODO: Move this to config

		// Receiver email address
		to := []string{
			"villainousteam2018@gmail.com",
		}

		// smtp server configuration
		smtpServer := smtpServer{
			host: "smtp.gmail.com",
			port: "587",
		}

		// Message
		m := fmt.Sprintf("To: villainousteam2018@gmail.com\r\nSubject: %v\r\n\r\n%v\r\n%v\r\n\r\n%v", *params.Category, *params.Email, *params.Name, *params.Message)
		message := []byte(m)
		// Authentication
		auth := smtp.PlainAuth("", from, password, smtpServer.host)

		// Sending email
		err := smtp.SendMail(smtpServer.Address(), auth, from, to, message)
		if err != nil {
			fmt.Println(err)
			return
		}

		result := ResultSend{
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
