package users

import (
	"database/sql"
	"fmt"
	"time"
)

// Database defines the users database
type Database struct {
	db *sql.DB
}

// New creates a new users database
func New(db *sql.DB) *Database {
	return &Database{
		db: db,
	}
}

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

// Users defines a set of users
type Users struct {
	Users []*User `json:"users"`
	Total int     `json:"total"`
}

const (
	// stmtInsert defines the SQL statement to
	// insert a new user into the database
	stmtInsert = `
INSERT INTO users (
	username, 
	password, 
	email, 
	f_name, 
	m_name, 
	l_name, 
	title, 
	address, 
	city, 
	province, 
	zip, 
	country, 
	birth_date, 
	description, 
	role, 
	privilege_id, 
	status, 
	facebook_url, 
	twitter_url, 
	instagram_url, 
	twitch_url, 
	youtube_url,
	other_url, 
	ps4_gamertag, 
	xbox_gamertag, 
	steam_gamertag, 
	created_at, 
	updated_at
)
VALUES (
	?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?
)
`

	// stmtSelect defines the SQL statement to
	// select a set of users.
	stmtSelect = `
SELECT 
	id,
	username, 
	password, 
	email, 
	f_name, 
	m_name, 
	l_name, 
	title, 
	address, 
	city, 
	province, 
	zip, 
	country, 
	birth_date, 
	description, 
	role, 
	privilege_id, 
	status, 
	facebook_url, 
	twitter_url, 
	instagram_url, 
	twitch_url, 
	youtube_url,
	other_url, 
	ps4_gamertag, 
	xbox_gamertag, 
	steam_gamertag, 
	created_at, 
	updated_at
FROM users
%s
LIMIT %v, %v
`

	// stmtSelectCount defines the SQL statement to
	// select the total number of users according to the filters.
	stmtSelectCount = `
SELECT COUNT(*)
FROM users
%s
`

	// stmtSelectByID defines the SQL statement to
	// select a user by their ID.
	stmtSelectByID = `
SELECT 
	id,
	username, 
	password, 
	email, 
	f_name, 
	m_name, 
	l_name, 
	title, 
	address, 
	city, 
	province, 
	zip, 
	country, 
	birth_date, 
	description, 
	role, 
	privilege_id, 
	status, 
	facebook_url, 
	twitter_url, 
	instagram_url, 
	twitch_url, 
	youtube_url,
	other_url, 
	ps4_gamertag, 
	xbox_gamertag, 
	steam_gamertag, 
	created_at, 
	updated_at
FROM users
WHERE id=?
`

	// stmtSelectByEmail defines the SQL statement to
	// select a user by their email.
	stmtSelectByEmail = `
SELECT 
	id,
	username, 
	password, 
	email, 
	f_name, 
	m_name, 
	l_name, 
	title, 
	address, 
	city, 
	province, 
	zip, 
	country, 
	birth_date, 
	description, 
	role, 
	privilege_id, 
	status, 
	facebook_url, 
	twitter_url, 
	instagram_url, 
	twitch_url, 
	youtube_url,
	other_url, 
	ps4_gamertag, 
	xbox_gamertag, 
	steam_gamertag, 
	created_at, 
	updated_at
FROM users
WHERE email=?
`

	// stmtUpdate defines the SQL statement to
	// update a user.
	stmtUpdate = `
UPDATE users
SET %s
WHERE id=?
`

	// stmtDelete defines the SQL statement to
	// remove a user.
	stmtDelete = `
UPDATE users
SET status=2
WHERE id=?
`
)

// NewParams defines the parameters for the New method
// Similar to User type but does not have ID or
// created/updated timestamps
type NewParams struct {
	Username      string `json:"username"`
	Password      string `json:"password"`
	Email         string `json:"email"`
	FName         string `json:"f_name"`
	MName         string `json:"m_name"`
	LName         string `json:"l_name"`
	Title         string `json:"title"`
	Address       string `json:"address"`
	City          string `json:"city"`
	Province      string `json:"province"`
	Zip           string `json:"zip"`
	Country       string `json:"country"`
	BirthDate     string `json:"birth_date"`
	Description   string `json:"description"`
	Role          int    `json:"role"`
	PrivilegeID   int    `json:"privilege_id"`
	Status        int    `json:"status"`
	FacebookURL   string `json:"facebook_url"`
	TwitterURL    string `json:"twitter_url"`
	InstagramURL  string `json:"instagram_url"`
	TwitchURL     string `json:"twitch_url"`
	YoutubeURL    string `json:"youtube_url"`
	OtherURL      string `json:"other_url"`
	PS4Gamertag   string `json:"ps4_gamertag"`
	XBoxGamertag  string `json:"xbox_gamertag"`
	SteamGamertag string `json:"steam_gamertag"`
}

// New creates a new user.
func (db *Database) New(params *NewParams) (*User, error) {
	// Create a new user
	user := &User{
		Username:      params.Username,
		Password:      params.Password,
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
		CreatedAt:     time.Now(),
		UpdatedAt:     time.Now(),
	}

	// Create variable to hold the result
	var res sql.Result
	var err error

	// Execute the query
	if res, err = db.db.Exec(stmtInsert, user.Username, user.Password, user.Email, user.FName, user.MName, user.LName, user.Title, user.Address, user.City, user.Province, user.Zip, user.Country, user.BirthDate, user.Description, user.Role, user.PrivilegeID, user.Status, user.FacebookURL, user.TwitterURL, user.InstagramURL, user.TwitchURL, user.YoutubeURL, user.OtherURL, user.PS4Gamertag, user.XBoxGamertag, user.SteamGamertag, user.CreatedAt, user.UpdatedAt); err != nil {
		return nil, err
	}

	// Get last insert ID
	id, err := res.LastInsertId()
	if err != nil {
		return nil, err
	}
	user.ID = int(id)

	return user, nil
}

// GetByID retrieves a user by their ID
func (db *Database) GetByID(id int) (*User, error) {
	// Create a new user
	user := &User{}

	// Execute the query
	err := db.db.QueryRow(stmtSelectByID, id).Scan(&user.ID, &user.Username, &user.Password, &user.Email, &user.FName, &user.MName, &user.LName, &user.Title, &user.Address, &user.City, &user.Province, &user.Zip, &user.Country, &user.BirthDate, &user.Description, &user.Role, &user.PrivilegeID, &user.Status, &user.FacebookURL, &user.TwitterURL, &user.InstagramURL, &user.TwitchURL, &user.YoutubeURL, &user.OtherURL, &user.PS4Gamertag, &user.XBoxGamertag, &user.SteamGamertag, &user.CreatedAt, &user.UpdatedAt)
	switch {
	case err == sql.ErrNoRows:
		return nil, ErrUserNotFound
	case err != nil:
		return nil, err
	}

	return user, nil
}

// GetByEmail retrieves a user by their email
func (db *Database) GetByEmail(email string) (*User, error) {
	// Create a new user
	user := &User{}

	// Execute the query
	err := db.db.QueryRow(stmtSelectByEmail, email).Scan(&user.ID, &user.Username, &user.Password, &user.Email, &user.FName, &user.MName, &user.LName, &user.Title, &user.Address, &user.City, &user.Province, &user.Zip, &user.Country, &user.BirthDate, &user.Description, &user.Role, &user.PrivilegeID, &user.Status, &user.FacebookURL, &user.TwitterURL, &user.InstagramURL, &user.TwitchURL, &user.YoutubeURL, &user.OtherURL, &user.PS4Gamertag, &user.XBoxGamertag, &user.SteamGamertag, &user.CreatedAt, &user.UpdatedAt)
	switch {
	case err == sql.ErrNoRows:
		return nil, ErrUserNotFound
	case err != nil:
		return nil, err
	}

	return user, nil
}

// GetParams defines the parameters for the Get method.
type GetParams struct {
	ID            *int       `json:"id"`
	Username      *string    `json:"username"`
	Email         *string    `json:"email"`
	FName         *string    `json:"f_name"`
	MName         *string    `json:"m_name"`
	LName         *string    `json:"l_name"`
	Title         *string    `json:"title"`
	Address       *string    `json:"address"`
	City          *string    `json:"city"`
	Province      *string    `json:"province"`
	Zip           *string    `json:"zip"`
	Country       *string    `json:"country"`
	BirthDate     *string    `json:"birth_date"`
	Description   *string    `json:"description"`
	Role          *int       `json:"role"`
	PrivilegeID   *int       `json:"privilege_id"`
	Status        *int       `json:"status"`
	FacebookURL   *string    `json:"facebook_url"`
	TwitterURL    *string    `json:"twitter_url"`
	InstagramURL  *string    `json:"instagram_url"`
	TwitchURL     *string    `json:"twitch_url"`
	YoutubeURL    *string    `json:"youtube_url"`
	OtherURL      *string    `json:"other_url"`
	PS4Gamertag   *string    `json:"ps4_gamertag"`
	XBoxGamertag  *string    `json:"xbox_gamertag"`
	SteamGamertag *string    `json:"steam_gamertag"`
	CreatedAt     *time.Time `json:"created_at"`
	UpdatedAt     *time.Time `json:"updated_at"`
	Offset        int        `json:"offset"`
	Limit         int        `json:"limit"`
}

// GetAll retrieves all users
func (db *Database) Get(params *GetParams) (*Users, error) {
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

	// Create new users
	users := &Users{
		Users: []*User{},
	}

	// Execute the query and receive rows of users
	userRows, err := db.db.Query(query, queryValues...)
	if err != nil {
		return nil, err
	}
	defer userRows.Close()

	for userRows.Next() {
		// Create a new user
		user := &User{}

		// Place data from user row into user
		if err := userRows.Scan(&user.ID, &user.Username, &user.Password, &user.Email, &user.FName, &user.MName, &user.LName, &user.Title, &user.Address, &user.City, &user.Province, &user.Zip, &user.Country, &user.BirthDate, &user.Description, &user.Role, &user.PrivilegeID, &user.Status, &user.FacebookURL, &user.TwitterURL, &user.InstagramURL, &user.TwitchURL, &user.YoutubeURL, &user.OtherURL, &user.PS4Gamertag, &user.XBoxGamertag, &user.SteamGamertag, &user.CreatedAt, &user.UpdatedAt); err != nil {
			return nil, err
		}

		// Add user to result users slice
		users.Users = append(users.Users, user)
	}
	if err = userRows.Err(); err != nil {
		return nil, err
	}

	// Build the total count query
	queryCount := fmt.Sprintf(stmtSelectCount, queryFields)

	// Get total count
	var total int
	if err = db.db.QueryRow(queryCount, queryValues...).Scan(&total); err != nil {
		return nil, err
	}
	users.Total = total

	return users, nil
}

// UpdateParams defines the parameters for the Update method.
type UpdateParams struct {
	Username      *string `json:"username"`
	Password      *string `json:"password"`
	FName         *string `json:"f_name"`
	MName         *string `json:"m_name"`
	LName         *string `json:"l_name"`
	Title         *string `json:"title"`
	Address       *string `json:"address"`
	City          *string `json:"city"`
	Province      *string `json:"province"`
	Zip           *string `json:"zip"`
	Country       *string `json:"country"`
	BirthDate     *string `json:"birth_date"`
	Description   *string `json:"description"`
	Role          *int    `json:"role"`
	PrivilegeID   *int    `json:"privilege_id"`
	Status        *int    `json:"status"`
	FacebookURL   *string `json:"facebook_url"`
	TwitterURL    *string `json:"twitter_url"`
	InstagramURL  *string `json:"instagram_url"`
	TwitchURL     *string `json:"twitch_url"`
	YoutubeURL    *string `json:"youtube_url"`
	OtherURL      *string `json:"other_url"`
	PS4Gamertag   *string `json:"ps4_gamertag"`
	XBoxGamertag  *string `json:"xbox_gamertag"`
	SteamGamertag *string `json:"steam_gamertag"`
}

// Update updates a user
func (db *Database) Update(id int, params *UpdateParams) (*User, error) {
	// Create variables to hold the query fields
	// being updated and their new values
	var queryFields string
	var queryValues []interface{}

	// Handle username field.
	if params.Username != nil {
		if queryFields == "" {
			queryFields = "username=?"
		} else {
			queryFields += ", username=?"
		}

		queryValues = append(queryValues, *params.Username)
	}

	// Handle password field.
	if params.Password != nil {
		if queryFields == "" {
			queryFields = "password=?"
		} else {
			queryFields += ", password=?"
		}

		queryValues = append(queryValues, *params.Password)
	}

	// Handle f_name field.
	if params.FName != nil {
		if queryFields == "" {
			queryFields = "f_name=?"
		} else {
			queryFields += ", f_name=?"
		}

		queryValues = append(queryValues, *params.FName)
	}

	// Handle m_name field.
	if params.MName != nil {
		if queryFields == "" {
			queryFields = "m_name=?"
		} else {
			queryFields += ", m_name=?"
		}

		queryValues = append(queryValues, *params.MName)
	}

	// Handle l_name field.
	if params.LName != nil {
		if queryFields == "" {
			queryFields = "l_name=?"
		} else {
			queryFields += ", l_name=?"
		}

		queryValues = append(queryValues, *params.LName)
	}

	// Handle title field.
	if params.Title != nil {
		if queryFields == "" {
			queryFields = "title=?"
		} else {
			queryFields += ", title=?"
		}

		queryValues = append(queryValues, *params.Title)
	}
	// Handle address field.
	if params.Address != nil {
		if queryFields == "" {
			queryFields = "address=?"
		} else {
			queryFields += ", address=?"
		}

		queryValues = append(queryValues, *params.Address)
	}
	// Handle city field.
	if params.City != nil {
		if queryFields == "" {
			queryFields = "city=?"
		} else {
			queryFields += ", city=?"
		}

		queryValues = append(queryValues, *params.City)
	}
	// Handle province field.
	if params.Province != nil {
		if queryFields == "" {
			queryFields = "province=?"
		} else {
			queryFields += ", province=?"
		}

		queryValues = append(queryValues, *params.Province)
	}
	// Handle zip field.
	if params.Zip != nil {
		if queryFields == "" {
			queryFields = "zip=?"
		} else {
			queryFields += ", zip=?"
		}

		queryValues = append(queryValues, *params.Zip)
	}
	// Handle country field.
	if params.Country != nil {
		if queryFields == "" {
			queryFields = "country=?"
		} else {
			queryFields += ", country=?"
		}

		queryValues = append(queryValues, *params.Country)
	}
	// Handle birth_date field.
	if params.BirthDate != nil {
		if queryFields == "" {
			queryFields = "birth_date=?"
		} else {
			queryFields += ", birth_date=?"
		}

		queryValues = append(queryValues, *params.BirthDate)
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
	// Handle role field.
	if params.Role != nil {
		if queryFields == "" {
			queryFields = "role=?"
		} else {
			queryFields += ", role=?"
		}

		queryValues = append(queryValues, *params.Role)
	}
	// Handle privilege_id field.
	if params.PrivilegeID != nil {
		if queryFields == "" {
			queryFields = "privilege_id=?"
		} else {
			queryFields += ", privilege_id=?"
		}

		queryValues = append(queryValues, *params.PrivilegeID)
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
	// Handle facebook_url field.
	if params.FacebookURL != nil {
		if queryFields == "" {
			queryFields = "facebook_url=?"
		} else {
			queryFields += ", facebook_url=?"
		}

		queryValues = append(queryValues, *params.FacebookURL)
	}
	// Handle twitter_url field.
	if params.TwitterURL != nil {
		if queryFields == "" {
			queryFields = "twitter_url=?"
		} else {
			queryFields += ", twitter_url=?"
		}

		queryValues = append(queryValues, *params.TwitterURL)
	}
	// Handle instagram_url field.
	if params.InstagramURL != nil {
		if queryFields == "" {
			queryFields = "instagram_url=?"
		} else {
			queryFields += ", instagram_url=?"
		}

		queryValues = append(queryValues, *params.InstagramURL)
	}
	// Handle twitch_url field.
	if params.TwitchURL != nil {
		if queryFields == "" {
			queryFields = "twitch_url=?"
		} else {
			queryFields += ", twitch_url=?"
		}

		queryValues = append(queryValues, *params.TwitchURL)
	}
	// Handle youtube_url field.
	if params.YoutubeURL != nil {
		if queryFields == "" {
			queryFields = "youtube_url=?"
		} else {
			queryFields += ", youtube_url=?"
		}

		queryValues = append(queryValues, *params.YoutubeURL)
	}
	// Handle other_url field.
	if params.OtherURL != nil {
		if queryFields == "" {
			queryFields = "other_url=?"
		} else {
			queryFields += ", other_url=?"
		}

		queryValues = append(queryValues, *params.OtherURL)
	}
	// Handle ps4_gamertag field.
	if params.PS4Gamertag != nil {
		if queryFields == "" {
			queryFields = "ps4_gamertag=?"
		} else {
			queryFields += ", ps4_gamertag=?"
		}

		queryValues = append(queryValues, *params.PS4Gamertag)
	}
	// Handle xbox_gamertag field.
	if params.XBoxGamertag != nil {
		if queryFields == "" {
			queryFields = "xbox_gamertag=?"
		} else {
			queryFields += ", xbox_gamertag=?"
		}

		queryValues = append(queryValues, *params.XBoxGamertag)
	}
	// Handle steam_gamertag field.
	if params.SteamGamertag != nil {
		if queryFields == "" {
			queryFields = "steam_gamertag=?"
		} else {
			queryFields += ", steam_gamertag=?"
		}

		queryValues = append(queryValues, *params.SteamGamertag)
	}

	// Check if the query is empty
	if queryFields == "" {
		return db.GetByID(id)
	} else {
		// Handle updated_at
		queryFields += ", updated_at=?"
		queryValues = append(queryValues, time.Now())
	}

	// Build the full query.
	query := fmt.Sprintf(stmtUpdate, queryFields)
	queryValues = append(queryValues, id)

	// Execute the query
	_, err := db.db.Exec(query, queryValues...)
	if err != nil {
		return nil, err
	}

	// Since the GetByID method is straight forward,
	// we can use this method to retrieve the updated
	// user. Anything more complicated should use the
	// original statement constants.
	return db.GetByID(id)
}

// Delete updates a user's status to 'deleted'
func (db *Database) Delete(id int) error {
	// Execute the query
	_, err := db.db.Exec(stmtDelete, id)
	if err != nil {
		return err
	} else {
		return nil
	}
}
