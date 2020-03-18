package users

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"time"

	apictx "vil/api/context"
	"vil/api/errors"
	"vil/api/middleware/auth"
	"vil/api/render"
	serverrors "vil/services/errors"
	servusers "vil/services/users"

	"github.com/beeker1121/httprouter"
)

// User defines a user
type User struct {
	ID            int       `json:"id"`
	Username      string    `json:"username"`
	Password      string    `json:"password"`
	Email         string    `json:"email"`
	FName         string    `json:"f_name"`
	MName         string    `json:"m_name"`
	LName         string    `json:"l_name"`
	Title         string    `json:"title"`
	Address       string    `json:"address"`
	City          string    `json:"city"`
	Province      string    `json:"province"`
	Zip           string    `json:"zip"`
	Country       string    `json:"country"`
	BirthDate     string    `json:"birth_date"`
	Description   string    `json:"description"`
	Role          int       `json:"role"`
	PrivilegeID   int       `json:"privilege_id"`
	Status        int       `json:"status"`
	FacebookURL   string    `json:"facebook_url"`
	TwitterURL    string    `json:"twitter_url"`
	InstagramURL  string    `json:"instagram_url"`
	TwitchURL     string    `json:"twitch_url"`
	YoutubeURL    string    `json:"youtube_url"`
	OtherURL      string    `json:"other_url"`
	PS4Gamertag   string    `json:"ps4_gamertag"`
	XBoxGamertag  string    `json:"xbox_gamertag"`
	SteamGamertag string    `json:"steam_gamertag"`
	CreatedAt     time.Time `json:"created_at"`
	UpdatedAt     time.Time `json:"updated_at"`
}

// Meta defines the response top level meta object.
type Meta struct {
	Offset int `json:"offset"`
	Limit  int `json:"limit"`
	Total  int `json:"total"`
}

// Links defines the response top level links object.
type Links struct {
	Prev *string `json:"prev"`
	Next *string `json:"next"`
}

// ResultPost defines the response data for the HandlePost handler.
type ResultPost struct {
	Data *User `json:"data"`
}

// ResultGet defines the response data for the HandleGet handler
type ResultGet struct {
	Data  []*User `json:"data"`
	Meta  Meta    `json:"meta"`
	Links Links   `json:"links"`
}

type ResultGetUser struct {
	Data *User `json:"data"`
}

type ResultUpdate struct {
	Data *User `json:"data"`
}

type DataDelete struct {
	ID int `json:"id"`
}
type ResultDelete struct {
	Data DataDelete `json:"data"`
}

// New creates the routes for the users endpoints of the API
func New(ac *apictx.Context, router *httprouter.Router) {
	// Handle the routes
	router.POST("/api/v1/users", auth.AuthenticateEndpoint(ac, HandlePost(ac)))
	router.GET("/api/v1/users", HandleGet(ac))
	router.GET("/api/v1/users/:id", HandleGetUser(ac))
	router.GET("/api/v1/me", auth.AuthenticateEndpoint(ac, HandleGetMe(ac)))
	router.PUT("/api/v1/users/:id", auth.AuthenticateEndpoint(ac, HandleUpdate(ac)))
	router.DELETE("/api/v1/users/:id", auth.AuthenticateEndpoint(ac, HandleDelete(ac)))
}

// HandlePost handles the /api/v1/users POST route of the API
func HandlePost(ac *apictx.Context) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Parse the parameters from the request body
		var params servusers.NewParams
		if err := json.NewDecoder(r.Body).Decode(&params); err != nil {
			fmt.Println("Here1")
			errors.Default(ac.Logger, w, errors.ErrBadRequest)
			return
		}

		_, err := auth.GetUserFromRequest(r)
		if err != nil {
			fmt.Println("Here2")
			errors.Default(ac.Logger, w, errors.ErrInternalServerError)
			return
		}

		// Try to create a new user
		user, err := ac.Services.Users.New(&params)
		if pes, ok := err.(*serverrors.ParamErrors); ok && err != nil {
			fmt.Printf("%v\n", pes)
			errors.Params(ac.Logger, w, http.StatusBadRequest, pes)
			return
		} else if err != nil {
			fmt.Println("Here4")
			fmt.Printf("%v\n", err)
			ac.Logger.Printf("users.New() service error: %s\n", err)
			errors.Default(ac.Logger, w, errors.ErrInternalServerError)
			return
		}

		// Create a new Result
		result := ResultPost{
			Data: &User{
				ID:            user.ID,
				Username:      user.Username,
				Password:      user.Password,
				Email:         user.Email,
				FName:         user.FName,
				MName:         user.MName,
				LName:         user.LName,
				Address:       user.Address,
				City:          user.City,
				Province:      user.Province,
				Zip:           user.Zip,
				Country:       user.Country,
				BirthDate:     user.BirthDate,
				Description:   user.Description,
				Role:          user.Role,
				PrivilegeID:   user.PrivilegeID,
				Status:        user.Status,
				FacebookURL:   user.FacebookURL,
				TwitterURL:    user.TwitterURL,
				InstagramURL:  user.InstagramURL,
				TwitchURL:     user.TwitchURL,
				YoutubeURL:    user.YoutubeURL,
				OtherURL:      user.OtherURL,
				PS4Gamertag:   user.PS4Gamertag,
				XBoxGamertag:  user.XBoxGamertag,
				SteamGamertag: user.SteamGamertag,
				CreatedAt:     user.CreatedAt,
				UpdatedAt:     user.UpdatedAt,
			},
		}

		// Render output
		if err := render.JSON(w, true, result); err != nil {
			fmt.Println("Here5")
			ac.Logger.Printf("render.JSON() error: %s\n", err)
			errors.Default(ac.Logger, w, errors.ErrInternalServerError)
			return
		}
	}
}

// HandleGetAll handles the /api/v1/users GET route of the API
func HandleGet(ac *apictx.Context) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Get this user from the request context
		// _, err := auth.GetUserFromRequest(r)
		// if err != nil {
		// 	errors.Default(ac.Logger, w, errors.ErrInternalServerError)
		// 	return
		// }

		// Create a new GetParams
		params := &servusers.GetParams{}

		// Create a new API Errors
		errs := &errors.Errors{}

		// Handle created
		if createdqs, ok := r.URL.Query()["created_at"]; ok && len(createdqs) == 1 {
			t, err := time.Parse(time.RFC3339, createdqs[0])
			if err != nil {
				errs.Add(errors.New(http.StatusBadRequest, "created", ErrCreatedInvalid.Error()))
			} else {
				params.CreatedAt = &t
			}
		}

		// Handle created
		if statusqs, ok := r.URL.Query()["status"]; ok && len(statusqs) == 1 {
			status, err := strconv.ParseInt(statusqs[0], 10, 32)
			if err != nil {
				errs.Add(errors.New(http.StatusBadRequest, "status", ErrLimitInvalid.Error()))
			} else {
				params.Status = int(status)
			}
		} else {
			params.Status = 0
		}

		// Handle offset
		if offsetqs, ok := r.URL.Query()["offset"]; ok && len(offsetqs) == 1 {
			offset64, err := strconv.ParseInt(offsetqs[0], 10, 32)
			if err != nil {
				errs.Add(errors.New(http.StatusBadRequest, "offset", ErrLimitInvalid.Error()))
			} else {
				params.Offset = int(offset64)
			}
		} else {
			params.Offset = 0
		}

		// Handle limit
		if limitqs, ok := r.URL.Query()["limit"]; ok && len(limitqs) == 1 {
			limit64, err := strconv.ParseInt(limitqs[0], 10, 32)
			if err != nil {
				errs.Add(errors.New(http.StatusBadRequest, "limit", ErrLimitInvalid.Error()))
			} else {
				if int(limit64) > ac.Config.LimitMax {
					errs.Add(errors.New(http.StatusBadRequest, "limit", ErrLimitMax.Error()+" of "+strconv.FormatUint(uint64(ac.Config.LimitMax), 10)))
				} else {
					params.Limit = int(limit64)
				}
			}
		} else {
			params.Limit = ac.Config.LimitDefault
		}

		// Return if there were errors
		if errs.Length() > 0 {
			errors.Multiple(ac.Logger, w, http.StatusBadRequest, errs)
			return
		}

		// Try to get the users
		users, err := ac.Services.Users.Get(params)
		if err != nil {
			ac.Logger.Printf("users.Get() service error: %s\n", err)
			errors.Default(ac.Logger, w, errors.ErrInternalServerError)
			return
		}

		// Create a new Result
		result := ResultGet{
			Data: []*User{},
			Meta: Meta{
				Offset: params.Offset,
				Limit:  params.Limit,
				Total:  users.Total,
			},
			Links: Links{},
		}

		// Loop through the users
		for _, u := range users.Users {
			// Copy the User type over
			user := &User{
				ID:            u.ID,
				Username:      u.Username,
				Password:      u.Password,
				Email:         u.Email,
				FName:         u.FName,
				MName:         u.MName,
				LName:         u.LName,
				Address:       u.Address,
				City:          u.City,
				Province:      u.Province,
				Zip:           u.Zip,
				Country:       u.Country,
				BirthDate:     u.BirthDate,
				Description:   u.Description,
				Role:          u.Role,
				PrivilegeID:   u.PrivilegeID,
				Status:        u.Status,
				FacebookURL:   u.FacebookURL,
				TwitterURL:    u.TwitterURL,
				InstagramURL:  u.InstagramURL,
				TwitchURL:     u.TwitchURL,
				YoutubeURL:    u.YoutubeURL,
				OtherURL:      u.OtherURL,
				PS4Gamertag:   u.PS4Gamertag,
				XBoxGamertag:  u.XBoxGamertag,
				SteamGamertag: u.SteamGamertag,
				CreatedAt:     u.CreatedAt,
				UpdatedAt:     u.UpdatedAt,
			}

			result.Data = append(result.Data, user)
		}

		// Handle previous link
		if params.Offset > 0 {
			limitstr := "&limit=" + strconv.FormatInt(int64(params.Limit), 10)

			offsetstr := "?offset="
			if params.Offset-params.Limit < 0 {
				offsetstr += "0"
			} else {
				offsetstr += strconv.FormatInt(int64(params.Offset-params.Limit), 10)
			}

			prev := "//" + ac.Config.APIHost + "/api/v1/users" + offsetstr + limitstr
			result.Links.Prev = &prev
		}

		// Handle next link.
		if params.Offset+params.Limit < users.Total {
			offsetstr := "?offset=" + strconv.FormatInt(int64(params.Offset+params.Limit), 10)
			limitstr := "&limit=" + strconv.FormatInt(int64(params.Limit), 10)

			next := "//" + ac.Config.APIHost + "/api/v1/users" + offsetstr + limitstr
			result.Links.Next = &next
		}

		// Render output.
		if err := render.JSON(w, true, result); err != nil {
			ac.Logger.Printf("render.JSON() error: %s\n", err)
			errors.Default(ac.Logger, w, errors.ErrInternalServerError)
			return
		}

	}
}

// HandleGetUser handles the /api/v1/users GET route of the API
func HandleGetUser(ac *apictx.Context) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var id int
		id64, err := strconv.ParseInt(httprouter.GetParam(r, "id"), 10, 32)
		if err != nil {
			errors.Default(ac.Logger, w, errors.ErrBadRequest)
			return
		}
		id = int(id64)

		// Get this user from the request context
		// This validates that this user exists and has their token
		// _, err = auth.GetUserFromRequest(r)
		// if err != nil {
		// 	errors.Default(ac.Logger, w, errors.ErrInternalServerError)
		// 	return
		// }

		// Try to get this user
		user, err := ac.Services.Users.GetByID(id)
		if err == servusers.ErrUserNotFound {
			errors.Default(ac.Logger, w, errors.New(http.StatusNotFound, "", err.Error()))
			return
		} else if err != nil {
			ac.Logger.Printf("users.GetByID() service error: %s\n", err)
			errors.Default(ac.Logger, w, errors.ErrInternalServerError)
			return
		}

		// Create a new result
		result := ResultGetUser{
			Data: &User{
				ID:            user.ID,
				Username:      user.Username,
				Password:      user.Password,
				Email:         user.Email,
				FName:         user.FName,
				MName:         user.MName,
				LName:         user.LName,
				Title:         user.Title,
				Address:       user.Address,
				City:          user.City,
				Province:      user.Province,
				Zip:           user.Zip,
				Country:       user.Country,
				BirthDate:     user.BirthDate,
				Description:   user.Description,
				Role:          user.Role,
				PrivilegeID:   user.PrivilegeID,
				Status:        user.Status,
				FacebookURL:   user.FacebookURL,
				TwitterURL:    user.TwitterURL,
				InstagramURL:  user.InstagramURL,
				TwitchURL:     user.TwitchURL,
				YoutubeURL:    user.YoutubeURL,
				OtherURL:      user.OtherURL,
				PS4Gamertag:   user.PS4Gamertag,
				XBoxGamertag:  user.XBoxGamertag,
				SteamGamertag: user.SteamGamertag,
				CreatedAt:     user.CreatedAt,
				UpdatedAt:     user.UpdatedAt,
			},
		}

		// Render output.
		if err := render.JSON(w, true, result); err != nil {
			ac.Logger.Printf("render.JSON() error: %s\n", err)
			errors.Default(ac.Logger, w, errors.ErrInternalServerError)
			return
		}

	}
}

// HandleGetMe handles the /api/v1/me GET route of the API
func HandleGetMe(ac *apictx.Context) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Get this user from the request context
		// This validates that this user exists and has their token
		user, err := auth.GetUserFromRequest(r)
		if err != nil {
			errors.Default(ac.Logger, w, errors.ErrInternalServerError)
			return
		}

		// Create a new result
		result := ResultGetUser{
			Data: &User{
				ID:            user.ID,
				Username:      user.Username,
				Password:      user.Password,
				Email:         user.Email,
				FName:         user.FName,
				MName:         user.MName,
				LName:         user.LName,
				Title:         user.Title,
				Address:       user.Address,
				City:          user.City,
				Province:      user.Province,
				Zip:           user.Zip,
				Country:       user.Country,
				BirthDate:     user.BirthDate,
				Description:   user.Description,
				Role:          user.Role,
				PrivilegeID:   user.PrivilegeID,
				Status:        user.Status,
				FacebookURL:   user.FacebookURL,
				TwitterURL:    user.TwitterURL,
				InstagramURL:  user.InstagramURL,
				TwitchURL:     user.TwitchURL,
				YoutubeURL:    user.YoutubeURL,
				OtherURL:      user.OtherURL,
				PS4Gamertag:   user.PS4Gamertag,
				XBoxGamertag:  user.XBoxGamertag,
				SteamGamertag: user.SteamGamertag,
				CreatedAt:     user.CreatedAt,
				UpdatedAt:     user.UpdatedAt,
			},
		}

		// Render output.
		if err := render.JSON(w, true, result); err != nil {
			ac.Logger.Printf("render.JSON() error: %s\n", err)
			errors.Default(ac.Logger, w, errors.ErrInternalServerError)
			return
		}
	}
}

// HandleUpdate handles the /api/v1/users/:id PUT route of the API
func HandleUpdate(ac *apictx.Context) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Parse the parameters from the request body
		var params servusers.UpdateParams
		if err := json.NewDecoder(r.Body).Decode(&params); err != nil {
			errors.Default(ac.Logger, w, errors.ErrBadRequest)
			return
		}

		// Try to get the user ID
		var id int
		id64, err := strconv.ParseInt(httprouter.GetParam(r, "id"), 10, 32)
		if err != nil {
			errors.Default(ac.Logger, w, errors.ErrBadRequest)
			return
		}
		id = int(id64)

		// Get this user from the request context
		_, err = auth.GetUserFromRequest(r)
		if err != nil {
			errors.Default(ac.Logger, w, errors.ErrInternalServerError)
			return
		}

		// Try to update this user
		user, err := ac.Services.Users.UpdateByID(id, &params)
		if pes, ok := err.(*serverrors.ParamErrors); ok && err != nil {
			errors.Params(ac.Logger, w, http.StatusBadRequest, pes)
			return
		} else if err == servusers.ErrUserNotFound {
			errors.Default(ac.Logger, w, errors.New(http.StatusNotFound, "", err.Error()))
			return
		} else if err != nil {
			ac.Logger.Printf("*users.New() service error: %s\n", err)
			errors.Default(ac.Logger, w, errors.ErrInternalServerError)
			return
		}

		// Create a new Result
		result := ResultUpdate{
			Data: &User{
				ID:            user.ID,
				Username:      user.Username,
				Password:      user.Password,
				Email:         user.Email,
				FName:         user.FName,
				MName:         user.MName,
				LName:         user.LName,
				Title:         user.Title,
				Address:       user.Address,
				City:          user.City,
				Province:      user.Province,
				Zip:           user.Zip,
				Country:       user.Country,
				BirthDate:     user.BirthDate,
				Description:   user.Description,
				Role:          user.Role,
				PrivilegeID:   user.PrivilegeID,
				Status:        user.Status,
				FacebookURL:   user.FacebookURL,
				TwitterURL:    user.TwitterURL,
				InstagramURL:  user.InstagramURL,
				TwitchURL:     user.TwitchURL,
				YoutubeURL:    user.YoutubeURL,
				OtherURL:      user.OtherURL,
				PS4Gamertag:   user.PS4Gamertag,
				XBoxGamertag:  user.XBoxGamertag,
				SteamGamertag: user.SteamGamertag,
				CreatedAt:     user.CreatedAt,
				UpdatedAt:     user.UpdatedAt,
			},
		}

		// Render output
		if err := render.JSON(w, true, result); err != nil {
			ac.Logger.Printf("render.JSON() error: %s\n", err)
			errors.Default(ac.Logger, w, errors.ErrInternalServerError)
			return
		}
	}
}

// HandleUpdate handles the /api/v1/users/:id DELETE route of the API
func HandleDelete(ac *apictx.Context) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		// Try to get the user ID
		var id int
		id64, err := strconv.ParseInt(httprouter.GetParam(r, "id"), 10, 32)
		if err != nil {
			errors.Default(ac.Logger, w, errors.ErrBadRequest)
			return
		}
		id = int(id64)

		fmt.Printf("%v\n", id)

		// Get this user from the request context
		_, err = auth.GetUserFromRequest(r)
		if err != nil {
			errors.Default(ac.Logger, w, errors.ErrInternalServerError)
			return
		}

		// Try to delete this user
		err = ac.Services.Users.DeleteByID(id)
		if pes, ok := err.(*serverrors.ParamErrors); ok && err != nil {
			errors.Params(ac.Logger, w, http.StatusBadRequest, pes)
			return
		} else if err == servusers.ErrUserNotFound {
			errors.Default(ac.Logger, w, errors.New(http.StatusNotFound, "", err.Error()))
			return
		} else if err != nil {
			ac.Logger.Printf("*users.DeleteByID() service error: %s\n", err)
			errors.Default(ac.Logger, w, errors.ErrInternalServerError)
			return
		}

		// Create a new Result
		result := ResultDelete{
			Data: DataDelete{
				ID: id,
			},
		}

		// Render output
		if err := render.JSON(w, true, result); err != nil {
			ac.Logger.Printf("render.JSON() error: %s\n", err)
			errors.Default(ac.Logger, w, errors.ErrInternalServerError)
			return
		}
	}
}
