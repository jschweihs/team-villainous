package events

import (
	"database/sql"
	"fmt"
	"strconv"
	"strings"
	"time"
)

// Database defines the events database
type Database struct {
	db *sql.DB
}

// New creates a new events database
func New(db *sql.DB) *Database {
	return &Database{
		db: db,
	}
}

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
	CreatedBy     int       `json:"created_by"`
	ModifiedBy    int       `json:"modified_by"`
	CreatedAt     time.Time `json:"created_at"`
	ModifiedAt    time.Time `json:"modified_at"`
	Placements    []int     `json:"placements"`
	Users         []int     `json:"users"`
}

// Events defines a set of events
type Events struct {
	Events []*Event `json:"events"`
	Total  int      `json:"total"`
}

// Placement defines a placement
type Placement struct {
	EventID   int
	Placement int
}

// User defines a user
type User struct {
	EventID int
	UserID  int
}

const (
	// stmtInsert defines the SQL statement to
	// insert a new event into the database
	stmtInsert = `
INSERT INTO events (
	name,
	location,
	start_datetime,
	show_start_time,
	end_datetime,
	show_end_time,
	type,
	game_id,
	description,
	referral_url,
	status,
	created_by,
	modified_by,
	created_at,
	modified_at
)
VALUES (
	?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?
)
`

	// stmtInsertPlacement defines the SQL statement to
	// insert a new event into the database
	stmtInsertPlacement = `
INSERT INTO events_placements (
	event_id,
	placement
)
VALUES (
	?, ?
)
`

	// stmtInsertPlacement defines the SQL statement to
	// insert a new event into the database
	stmtInsertUser = `
INSERT INTO events_users (
	event_id,
	user_id
)
VALUES (
	?, ?
)
`

	// stmtSelect defines the SQL statement to
	// select a set of events.
	stmtSelect = `
SELECT 
	id,
	name,
	location,
	start_datetime,
	show_start_time,
	end_datetime,
	show_end_time,
	type,
	game_id,
	description,
	referral_url,
	status,
	created_by,
	modified_by,
	created_at,
	modified_at,
	COALESCE(GROUP_CONCAT(DISTINCT events_placements.placement), '') AS placements,
	COALESCE(GROUP_CONCAT(DISTINCT events_users.user_id), '') AS users
FROM events
LEFT JOIN events_placements ON events.id=events_placements.event_id
LEFT JOIN events_users ON events.id=events_users.event_id 
%s
GROUP BY events.id
LIMIT %v, %v
`

	// stmtSelectCount defines the SQL statement to
	// select the total number of events according to the filters.
	stmtSelectCount = `
SELECT COUNT(*)
FROM events
%s
`

	// stmtSelectByID defines the SQL statement to
	// select an event by their ID.
	stmtSelectByID = `
SELECT 
	id,
	name,
	location,
	start_datetime,
	show_start_time,
	end_datetime,
	show_end_time,
	type,
	game_id,
	description,
	referral_url,
	status,
	created_by,
	modified_by,
	created_at,
	modified_at,
	COALESCE(GROUP_CONCAT(DISTINCT events_placements.placement), '') AS placements,
	COALESCE(GROUP_CONCAT(DISTINCT events_users.user_id), '') AS users 
FROM events
LEFT JOIN events_placements ON events.id=events_placements.event_id
LEFT JOIN events_users ON events.id=events_users.event_id 
WHERE id=?
GROUP BY events.id
`

	// stmtUpdate defines the SQL statement to
	// update an event.
	stmtUpdate = `
UPDATE events
SET %s
WHERE id=?
`

	// stmtDelete defines the SQL statement to
	// remove an event.
	stmtDelete = `
UPDATE events
SET status=2
WHERE id=?
`

	// stmtDeletePlacementsByEventID defines the SQL statement to
	// remove all placements for an event.
	stmtDeletePlacementsByEventID = `
DELETE FROM events_placements
WHERE event_id=?
`

	// stmtDeleteUsersByEventID defines the SQL statement to
	// remove all users for an event.
	stmtDeleteUsersByEventID = `
DELETE FROM events_users
WHERE event_id=?
`
)

// NewParams defines the parameters for the New method
// Similar to Event type but does not have ID or
// created/updated timestamps
type NewParams struct {
	Name          string    `json:"name"`
	Location      string    `json:"location"`
	StartDatetime time.Time `json:"start_datetime"`
	ShowStartTime bool      `json:"show_start_time"`
	EndDatetime   time.Time `json:"end_datetime"`
	ShowEndTime   bool      `json:"show_end_time"`
	Type          int       `json:"type"`
	GameID        int       `json:"game_id"`
	Description   string    `json:"description"`
	ReferralURL   string    `json:"referral_url"`
	Status        int       `json:"status"`
	Placements    []int     `json:"placements"`
	Users         []int     `json:"users"`
	CurrentUser   int       `json:"current_user"`
}

// New creates a new event.
func (db *Database) New(params *NewParams) (*Event, error) {
	// Create a new event
	event := &Event{
		Name:          params.Name,
		Location:      params.Location,
		StartDatetime: params.StartDatetime,
		ShowStartTime: params.ShowStartTime,
		EndDatetime:   params.EndDatetime,
		ShowEndTime:   params.ShowEndTime,
		Type:          params.Type,
		GameID:        params.GameID,
		Description:   params.Description,
		ReferralURL:   params.ReferralURL,
		Status:        params.Status,
		CreatedBy:     params.CurrentUser,
		ModifiedBy:    params.CurrentUser,
		CreatedAt:     time.Now(),
		ModifiedAt:    time.Now(),
	}

	// Create variable to hold the result
	var res sql.Result
	var err error

	// Begin database transaction
	tx, err := db.db.Begin()
	if err != nil {
		fmt.Printf("err1: %v\n", err)
		return nil, err
	}

	// Execute the query
	if res, err = tx.Exec(stmtInsert, event.Name, event.Location, event.StartDatetime, event.ShowStartTime, event.EndDatetime, event.ShowEndTime, event.Type, event.GameID, event.Description, event.ReferralURL, event.Status, event.CreatedBy, event.ModifiedBy, event.CreatedAt, event.ModifiedAt); err != nil {
		tx.Rollback()
		fmt.Printf("err2: %v\n", err)
		return nil, err
	}

	// Get last insert ID
	id, err := res.LastInsertId()
	if err != nil {
		fmt.Printf("err3: %v\n", err)
		tx.Rollback()
		return nil, err
	}
	event.ID = int(id)

	// Handle placements
	for _, p := range params.Placements {
		// Create a new placement
		placement := &Placement{
			EventID:   event.ID,
			Placement: p,
		}

		//Execute query
		if res, err = tx.Exec(stmtInsertPlacement, placement.EventID, placement.Placement); err != nil {
			tx.Rollback()
			fmt.Printf("err4: %v\n", err)
			return nil, err
		}

		// Add placement to event
		event.Placements = append(event.Placements, p)
	}

	// Handle users
	for _, u := range params.Users {
		// Create a new user
		user := &User{
			EventID: event.ID,
			UserID:  u,
		}

		//Execute query
		if res, err = tx.Exec(stmtInsertUser, user.EventID, user.UserID); err != nil {
			tx.Rollback()
			fmt.Printf("err5: %v\n", err)
			return nil, err
		}

		// Add user to event
		event.Users = append(event.Users, u)
	}

	// Commit transaction
	err = tx.Commit()
	if err != nil {
		tx.Rollback()
		fmt.Printf("err6: %v\n", err)
		return nil, err
	}

	return event, nil
}

// EventRow defines the event data returned from the database
type EventRow struct {
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
	CreatedBy     int       `json:"created_by"`
	ModifiedBy    int       `json:"modified_by"`
	CreatedAt     time.Time `json:"created_at"`
	ModifiedAt    time.Time `json:"modified_at"`
	Placements    string    `json:"placements"`
	Users         string    `json:"users"`
}

// GetByID retrieves a event by their ID
func (db *Database) GetByID(id int) (*Event, error) {
	// Create a new eventRow
	eventRow := &EventRow{}

	// Execute the query
	err := db.db.QueryRow(stmtSelectByID, id).Scan(&eventRow.ID, &eventRow.Name, &eventRow.Location, &eventRow.StartDatetime, &eventRow.ShowStartTime, &eventRow.EndDatetime, &eventRow.ShowEndTime, &eventRow.Type, &eventRow.GameID, &eventRow.Description, &eventRow.ReferralURL, &eventRow.Status, &eventRow.CreatedBy, &eventRow.ModifiedBy, &eventRow.CreatedAt, &eventRow.ModifiedAt, &eventRow.Placements, &eventRow.Users)
	switch {
	case err == sql.ErrNoRows:
		return nil, ErrEventNotFound
	case err != nil:
		return nil, err
	}

	// Create new event
	event := &Event{
		ID:            eventRow.ID,
		Name:          eventRow.Name,
		Location:      eventRow.Location,
		StartDatetime: eventRow.StartDatetime,
		ShowStartTime: eventRow.ShowStartTime,
		EndDatetime:   eventRow.EndDatetime,
		ShowEndTime:   eventRow.ShowEndTime,
		Type:          eventRow.Type,
		GameID:        eventRow.GameID,
		Description:   eventRow.Description,
		ReferralURL:   eventRow.ReferralURL,
		Status:        eventRow.Status,
		CreatedBy:     eventRow.CreatedBy,
		ModifiedBy:    eventRow.ModifiedBy,
		CreatedAt:     eventRow.CreatedAt,
		ModifiedAt:    eventRow.ModifiedAt,
	}

	// Handle placements
	event.Placements = []int{}
	xp := strings.Split(eventRow.Placements, ",")
	for _, p := range xp {
		if p != "" {
			placement, err := strconv.Atoi(p)
			if err != nil {
				return nil, err
			}
			event.Placements = append(event.Placements, placement)
		}
	}

	// Handle users
	event.Users = []int{}
	xu := strings.Split(eventRow.Users, ",")
	for _, u := range xu {
		if u != "" {
			userID, err := strconv.Atoi(u)
			if err != nil {
				return nil, err
			}
			event.Users = append(event.Users, userID)
		}
	}

	return event, nil
}

// GetParams defines the parameters for the Get method.
type GetParams struct {
	ID            *int       `json:"id"`
	Name          *string    `json:"name"`
	Location      *string    `json:"location"`
	StartDatetime *time.Time `json:"start_datetime"`
	ShowStartTime *bool      `json:"show_start_time"`
	EndDatetime   *time.Time `json:"end_datetime"`
	ShowEndTime   *bool      `json:"show_end_time"`
	Type          *int       `json:"type"`
	GameID        *int       `json:"game_id"`
	Description   *string    `json:"description"`
	ReferralURL   *string    `json:"referralURL"`
	Status        *int       `json:"status"`
	CreatedBy     *int       `json:"created_by"`
	ModifiedBy    *int       `json:"modified_by"`
	CreatedAt     *time.Time `json:"created_at"`
	ModifiedAt    *time.Time `json:"modified_at"`
	Placements    *[]int     `json:"placements"`
	Users         *[]int     `json:"users"`
	Offset        int        `json:"offset"`
	Limit         int        `json:"limit"`
}

// Get retrieves a set of events
func (db *Database) Get(params *GetParams) (*Events, error) {
	// Create variables to hold the query fields
	// being filtered on and their values
	var queryFields string
	var queryValues []interface{}

	// Handle the id field
	if params.ID != nil {
		if queryFields == "" {
			queryFields = "WHERE id=?"
		} else {
			queryFields = "AND id=?"
		}

		queryValues = append(queryValues, *params.ID)
	}

	// Handle created_at field
	if params.CreatedAt != nil {
		if queryFields == "" {
			queryFields = "WHERE created_at=?"
		} else {
			queryFields += " AND created_at=?"
		}

		queryValues = append(queryValues, *params.CreatedAt)
	}

	// Handle status field
	if params.Status != nil {
		if queryFields == "" {
			queryFields = "WHERE status=?"
		} else {
			queryFields += " AND status=?"
		}

		queryValues = append(queryValues, *params.Status)
	}

	// TODO: Handle all existing query fields

	query := fmt.Sprintf(stmtSelect, queryFields, params.Offset, params.Limit)

	// Create new events
	events := &Events{
		Events: []*Event{},
	}

	// Execute the query and receive rows of events
	eventRows, err := db.db.Query(query, queryValues...)
	if err != nil {
		return nil, err
	}
	defer eventRows.Close()

	for eventRows.Next() {
		// Create a new event
		eventRow := &EventRow{}

		// Place data from event row into event
		if err := eventRows.Scan(&eventRow.ID, &eventRow.Name, &eventRow.Location, &eventRow.StartDatetime, &eventRow.ShowStartTime, &eventRow.EndDatetime, &eventRow.ShowEndTime, &eventRow.Type, &eventRow.GameID, &eventRow.Description, &eventRow.ReferralURL, &eventRow.Status, &eventRow.CreatedBy, &eventRow.ModifiedBy, &eventRow.CreatedAt, &eventRow.ModifiedAt, &eventRow.Placements, &eventRow.Users); err != nil {
			return nil, err
		}

		// Create new event
		event := &Event{
			ID:            eventRow.ID,
			Name:          eventRow.Name,
			Location:      eventRow.Location,
			StartDatetime: eventRow.StartDatetime,
			ShowStartTime: eventRow.ShowStartTime,
			EndDatetime:   eventRow.EndDatetime,
			ShowEndTime:   eventRow.ShowEndTime,
			Type:          eventRow.Type,
			GameID:        eventRow.GameID,
			Description:   eventRow.Description,
			ReferralURL:   eventRow.ReferralURL,
			Status:        eventRow.Status,
			CreatedBy:     eventRow.CreatedBy,
			ModifiedBy:    eventRow.ModifiedBy,
			CreatedAt:     eventRow.CreatedAt,
			ModifiedAt:    eventRow.ModifiedAt,
		}

		// Handle placements
		event.Placements = []int{}
		xp := strings.Split(eventRow.Placements, ",")
		for _, p := range xp {
			if p != "" {
				placement, err := strconv.Atoi(p)
				if err != nil {
					return nil, err
				}
				event.Placements = append(event.Placements, placement)
			}
		}

		// Handle users
		event.Users = []int{}
		xu := strings.Split(eventRow.Users, ",")
		for _, u := range xu {
			if u != "" {
				userID, err := strconv.Atoi(u)
				if err != nil {
					return nil, err
				}
				event.Users = append(event.Users, userID)
			}
		}

		// Add event to events
		events.Events = append(events.Events, event)
	}

	if err = eventRows.Err(); err != nil {
		return nil, err
	}

	// Build the total count query
	queryCount := fmt.Sprintf(stmtSelectCount, queryFields)

	// Get total count
	var total int
	if err = db.db.QueryRow(queryCount, queryValues...).Scan(&total); err != nil {
		return nil, err
	}
	events.Total = total

	return events, nil
}

// UpdateParams defines the parameters for the Update method.
type UpdateParams struct {
	Name          *string    `json:"name"`
	Location      *string    `json:"location"`
	StartDatetime *time.Time `json:"start_datetime"`
	ShowStartTime *bool      `json:"show_start_time"`
	EndDatetime   *time.Time `json:"end_datetime"`
	ShowEndTime   *bool      `json:"show_end_time"`
	Type          *int       `json:"type"`
	GameID        *int       `json:"game_id"`
	Description   *string    `json:"description"`
	ReferralURL   *string    `json:"referralURL"`
	Status        *int       `json:"status"`
	Placements    *[]int     `json:"placements"`
	Users         *[]int     `json:"users"`
	CurrentUser   *int       `json:"current_user"`
}

// Update updates an event
func (db *Database) Update(id int, params *UpdateParams) (*Event, error) {
	// Create variables to hold the query fields
	// being updated and their new values
	var queryFields string
	var queryValues []interface{}

	// Begin database transaction
	tx, err := db.db.Begin()
	if err != nil {
		fmt.Printf("err %v\n", err)
		return nil, err
	}

	// Handle name field.
	if params.Name != nil {
		if queryFields == "" {
			queryFields = "name=?"
		} else {
			queryFields += ", name=?"
		}

		queryValues = append(queryValues, *params.Name)
	}

	// Handle location field.
	if params.Location != nil {
		if queryFields == "" {
			queryFields = "location=?"
		} else {
			queryFields += ", location=?"
		}

		queryValues = append(queryValues, *params.Location)
	}

	// Handle start_datetime field.
	if params.StartDatetime != nil {
		if queryFields == "" {
			queryFields = "start_datetime=?"
		} else {
			queryFields += ", start_datetime=?"
		}

		queryValues = append(queryValues, *params.StartDatetime)
	}

	// Handle show_start_time field.
	if params.ShowStartTime != nil {
		if queryFields == "" {
			queryFields = "show_start_time=?"
		} else {
			queryFields += ", show_start_time=?"
		}

		queryValues = append(queryValues, *params.ShowStartTime)
	}

	// Handle end_datetime field.
	if params.EndDatetime != nil {
		if queryFields == "" {
			queryFields = "end_datetime=?"
		} else {
			queryFields += ", end_datetime=?"
		}

		queryValues = append(queryValues, *params.EndDatetime)
	}

	// Handle show_end_time field.
	if params.ShowEndTime != nil {
		if queryFields == "" {
			queryFields = "show_end_time=?"
		} else {
			queryFields += ", show_end_time=?"
		}

		queryValues = append(queryValues, *params.ShowEndTime)
	}
	// Handle type field.
	if params.Type != nil {
		if queryFields == "" {
			queryFields = "type=?"
		} else {
			queryFields += ", type=?"
		}

		queryValues = append(queryValues, *params.Type)
	}
	// Handle game_id field.
	if params.GameID != nil {
		if queryFields == "" {
			queryFields = "game_id=?"
		} else {
			queryFields += ", game_id=?"
		}

		queryValues = append(queryValues, *params.GameID)
	}
	// Handle description field.
	if params.Description != nil {
		if queryFields == "" {
			queryFields = "description=?"
		} else {
			queryFields += ", description=?"
		}

		queryValues = append(queryValues, *params.Description)
	}
	// Handle referral_url field.
	if params.ReferralURL != nil {
		if queryFields == "" {
			queryFields = "referral_url=?"
		} else {
			queryFields += ", referral_url=?"
		}

		queryValues = append(queryValues, *params.ReferralURL)
	}
	// Handle status field.
	if params.Status != nil {
		if queryFields == "" {
			queryFields = "status=?"
		} else {
			queryFields += ", status=?"
		}

		queryValues = append(queryValues, *params.Status)
	}

	// Check if the query is empty
	if queryFields == "" {
		return db.GetByID(id)
	} else {
		// Handle modified_by
		queryFields += ", modified_by=?"
		queryValues = append(queryValues, *params.CurrentUser)
		// Handle modified_at
		queryFields += ", modified_at=?"
		queryValues = append(queryValues, time.Now())
	}

	// Build the full query.
	query := fmt.Sprintf(stmtUpdate, queryFields)
	queryValues = append(queryValues, id)

	// Execute the query
	_, err = tx.Exec(query, queryValues...)
	if err != nil {
		fmt.Printf("err %v\n", err)
		tx.Rollback()
		return nil, err
	}

	// Handle placements
	if params.Placements != nil {
		// Delete existing placements
		_, err = tx.Exec(stmtDeletePlacementsByEventID, id)
		if err != nil {
			fmt.Printf("err %v\n", err)
			tx.Rollback()
			return nil, err
		}

		// Add new placements
		for _, p := range *(*params).Placements {
			// Create a new placement
			placement := &Placement{
				EventID:   id,
				Placement: p,
			}

			//Execute query
			if _, err := tx.Exec(stmtInsertPlacement, placement.EventID, placement.Placement); err != nil {
				fmt.Printf("err %v\n", err)
				tx.Rollback()
				return nil, err
			}
		}
	}

	// Handle users
	if params.Users != nil {
		// Delete existing placements
		_, err = tx.Exec(stmtDeleteUsersByEventID, id)
		if err != nil {
			fmt.Printf("err %v\n", err)
			tx.Rollback()
			return nil, err
		}

		// Add new placements
		for _, u := range *(*params).Users {
			// Create a new placement
			placement := &User{
				EventID: id,
				UserID:  u,
			}

			//Execute query
			if _, err := tx.Exec(stmtInsertUser, placement.EventID, placement.UserID); err != nil {
				fmt.Printf("err %v\n", err)
				tx.Rollback()
				return nil, err
			}
		}
	}

	// Commit transaction
	err = tx.Commit()
	if err != nil {
		fmt.Printf("err %v\n", err)
		return nil, err
	}

	// Since the GetByID method is straight forward,
	// we can use this method to retrieve the updated
	// event. Anything more complicated should use the
	// original statement constants.
	return db.GetByID(id)
}

// Delete updates an event's status to 'deleted'
func (db *Database) Delete(id int) error {
	// Begin database transaction
	tx, err := db.db.Begin()
	if err != nil {
		return err
	}

	// Execute the query
	_, err = tx.Exec(stmtDelete, id)
	if err != nil {
		tx.Rollback()
		return err
	}

	// Delete existing placements
	// _, err = tx.Exec(stmtDeletePlacementsByEventID, id)
	// if err != nil {
	// 	tx.Rollback()
	// 	return err
	// }

	// Delete existing users
	// _, err = tx.Exec(stmtDeleteUsersByEventID, id)
	// if err != nil {
	// 	tx.Rollback()
	// 	return err
	// }

	// Commit transaction
	err = tx.Commit()
	if err != nil {
		return err
	}

	return nil
}
