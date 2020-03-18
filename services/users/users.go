package users

import (
	"vil/database"
	dbusers "vil/database/users"
	"vil/services/errors"

	"golang.org/x/crypto/bcrypt"
)

// Service defines the users service
type Service struct {
	db *database.Database
}

// New returns a new users service
func New(db *database.Database) *Service {
	return &Service{
		db: db,
	}
}

// Users defines a user
type User dbusers.User

// Users defines a slice of users
type Users struct {
	Users []*User `json:"users"`
	Total int     `json:"total"`
}

// NewParams defines the parameters for the New method.
type NewParams dbusers.NewParams

// New creates a new user
func (s *Service) New(params *NewParams) (*User, error) {
	// Create a new ParamErrors
	pes := errors.NewParamErrors()

	// Check email
	if params.Email == "" {
		pes.Add(errors.NewParamError("email", ErrEmailEmpty))
	} else {
		_, err := s.db.Users.GetByEmail(params.Email)
		if err == nil {
			pes.Add(errors.NewParamError("email", ErrEmailExists))
		} else if err != nil && err != dbusers.ErrUserNotFound {
			return nil, err
		}
	}

	// Check password for validity
	// TODO: Increase quality of password check
	if len(params.Password) < 8 {
		pes.Add(errors.NewParamError("password", ErrPassword))
	}

	// Handle birth date
	if len(params.BirthDate) != 10 {
		params.BirthDate = "0000-00-00"
	}

	// Return if there were parameter errors
	if pes.Length() > 0 {
		return nil, pes
	}

	// Hash the password
	pwhash, err := bcrypt.GenerateFromPassword([]byte(params.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}

	// Create this user in the database
	dbu, err := s.db.Users.New(&dbusers.NewParams{
		Username:      params.Username,
		Password:      string(pwhash),
		Email:         params.Email,
		FName:         params.FName,
		MName:         params.MName,
		LName:         params.LName,
		Address:       params.Address,
		City:          params.City,
		Province:      params.Province,
		Zip:           params.Zip,
		Country:       params.Country,
		BirthDate:     params.BirthDate,
		Description:   params.Description,
		Role:          params.Role,
		PrivilegeID:   params.PrivilegeID,
		Status:        params.Status,
		FacebookURL:   params.FacebookURL,
		TwitterURL:    params.TwitterURL,
		InstagramURL:  params.InstagramURL,
		TwitchURL:     params.TwitchURL,
		YoutubeURL:    params.YoutubeURL,
		OtherURL:      params.OtherURL,
		PS4Gamertag:   params.PS4Gamertag,
		XBoxGamertag:  params.XBoxGamertag,
		SteamGamertag: params.SteamGamertag,
	})
	if err != nil {
		return nil, err
	}

	// Create a new user
	user := &User{
		ID:            dbu.ID,
		Username:      dbu.Username,
		Password:      dbu.Password,
		Email:         dbu.Email,
		FName:         dbu.FName,
		MName:         dbu.MName,
		LName:         dbu.LName,
		Address:       dbu.Address,
		City:          dbu.City,
		Province:      dbu.Province,
		Zip:           dbu.Zip,
		Country:       dbu.Country,
		BirthDate:     dbu.BirthDate,
		Description:   dbu.Description,
		Role:          dbu.Role,
		PrivilegeID:   dbu.PrivilegeID,
		Status:        dbu.Status,
		FacebookURL:   dbu.FacebookURL,
		TwitterURL:    dbu.TwitterURL,
		InstagramURL:  dbu.InstagramURL,
		TwitchURL:     dbu.TwitchURL,
		YoutubeURL:    dbu.YoutubeURL,
		OtherURL:      dbu.OtherURL,
		PS4Gamertag:   dbu.PS4Gamertag,
		XBoxGamertag:  dbu.XBoxGamertag,
		SteamGamertag: dbu.SteamGamertag,
		CreatedAt:     dbu.CreatedAt,
		UpdatedAt:     dbu.UpdatedAt,
	}

	return user, nil
}

// GetParams defines the parameters for the Get method
type GetParams dbusers.GetParams

func (s *Service) Get(params *GetParams) (*Users, error) {
	// Try to pull the users from the database
	dbus, err := s.db.Users.Get(&dbusers.GetParams{
		ID:            params.ID,
		Username:      params.Username,
		Email:         params.Email,
		FName:         params.FName,
		MName:         params.MName,
		LName:         params.LName,
		Title:         params.Title,
		Address:       params.Address,
		City:          params.City,
		Province:      params.Province,
		Zip:           params.Zip,
		Country:       params.Country,
		BirthDate:     params.BirthDate,
		Description:   params.Description,
		Role:          params.Role,
		PrivilegeID:   params.PrivilegeID,
		Status:        params.Status,
		FacebookURL:   params.FacebookURL,
		TwitterURL:    params.TwitterURL,
		InstagramURL:  params.InstagramURL,
		TwitchURL:     params.TwitchURL,
		YoutubeURL:    params.YoutubeURL,
		OtherURL:      params.OtherURL,
		PS4Gamertag:   params.PS4Gamertag,
		XBoxGamertag:  params.XBoxGamertag,
		SteamGamertag: params.SteamGamertag,
		CreatedAt:     params.CreatedAt,
		UpdatedAt:     params.UpdatedAt,
		Offset:        params.Offset,
		Limit:         params.Limit,
	})
	if err != nil {
		return nil, err
	}

	// Create a new Users
	users := &Users{
		Users: []*User{},
		Total: dbus.Total,
	}

	// Move user rows into users
	for _, dbu := range dbus.Users {
		user := &User{
			ID:            dbu.ID,
			Username:      dbu.Username,
			Password:      dbu.Password,
			Email:         dbu.Email,
			FName:         dbu.FName,
			MName:         dbu.MName,
			LName:         dbu.LName,
			Address:       dbu.Address,
			City:          dbu.City,
			Province:      dbu.Province,
			Zip:           dbu.Zip,
			Country:       dbu.Country,
			BirthDate:     dbu.BirthDate,
			Description:   dbu.Description,
			Role:          dbu.Role,
			PrivilegeID:   dbu.PrivilegeID,
			Status:        dbu.Status,
			FacebookURL:   dbu.FacebookURL,
			TwitterURL:    dbu.TwitterURL,
			InstagramURL:  dbu.InstagramURL,
			TwitchURL:     dbu.TwitchURL,
			YoutubeURL:    dbu.YoutubeURL,
			OtherURL:      dbu.OtherURL,
			PS4Gamertag:   dbu.PS4Gamertag,
			XBoxGamertag:  dbu.XBoxGamertag,
			SteamGamertag: dbu.SteamGamertag,
			CreatedAt:     dbu.CreatedAt,
			UpdatedAt:     dbu.UpdatedAt,
		}

		users.Users = append(users.Users, user)

	}

	return users, nil
}

// GetByID retrieves a user by their ID
func (s *Service) GetByID(id int) (*User, error) {
	// Get user from database
	u, err := s.db.Users.GetByID(id)
	if err != nil {
		return nil, err
	}

	// Create a new user
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

	return user, nil
}

// Update params defines the parameters for the update methods
type UpdateParams dbusers.UpdateParams

// UpdateByID updates a user
func (s *Service) UpdateByID(id int, params *UpdateParams) (*User, error) {
	// Try and pull this user from the database
	dbu, err := s.db.Users.GetByID(id)
	if err == dbusers.ErrUserNotFound {
		return nil, ErrUserNotFound
	} else if err != nil {
		return nil, err
	}

	dbu, err = s.db.Users.Update(id, &dbusers.UpdateParams{
		Username:      params.Username,
		Password:      params.Password,
		FName:         params.FName,
		MName:         params.MName,
		LName:         params.LName,
		Title:         params.Title,
		Address:       params.Address,
		City:          params.City,
		Province:      params.Province,
		Zip:           params.Zip,
		Country:       params.Country,
		BirthDate:     params.BirthDate,
		Description:   params.Description,
		Role:          params.Role,
		PrivilegeID:   params.PrivilegeID,
		Status:        params.Status,
		FacebookURL:   params.FacebookURL,
		TwitterURL:    params.TwitterURL,
		InstagramURL:  params.InstagramURL,
		TwitchURL:     params.TwitchURL,
		YoutubeURL:    params.YoutubeURL,
		OtherURL:      params.OtherURL,
		PS4Gamertag:   params.PS4Gamertag,
		XBoxGamertag:  params.XBoxGamertag,
		SteamGamertag: params.SteamGamertag,
		UpdatedAt:     params.UpdatedAt,
	})
	if err != nil {
		return nil, err
	}

	user := &User{
		ID:            dbu.ID,
		Username:      dbu.Username,
		Password:      dbu.Password,
		Email:         dbu.Email,
		FName:         dbu.FName,
		MName:         dbu.MName,
		LName:         dbu.LName,
		Title:         dbu.Title,
		Address:       dbu.Address,
		City:          dbu.City,
		Province:      dbu.Province,
		Zip:           dbu.Zip,
		Country:       dbu.Country,
		BirthDate:     dbu.BirthDate,
		Description:   dbu.Description,
		Role:          dbu.Role,
		PrivilegeID:   dbu.PrivilegeID,
		Status:        dbu.Status,
		FacebookURL:   dbu.FacebookURL,
		TwitterURL:    dbu.TwitterURL,
		InstagramURL:  dbu.InstagramURL,
		TwitchURL:     dbu.TwitchURL,
		YoutubeURL:    dbu.YoutubeURL,
		OtherURL:      dbu.OtherURL,
		PS4Gamertag:   dbu.PS4Gamertag,
		XBoxGamertag:  dbu.XBoxGamertag,
		SteamGamertag: dbu.SteamGamertag,
		CreatedAt:     dbu.CreatedAt,
		UpdatedAt:     dbu.UpdatedAt,
	}

	return user, nil
}

// DeleteByID deletes a user
func (s *Service) DeleteByID(id int) error {
	// Try and pull this user from the database
	_, err := s.db.Users.GetByID(id)
	if err == dbusers.ErrUserNotFound {
		return ErrUserNotFound
	} else if err != nil {
		return err
	}

	err = s.db.Users.Delete(id)
	if err != nil {
		return err
	}

	return nil
}

// LoginParams defines the parameters for the Login method
type LoginParams struct {
	Email    string
	Password string
}

// Login checks if a user exists in the database and can log in.
func (s *Service) Login(params *LoginParams) (*User, error) {

	dbu, err := s.db.Users.GetByEmail(params.Email)
	if err == dbusers.ErrUserNotFound {
		return nil, ErrInvalidLogin
	} else if err != nil {
		return nil, err
	}

	// Validate the password
	if err = bcrypt.CompareHashAndPassword([]byte(dbu.Password), []byte(params.Password)); err != nil {
		return nil, ErrInvalidLogin
	}

	// Create a new user
	user := &User{
		ID:            dbu.ID,
		Username:      dbu.Username,
		Password:      dbu.Password,
		Email:         dbu.Email,
		FName:         dbu.FName,
		MName:         dbu.MName,
		LName:         dbu.LName,
		Address:       dbu.Address,
		City:          dbu.City,
		Province:      dbu.Province,
		Zip:           dbu.Zip,
		Country:       dbu.Country,
		BirthDate:     dbu.BirthDate,
		Description:   dbu.Description,
		Role:          dbu.Role,
		PrivilegeID:   dbu.PrivilegeID,
		Status:        dbu.Status,
		FacebookURL:   dbu.FacebookURL,
		TwitterURL:    dbu.TwitterURL,
		InstagramURL:  dbu.InstagramURL,
		TwitchURL:     dbu.TwitchURL,
		YoutubeURL:    dbu.YoutubeURL,
		OtherURL:      dbu.OtherURL,
		PS4Gamertag:   dbu.PS4Gamertag,
		XBoxGamertag:  dbu.XBoxGamertag,
		SteamGamertag: dbu.SteamGamertag,
		CreatedAt:     dbu.CreatedAt,
		UpdatedAt:     dbu.UpdatedAt,
	}

	return user, nil
}
