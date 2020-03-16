package roles

import (
	"vil/database"
	dbroles "vil/database/roles"
	"vil/services/errors"
)

// Service defines the roles service
type Service struct {
	db *database.Database
}

// New returns a new roles service
func New(db *database.Database) *Service {
	return &Service{
		db: db,
	}
}

// Roles defines a role
type Role dbroles.Role

// Roles defines a slice of roles
type Roles struct {
	Roles []*Role `json:"roles"`
	Total int     `json:"total"`
}

// NewParams defines the parameters for the New method.
type NewParams dbroles.NewParams

// New creates a new role
func (s *Service) New(params *NewParams) (*Role, error) {
	// Create a new ParamErrors
	pes := errors.NewParamErrors()

	// Check role name
	if params.Name == "" {
		pes.Add(errors.NewParamError("name", ErrRoleEmpty))
	} else {
		_, err := s.db.Roles.GetByName(params.Name)
		if err == nil {
			pes.Add(errors.NewParamError("name", ErrRoleExists))
		} else if err != nil && err != dbroles.ErrRoleNotFound {
			return nil, err
		}
	}

	// Return if there were parameter errors
	if pes.Length() > 0 {
		return nil, pes
	}

	// Create this role in the database
	dbr, err := s.db.Roles.New(&dbroles.NewParams{
		Name: params.Name,
	})
	if err != nil {
		return nil, err
	}

	// Create a new role
	role := &Role{
		ID:   dbr.ID,
		Name: dbr.Name,
	}

	return role, nil
}

// GetParams defines the parameters for the Get method
type GetParams dbroles.GetParams

func (s *Service) Get(params *GetParams) (*Roles, error) {
	// Try to pull the roles from the database
	dbrs, err := s.db.Roles.Get(&dbroles.GetParams{
		ID:     params.ID,
		Name:   params.Name,
		Offset: params.Offset,
		Limit:  params.Limit,
	})
	if err != nil {
		return nil, err
	}

	// Create a new Roles
	roles := &Roles{
		Roles: []*Role{},
		Total: dbrs.Total,
	}

	// Move role rows into roles
	for _, dbr := range dbrs.Roles {
		role := &Role{
			ID:   dbr.ID,
			Name: dbr.Name,
		}

		roles.Roles = append(roles.Roles, role)

	}

	return roles, nil
}

// GetByID retrieves a role by their ID
func (s *Service) GetByID(id int) (*Role, error) {
	// Get role from database
	r, err := s.db.Roles.GetByID(id)
	if err != nil {
		return nil, err
	}

	// Create a new role
	role := &Role{
		ID:   r.ID,
		Name: r.Name,
	}

	return role, nil
}

// Update params defines the parameters for the update methods
type UpdateParams dbroles.UpdateParams

// UpdateByID updates a role
func (s *Service) UpdateByID(id int, params *UpdateParams) (*Role, error) {
	// Try and pull this role from the database
	dbr, err := s.db.Roles.GetByID(id)
	if err == dbroles.ErrRoleNotFound {
		return nil, ErrRoleNotFound
	} else if err != nil {
		return nil, err
	}

	dbr, err = s.db.Roles.Update(id, &dbroles.UpdateParams{
		Name: params.Name,
	})
	if err != nil {
		return nil, err
	}

	role := &Role{
		ID:   dbr.ID,
		Name: dbr.Name,
	}

	return role, nil
}

// DeleteByID deletes a role
func (s *Service) DeleteByID(id int) error {
	// Try and pull this role from the database
	_, err := s.db.Roles.GetByID(id)
	if err == dbroles.ErrRoleNotFound {
		return ErrRoleNotFound
	} else if err != nil {
		return err
	}

	err = s.db.Roles.Delete(id)
	if err != nil {
		return err
	}

	return nil
}
