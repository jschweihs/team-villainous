package events

import (
	"vil/database"
	dbevents "vil/database/events"
	"vil/services/errors"
)

// Service defines the events service
type Service struct {
	db *database.Database
}

// New returns a new events service
func New(db *database.Database) *Service {
	return &Service{
		db: db,
	}
}

// Event defines an event
type Event dbevents.Event

// Placement defines a placement
type Placement dbevents.Placement

// UserID defines a user
type User dbevents.User

// Events defines a slice of events
type Events struct {
	Events []*Event `json:"events"`
	Total  int      `json:"total"`
}

// NewParams defines the parameters for the New method.
type NewParams dbevents.NewParams

// New creates a new event
func (s *Service) New(params *NewParams) (*Event, error) {
	// Create a new ParamErrors
	pes := errors.NewParamErrors()

	// Validate name
	if params.Name == "" {
		pes.Add(errors.NewParamError("name", ErrNameEmpty))
	}

	// TODO: Validate other params

	// Return if there were parameter errors
	if pes.Length() > 0 {
		return nil, pes
	}

	// Create this event in the database
	dbe, err := s.db.Events.New(&dbevents.NewParams{
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
		Placements:    params.Placements,
		Users:         params.Users,
		CurrentUser:   params.CurrentUser,
	})
	if err != nil {
		return nil, err
	}

	// Create a new event
	event := &Event{
		ID:            dbe.ID,
		Name:          dbe.Name,
		Location:      dbe.Location,
		StartDatetime: dbe.StartDatetime,
		ShowStartTime: dbe.ShowStartTime,
		Type:          dbe.Type,
		GameID:        dbe.GameID,
		Description:   dbe.Description,
		ReferralURL:   dbe.ReferralURL,
		Status:        dbe.Status,
		Placements:    dbe.Placements,
		Users:         dbe.Users,
		CreatedBy:     dbe.CreatedBy,
		ModifiedBy:    dbe.ModifiedBy,
		CreatedAt:     dbe.CreatedAt,
		ModifiedAt:    dbe.ModifiedAt,
	}

	return event, nil
}

// GetParams defines the parameters for the Get method
type GetParams dbevents.GetParams

func (s *Service) Get(params *GetParams) (*Events, error) {
	// Try to pull the events from the database
	dbes, err := s.db.Events.Get(&dbevents.GetParams{
		ID:            params.ID,
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
		CreatedBy:     params.CreatedBy,
		ModifiedBy:    params.ModifiedBy,
		CreatedAt:     params.CreatedAt,
		ModifiedAt:    params.ModifiedAt,
		Offset:        params.Offset,
		Limit:         params.Limit,
	})
	if err != nil {
		return nil, err
	}

	// Create a new Events
	events := &Events{
		Events: []*Event{},
		Total:  dbes.Total,
	}

	// Move event rows into events
	for _, dbe := range dbes.Events {
		event := &Event{
			ID:            dbe.ID,
			Name:          dbe.Name,
			Location:      dbe.Location,
			StartDatetime: dbe.StartDatetime,
			ShowStartTime: dbe.ShowStartTime,
			Type:          dbe.Type,
			GameID:        dbe.GameID,
			Description:   dbe.Description,
			ReferralURL:   dbe.ReferralURL,
			Status:        dbe.Status,
			Placements:    dbe.Placements,
			Users:         dbe.Users,
			CreatedBy:     dbe.CreatedBy,
			ModifiedBy:    dbe.ModifiedBy,
			CreatedAt:     dbe.CreatedAt,
			ModifiedAt:    dbe.ModifiedAt,
		}

		events.Events = append(events.Events, event)

	}

	return events, nil
}

// GetByID retrieves an event by their ID
func (s *Service) GetByID(id int) (*Event, error) {
	// Get event from database
	e, err := s.db.Events.GetByID(id)
	if err != nil {
		return nil, err
	}

	// Create a new event
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

	return event, nil
}

// Update params defines the parameters for the update methods
type UpdateParams dbevents.UpdateParams

// UpdateByID updates an event
func (s *Service) UpdateByID(id int, params *UpdateParams) (*Event, error) {
	// Try and pull this event from the database
	dbe, err := s.db.Events.GetByID(id)
	if err == dbevents.ErrEventNotFound {
		return nil, ErrEventNotFound
	} else if err != nil {
		return nil, err
	}

	dbe, err = s.db.Events.Update(id, &dbevents.UpdateParams{
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
		Placements:    params.Placements,
		Users:         params.Users,
		CurrentUser:   params.CurrentUser,
	})
	if err != nil {
		return nil, err
	}

	event := &Event{
		ID:            dbe.ID,
		Name:          dbe.Name,
		Location:      dbe.Location,
		StartDatetime: dbe.StartDatetime,
		ShowStartTime: dbe.ShowStartTime,
		Type:          dbe.Type,
		GameID:        dbe.GameID,
		Description:   dbe.Description,
		ReferralURL:   dbe.ReferralURL,
		Status:        dbe.Status,
		Placements:    dbe.Placements,
		Users:         dbe.Users,
		CreatedBy:     dbe.CreatedBy,
		ModifiedBy:    dbe.ModifiedBy,
		CreatedAt:     dbe.CreatedAt,
		ModifiedAt:    dbe.ModifiedAt,
	}

	return event, nil
}

// DeleteByID deletes an event
func (s *Service) DeleteByID(id int) error {
	// Try and pull this event from the database
	_, err := s.db.Events.GetByID(id)
	if err == dbevents.ErrEventNotFound {
		return ErrEventNotFound
	} else if err != nil {
		return err
	}

	err = s.db.Events.Delete(id)
	if err != nil {
		return err
	}

	return nil
}
