package users

import (
	"database/sql"
	"fmt"
	"reflect"
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
	Status        int        `json:"status"`
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
	if params.Status > 0 {
		if queryFields == "" {
			queryFields = "WHERE status=?"
		} else {
			queryFields += " AND status=?"
		}

		queryValues = append(queryValues, params.Status)
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
	Username      *string    `json:"username"`
	Password      *string    `json:"password"`
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
	UpdatedAt     *time.Time `json:"updated_at"`
}

// Update updates a user
func (db *Database) Update(id int, params *UpdateParams) (*User, error) {
	// Create variables to hold the query fields
	// being updated and their new values
	var queryFields string
	var queryValues []interface{}

	// Get fields and their values
	fields := reflect.TypeOf(*params)
	values := reflect.ValueOf(*params)
	size := fields.NumField()

	// Range through provided params to build query string
	for i := 0; i < size; i++ {
		// Get field and value from set of fields and values
		field := fields.Field(i)
		value := values.Field(i)

		// Only generate query based on values that are provided
		if !value.IsNil() {
			// Get value from value
			value = reflect.Indirect(value)

			// Assume json tag matches database column
			// If we change our mind about this assumption,
			// We can use another set of data for reference
			tag := field.Tag.Get("json")
			if queryFields == "" {
				queryFields = tag + "=?"
			} else {
				queryFields += ", " + tag + "=?"
			}
			queryValues = append(queryValues, value.String())
		}
	}

	// Check if the query is empty
	if queryFields == "" {
		return db.GetByID(id)
	}

	// Build the full query.
	query := fmt.Sprintf(stmtUpdate, queryFields)
	queryValues = append(queryValues, id)

	// Execute the query
	_, err := db.db.Exec(query, queryValues...)
	if err != nil {
		fmt.Println(err)
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
		fmt.Println(err)
		return err
	} else {
		return nil
	}
}
