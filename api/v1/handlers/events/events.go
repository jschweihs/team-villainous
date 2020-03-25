package events

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
	servevents "vil/services/events"

	"github.com/jschweihs/httprouter"
)

// Event defines an event
type Event struct {
	ID            int       `json:"id"`
	Name          string    `json:"name"`
	Location      string    `json:"location"`
	StartDatetime time.Time `json:"start_datetime"`
	ShowStartTime bool      `json:"show_start_time"`
	EndDatetime   time.Time `json:"end_datetime"`
	ShowEndTime   bool      `json:"show_end_time"`
	Type          int       `json:"type"`
	GameID        int       `json:"game_id"`
	Description   string    `json:"description"`
	ReferralURL   string    `json:"referralURL"`
	Status        int       `json:"status"`
	Placements    []int     `json:"placements"`
	Users         []int     `json:"users"`
	CreatedBy     int       `json:"created_by"`
	ModifiedBy    int       `json:"modified_by"`
	CreatedAt     time.Time `json:"created_at"`
	ModifiedAt    time.Time `json:"modified_at"`
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
	Data *Event `json:"data"`
}

// ResultGet defines the response data for the HandleGet handler
type ResultGet struct {
	Data  []*Event `json:"data"`
	Meta  Meta     `json:"meta"`
	Links Links    `json:"links"`
}

type ResultGetEvent struct {
	Data *Event `json:"data"`
}

type ResultUpdate struct {
	Data *Event `json:"data"`
}

type DataDelete struct {
	ID int `json:"id"`
}
type ResultDelete struct {
	Data DataDelete `json:"data"`
}

// New creates the routes for the events endpoints of the API
func New(ac *apictx.Context, router *httprouter.Router) {
	// Handle the routes
	router.POST("/api/v1/events", auth.AuthenticateEndpoint(ac, HandlePost(ac)))
	router.GET("/api/v1/events", HandleGet(ac))
	router.GET("/api/v1/events/:id", HandleGetEvent(ac))
	router.PUT("/api/v1/events/:id", auth.AuthenticateEndpoint(ac, HandleUpdate(ac)))
	router.DELETE("/api/v1/events/:id", auth.AuthenticateEndpoint(ac, HandleDelete(ac)))
}

// HandlePost handles the /api/v1/events POST route of the API
func HandlePost(ac *apictx.Context) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Parse the parameters from the request body
		var params servevents.NewParams
		if err := json.NewDecoder(r.Body).Decode(&params); err != nil {
			fmt.Println(err)
			errors.Default(ac.Logger, w, errors.ErrBadRequest)
			return
		}

		user, err := auth.GetUserFromRequest(r)
		if err != nil {
			fmt.Println(err)
			errors.Default(ac.Logger, w, errors.ErrInternalServerError)
			return
		}
		params.CurrentUser = user.ID

		// Try to create a new event
		event, err := ac.Services.Events.New(&params)
		if pes, ok := err.(*serverrors.ParamErrors); ok && err != nil {
			fmt.Println(err)
			errors.Params(ac.Logger, w, http.StatusBadRequest, pes)
			return
		} else if err != nil {
			fmt.Println(err)
			ac.Logger.Printf("events.New() service error: %s\n", err)
			errors.Default(ac.Logger, w, errors.ErrInternalServerError)
			return
		}

		// Create a new Result
		result := ResultPost{
			Data: &Event{
				ID:            event.ID,
				Name:          event.Name,
				Location:      event.Location,
				StartDatetime: event.StartDatetime,
				ShowStartTime: event.ShowStartTime,
				EndDatetime:   event.EndDatetime,
				ShowEndTime:   event.ShowEndTime,
				Type:          event.Type,
				GameID:        event.GameID,
				Description:   event.Description,
				ReferralURL:   event.ReferralURL,
				Status:        event.Status,
				Placements:    event.Placements,
				Users:         event.Users,
				CreatedBy:     event.CreatedBy,
				ModifiedBy:    event.ModifiedBy,
				CreatedAt:     event.CreatedAt,
				ModifiedAt:    event.ModifiedAt,
			},
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

// HandleGet handles the /api/v1/events GET route of the API
func HandleGet(ac *apictx.Context) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Create a new GetParams
		params := &servevents.GetParams{}

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

		// Handle status
		if statusqs, ok := r.URL.Query()["status"]; ok && len(statusqs) == 1 {
			status, err := strconv.ParseInt(statusqs[0], 10, 32)
			if err != nil {
				errs.Add(errors.New(http.StatusBadRequest, "status", ErrLimitInvalid.Error()))
			} else {
				i := int(status)
				params.Status = &i
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

		// Try to get the events
		events, err := ac.Services.Events.Get(params)
		if err != nil {
			ac.Logger.Printf("events.Get() service error: %s\n", err)
			errors.Default(ac.Logger, w, errors.ErrInternalServerError)
			return
		}

		// Create a new Result
		result := ResultGet{
			Data: []*Event{},
			Meta: Meta{
				Offset: params.Offset,
				Limit:  params.Limit,
				Total:  events.Total,
			},
			Links: Links{},
		}

		// Loop through the events
		for _, e := range events.Events {
			// Copy the Event type over
			event := &Event{
				ID:            e.ID,
				Name:          e.Name,
				Location:      e.Location,
				StartDatetime: e.StartDatetime,
				ShowStartTime: e.ShowStartTime,
				EndDatetime:   e.EndDatetime,
				ShowEndTime:   e.ShowEndTime,
				Type:          e.Type,
				GameID:        e.GameID,
				Description:   e.Description,
				ReferralURL:   e.ReferralURL,
				Status:        e.Status,
				Placements:    e.Placements,
				Users:         e.Users,
				CreatedBy:     e.CreatedBy,
				ModifiedBy:    e.ModifiedBy,
				CreatedAt:     e.CreatedAt,
				ModifiedAt:    e.ModifiedAt,
			}

			result.Data = append(result.Data, event)
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

			prev := "//" + ac.Config.APIHost + "/api/v1/events" + offsetstr + limitstr
			result.Links.Prev = &prev
		}

		// Handle next link.
		if params.Offset+params.Limit < events.Total {
			offsetstr := "?offset=" + strconv.FormatInt(int64(params.Offset+params.Limit), 10)
			limitstr := "&limit=" + strconv.FormatInt(int64(params.Limit), 10)

			next := "//" + ac.Config.APIHost + "/api/v1/events" + offsetstr + limitstr
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

// HandleGetEvent handles the /api/v1/events GET route of the API
func HandleGetEvent(ac *apictx.Context) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var id int
		id64, err := strconv.ParseInt(httprouter.GetParam(r, "id"), 10, 32)
		if err != nil {
			errors.Default(ac.Logger, w, errors.ErrBadRequest)
			return
		}
		id = int(id64)

		// Try to get this event
		event, err := ac.Services.Events.GetByID(id)
		if err == servevents.ErrEventNotFound {
			errors.Default(ac.Logger, w, errors.New(http.StatusNotFound, "", err.Error()))
			return
		} else if err != nil {
			ac.Logger.Printf("events.GetByID() service error: %s\n", err)
			errors.Default(ac.Logger, w, errors.ErrInternalServerError)
			return
		}

		// Create a new result
		result := ResultGetEvent{
			Data: &Event{
				ID:            event.ID,
				Name:          event.Name,
				Location:      event.Location,
				StartDatetime: event.StartDatetime,
				ShowStartTime: event.ShowStartTime,
				EndDatetime:   event.EndDatetime,
				ShowEndTime:   event.ShowEndTime,
				Type:          event.Type,
				GameID:        event.GameID,
				Description:   event.Description,
				ReferralURL:   event.ReferralURL,
				Status:        event.Status,
				Placements:    event.Placements,
				Users:         event.Users,
				CreatedBy:     event.CreatedBy,
				ModifiedBy:    event.ModifiedBy,
				CreatedAt:     event.CreatedAt,
				ModifiedAt:    event.ModifiedAt,
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

// HandleUpdate handles the /api/v1/events/:id PUT route of the API
func HandleUpdate(ac *apictx.Context) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Parse the parameters from the request body
		var params servevents.UpdateParams
		if err := json.NewDecoder(r.Body).Decode(&params); err != nil {
			errors.Default(ac.Logger, w, errors.ErrBadRequest)
			return
		}

		// Try to get the event ID
		var id int
		id64, err := strconv.ParseInt(httprouter.GetParam(r, "id"), 10, 32)
		if err != nil {
			errors.Default(ac.Logger, w, errors.ErrBadRequest)
			return
		}
		id = int(id64)

		// Get this user from the request context
		user, err := auth.GetUserFromRequest(r)
		if err != nil {
			errors.Default(ac.Logger, w, errors.ErrInternalServerError)
			return
		}
		uid := user.ID
		params.CurrentUser = &uid

		// Try to update this event
		event, err := ac.Services.Events.UpdateByID(id, &params)
		if pes, ok := err.(*serverrors.ParamErrors); ok && err != nil {
			errors.Params(ac.Logger, w, http.StatusBadRequest, pes)
			return
		} else if err == servevents.ErrEventNotFound {
			errors.Default(ac.Logger, w, errors.New(http.StatusNotFound, "", err.Error()))
			return
		} else if err != nil {
			ac.Logger.Printf("*events.New() service error: %s\n", err)
			errors.Default(ac.Logger, w, errors.ErrInternalServerError)
			return
		}

		// Create a new Result
		result := ResultUpdate{
			Data: &Event{
				ID:            event.ID,
				Name:          event.Name,
				Location:      event.Location,
				StartDatetime: event.StartDatetime,
				ShowStartTime: event.ShowStartTime,
				EndDatetime:   event.EndDatetime,
				ShowEndTime:   event.ShowEndTime,
				Type:          event.Type,
				GameID:        event.GameID,
				Description:   event.Description,
				ReferralURL:   event.ReferralURL,
				Status:        event.Status,
				Placements:    event.Placements,
				Users:         event.Users,
				CreatedBy:     event.CreatedBy,
				ModifiedBy:    event.ModifiedBy,
				CreatedAt:     event.CreatedAt,
				ModifiedAt:    event.ModifiedAt,
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

// HandleDelete handles the /api/v1/events/:id DELETE route of the API
func HandleDelete(ac *apictx.Context) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		// Try to get the event ID
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

		// Try to delete this event
		err = ac.Services.Events.DeleteByID(id)
		if pes, ok := err.(*serverrors.ParamErrors); ok && err != nil {
			errors.Params(ac.Logger, w, http.StatusBadRequest, pes)
			return
		} else if err == servevents.ErrEventNotFound {
			errors.Default(ac.Logger, w, errors.New(http.StatusNotFound, "", err.Error()))
			return
		} else if err != nil {
			ac.Logger.Printf("*events.DeleteByID() service error: %s\n", err)
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
