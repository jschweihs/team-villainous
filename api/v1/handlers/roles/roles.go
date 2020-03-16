package roles

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	apictx "vil/api/context"
	"vil/api/errors"
	"vil/api/middleware/auth"
	"vil/api/render"
	serverrors "vil/services/errors"
	servroles "vil/services/roles"

	"github.com/beeker1121/httprouter"
)

// Role defines a role
type Role struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
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
	Data *Role `json:"data"`
}

// ResultGet defines the response data for the HandleGet handler
type ResultGet struct {
	Data  []*Role `json:"data"`
	Meta  Meta    `json:"meta"`
	Links Links   `json:"links"`
}

type ResultGetRole struct {
	Data *Role `json:"data"`
}

type ResultUpdate struct {
	Data *Role `json:"data"`
}

type DataDelete struct {
	ID int `json:"id"`
}
type ResultDelete struct {
	Data DataDelete `json:"data"`
}

// New creates the routes for the roles endpoints of the API
func New(ac *apictx.Context, router *httprouter.Router) {
	// Handle the routes
	router.POST("/api/v1/roles", auth.AuthenticateEndpoint(ac, HandlePost(ac)))
	router.GET("/api/v1/roles", HandleGet(ac))
	router.GET("/api/v1/roles/:id", HandleGetRole(ac))
	router.PUT("/api/v1/roles/:id", auth.AuthenticateEndpoint(ac, HandleUpdate(ac)))
	router.DELETE("/api/v1/roles/:id", auth.AuthenticateEndpoint(ac, HandleDelete(ac)))
}

// HandlePost handles the /api/v1/roles POST route of the API
func HandlePost(ac *apictx.Context) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Parse the parameters from the request body
		var params servroles.NewParams
		if err := json.NewDecoder(r.Body).Decode(&params); err != nil {
			fmt.Printf("Could not decode json")
			errors.Default(ac.Logger, w, errors.ErrBadRequest)
			return
		}

		// Get this user from the request context
		// This validates that this user exists and has their token
		_, err := auth.GetUserFromRequest(r)
		if err != nil {
			errors.Default(ac.Logger, w, errors.ErrInternalServerError)
			return
		}

		// Try to create a new role
		role, err := ac.Services.Roles.New(&params)
		if pes, ok := err.(*serverrors.ParamErrors); ok && err != nil {
			errors.Params(ac.Logger, w, http.StatusBadRequest, pes)
			return
		} else if err != nil {
			ac.Logger.Printf("roles.New() service error: %s\n", err)
			errors.Default(ac.Logger, w, errors.ErrInternalServerError)
			return
		}

		// Create a new Result
		result := ResultPost{
			Data: &Role{
				ID:   role.ID,
				Name: role.Name,
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

// HandleGetAll handles the /api/v1/roles GET route of the API
func HandleGet(ac *apictx.Context) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		// Get this user from the request context
		// This validates that this user exists and has their token
		// _, err := auth.GetUserFromRequest(r)
		// if err != nil {
		// 	errors.Default(ac.Logger, w, errors.ErrInternalServerError)
		// 	return
		// }

		// Create a new GetParams
		params := &servroles.GetParams{}

		// Create a new API Errors
		errs := &errors.Errors{}

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

		// Try to get the roles
		roles, err := ac.Services.Roles.Get(params)
		if err != nil {
			ac.Logger.Printf("roles.Get() service error: %s\n", err)
			errors.Default(ac.Logger, w, errors.ErrInternalServerError)
			return
		}

		// Create a new Result
		result := ResultGet{
			Data: []*Role{},
			Meta: Meta{
				Offset: params.Offset,
				Limit:  params.Limit,
				Total:  roles.Total,
			},
			Links: Links{},
		}

		// Loop through the roles
		for _, r := range roles.Roles {
			// Copy the Role type over
			role := &Role{
				ID:   r.ID,
				Name: r.Name,
			}

			result.Data = append(result.Data, role)
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

			prev := "//" + ac.Config.APIHost + "/api/v1/roles" + offsetstr + limitstr
			result.Links.Prev = &prev
		}

		// Handle next link.
		if params.Offset+params.Limit < roles.Total {
			offsetstr := "?offset=" + strconv.FormatInt(int64(params.Offset+params.Limit), 10)
			limitstr := "&limit=" + strconv.FormatInt(int64(params.Limit), 10)

			next := "//" + ac.Config.APIHost + "/api/v1/roles" + offsetstr + limitstr
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

// HandleGetAll handles the /api/v1/roles GET route of the API
func HandleGetRole(ac *apictx.Context) http.HandlerFunc {
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

		// Try to get this role
		role, err := ac.Services.Roles.GetByID(id)
		if err == servroles.ErrRoleNotFound {
			errors.Default(ac.Logger, w, errors.New(http.StatusNotFound, "", err.Error()))
			return
		} else if err != nil {
			ac.Logger.Printf("roles.GetByID() service error: %s\n", err)
			errors.Default(ac.Logger, w, errors.ErrInternalServerError)
			return
		}

		// Create a new result
		result := ResultGetRole{
			Data: &Role{
				ID:   role.ID,
				Name: role.Name,
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

// HandleUpdate handles the /api/v1/roles/:id PUT route of the API
func HandleUpdate(ac *apictx.Context) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Parse the parameters from the request body
		var params servroles.UpdateParams
		if err := json.NewDecoder(r.Body).Decode(&params); err != nil {
			errors.Default(ac.Logger, w, errors.ErrBadRequest)
			return
		}

		// Try to get the role ID
		var id int
		id64, err := strconv.ParseInt(httprouter.GetParam(r, "id"), 10, 32)
		if err != nil {
			errors.Default(ac.Logger, w, errors.ErrBadRequest)
			return
		}
		id = int(id64)

		// Get this user from the request context
		// This validates that this user exists and has their token
		_, err = auth.GetUserFromRequest(r)
		if err != nil {
			errors.Default(ac.Logger, w, errors.ErrInternalServerError)
			return
		}

		// Try to update this role
		role, err := ac.Services.Roles.UpdateByID(id, &params)
		if pes, ok := err.(*serverrors.ParamErrors); ok && err != nil {
			errors.Params(ac.Logger, w, http.StatusBadRequest, pes)
			return
		} else if err == servroles.ErrRoleNotFound {
			errors.Default(ac.Logger, w, errors.New(http.StatusNotFound, "", err.Error()))
			return
		} else if err != nil {
			ac.Logger.Printf("*roles.New() service error: %s\n", err)
			errors.Default(ac.Logger, w, errors.ErrInternalServerError)
			return
		}

		// Create a new Result
		result := ResultUpdate{
			Data: &Role{
				ID:   role.ID,
				Name: role.Name,
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

// HandleUpdate handles the /api/v1/roles/:id PUT route of the API
func HandleDelete(ac *apictx.Context) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		// Try to get the role ID
		var id int
		id64, err := strconv.ParseInt(httprouter.GetParam(r, "id"), 10, 32)
		if err != nil {
			errors.Default(ac.Logger, w, errors.ErrBadRequest)
			return
		}
		id = int(id64)

		// Get this user from the request context
		// This validates that this user exists and has their token
		_, err = auth.GetUserFromRequest(r)
		if err != nil {
			errors.Default(ac.Logger, w, errors.ErrInternalServerError)
			return
		}

		// Try to delete this role
		err = ac.Services.Roles.DeleteByID(id)
		if pes, ok := err.(*serverrors.ParamErrors); ok && err != nil {
			errors.Params(ac.Logger, w, http.StatusBadRequest, pes)
			return
		} else if err == servroles.ErrRoleNotFound {
			errors.Default(ac.Logger, w, errors.New(http.StatusNotFound, "", err.Error()))
			return
		} else if err != nil {
			ac.Logger.Printf("*roles.DeleteByID() service error: %s\n", err)
			errors.Default(ac.Logger, w, errors.ErrInternalServerError)
			return
		}

		// Create a new Result
		result := ResultDelete{
			Data: DataDelete{
				ID: id,
			},
		}
		fmt.Printf("%v", result)
		// Render output
		if err := render.JSON(w, true, result); err != nil {
			ac.Logger.Printf("render.JSON() error: %s\n", err)
			errors.Default(ac.Logger, w, errors.ErrInternalServerError)
			return
		}
	}
}
