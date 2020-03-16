package blog

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
	servblog "vil/services/blog"
	serverrors "vil/services/errors"

	"github.com/beeker1121/httprouter"
)

// BlogEntry defines a blog entry
type BlogEntry struct {
	ID        int       `json:"id"`
	Title     string    `json:"title"`
	UserID    int       `json:"user_id"`
	Preview   string    `json:"preview"`
	Content   string    `json:"content"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
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
	Data *BlogEntry `json:"data"`
}

// ResultGet defines the response data for the HandleGet handler
type ResultGet struct {
	Data  []*BlogEntry `json:"data"`
	Meta  Meta         `json:"meta"`
	Links Links        `json:"links"`
}

type ResultGetBlogEntry struct {
	Data *BlogEntry `json:"data"`
}

type ResultUpdate struct {
	Data *BlogEntry `json:"data"`
}

type DataDelete struct {
	ID int `json:"id"`
}
type ResultDelete struct {
	Data DataDelete `json:"data"`
}

// New creates the routes for the blog entries endpoints of the API
func New(ac *apictx.Context, router *httprouter.Router) {
	// Handle the routes
	router.POST("/api/v1/blog", auth.AuthenticateEndpoint(ac, HandlePost(ac)))
	router.GET("/api/v1/blog", HandleGet(ac))
	router.GET("/api/v1/blog/:id", HandleGetBlogEntry(ac))
	router.PUT("/api/v1/blog/:id", auth.AuthenticateEndpoint(ac, HandleUpdate(ac)))
	router.DELETE("/api/v1/blog/:id", auth.AuthenticateEndpoint(ac, HandleDelete(ac)))
}

// HandlePost handles the /api/v1/blog POST route of the API
func HandlePost(ac *apictx.Context) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Parse the parameters from the request body
		var params servblog.NewParams
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

		// Try to create a new blog
		blogEntry, err := ac.Services.Blog.New(&params)
		if pes, ok := err.(*serverrors.ParamErrors); ok && err != nil {
			fmt.Println("Here")
			errors.Params(ac.Logger, w, http.StatusBadRequest, pes)
			return
		} else if err != nil {
			ac.Logger.Printf("blog.New() service error: %s\n", err)
			errors.Default(ac.Logger, w, errors.ErrInternalServerError)
			return
		}

		// Create a new Result
		result := ResultPost{
			Data: &BlogEntry{
				ID:        blogEntry.ID,
				Title:     blogEntry.Title,
				UserID:    blogEntry.UserID,
				Preview:   blogEntry.Preview,
				Content:   blogEntry.Content,
				CreatedAt: blogEntry.CreatedAt,
				UpdatedAt: blogEntry.UpdatedAt,
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

// HandleGetAll handles the /api/v1/blog GET route of the API
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
		params := &servblog.GetParams{}

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

		// Try to get the blog entries
		blogEntries, err := ac.Services.Blog.Get(params)
		if err != nil {
			ac.Logger.Printf("blogEntries.Get() service error: %s\n", err)
			errors.Default(ac.Logger, w, errors.ErrInternalServerError)
			return
		}

		// Create a new Result
		result := ResultGet{
			Data: []*BlogEntry{},
			Meta: Meta{
				Offset: params.Offset,
				Limit:  params.Limit,
				Total:  blogEntries.Total,
			},
			Links: Links{},
		}

		// Loop through the blog entries
		for _, be := range blogEntries.BlogEntries {
			// Copy the BlogEntry type over
			blogEntry := &BlogEntry{
				ID:        be.ID,
				Title:     be.Title,
				UserID:    be.UserID,
				Preview:   be.Preview,
				Content:   be.Content,
				CreatedAt: be.CreatedAt,
				UpdatedAt: be.UpdatedAt,
			}

			result.Data = append(result.Data, blogEntry)
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

			prev := "//" + ac.Config.APIHost + "/api/v1/blog" + offsetstr + limitstr
			result.Links.Prev = &prev
		}

		// Handle next link.
		if params.Offset+params.Limit < blogEntries.Total {
			offsetstr := "?offset=" + strconv.FormatInt(int64(params.Offset+params.Limit), 10)
			limitstr := "&limit=" + strconv.FormatInt(int64(params.Limit), 10)

			next := "//" + ac.Config.APIHost + "/api/v1/blog" + offsetstr + limitstr
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

// HandleGetAll handles the /api/v1/blog GET route of the API
func HandleGetBlogEntry(ac *apictx.Context) http.HandlerFunc {
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

		// Try to get this blog entry
		blogEntry, err := ac.Services.Blog.GetByID(id)
		if err == servblog.ErrBlogEntryNotFound {
			errors.Default(ac.Logger, w, errors.New(http.StatusNotFound, "", err.Error()))
			return
		} else if err != nil {
			ac.Logger.Printf("Blog.GetByID() service error: %s\n", err)
			errors.Default(ac.Logger, w, errors.ErrInternalServerError)
			return
		}

		// Create a new result
		result := ResultGetBlogEntry{
			Data: &BlogEntry{
				ID:        blogEntry.ID,
				Title:     blogEntry.Title,
				UserID:    blogEntry.UserID,
				Preview:   blogEntry.Preview,
				Content:   blogEntry.Content,
				CreatedAt: blogEntry.CreatedAt,
				UpdatedAt: blogEntry.UpdatedAt,
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

// HandleUpdate handles the /api/v1/blog/:id PUT route of the API
func HandleUpdate(ac *apictx.Context) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Parse the parameters from the request body
		var params servblog.UpdateParams
		if err := json.NewDecoder(r.Body).Decode(&params); err != nil {
			errors.Default(ac.Logger, w, errors.ErrBadRequest)
			return
		}

		// Try to get the blog ID
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

		// Try to update this blog entry
		blogEntry, err := ac.Services.Blog.UpdateByID(id, &params)
		if pes, ok := err.(*serverrors.ParamErrors); ok && err != nil {
			errors.Params(ac.Logger, w, http.StatusBadRequest, pes)
			return
		} else if err == servblog.ErrBlogEntryNotFound {
			errors.Default(ac.Logger, w, errors.New(http.StatusNotFound, "", err.Error()))
			return
		} else if err != nil {
			ac.Logger.Printf("*Blog.UpdateByID() service error: %s\n", err)
			errors.Default(ac.Logger, w, errors.ErrInternalServerError)
			return
		}

		// Create a new Result
		result := ResultUpdate{
			Data: &BlogEntry{
				ID:        blogEntry.ID,
				Title:     blogEntry.Title,
				UserID:    blogEntry.UserID,
				Preview:   blogEntry.Preview,
				Content:   blogEntry.Content,
				CreatedAt: blogEntry.CreatedAt,
				UpdatedAt: blogEntry.UpdatedAt,
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

// HandleUpdate handles the /api/v1/blog/:id PUT route of the API
func HandleDelete(ac *apictx.Context) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		// Try to get the blog ID
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

		// Try to delete this blog
		err = ac.Services.Blog.DeleteByID(id)
		if pes, ok := err.(*serverrors.ParamErrors); ok && err != nil {
			errors.Params(ac.Logger, w, http.StatusBadRequest, pes)
			return
		} else if err == servblog.ErrBlogEntryNotFound {
			errors.Default(ac.Logger, w, errors.New(http.StatusNotFound, "", err.Error()))
			return
		} else if err != nil {
			ac.Logger.Printf("*Blog.DeleteByID() service error: %s\n", err)
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
