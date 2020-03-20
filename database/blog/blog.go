package blog

import (
	"database/sql"
	"fmt"
	"reflect"
	"time"
)

// Database defines the blog database
type Database struct {
	db *sql.DB
}

// New creates a new blog database
func New(db *sql.DB) *Database {
	return &Database{
		db: db,
	}
}

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

// BlogEntries defines a set of blog entries
type BlogEntries struct {
	BlogEntries []*BlogEntry `json:"blog_entries"`
	Total       int          `json:"total"`
}

const (
	// stmtInsert defines the SQL statement to
	// insert a new blog entry into the database
	stmtInsert = `
INSERT INTO blog (
	title, 
	user_id, 
	preview, 
	content, 
	created_at, 
	updated_at
)
VALUES (
	?, ?, ?, ?, ?, ?
)
`

	// stmtSelect defines the SQL statement to
	// select a set of blog entries.
	stmtSelect = `
SELECT 
	id,
	title, 
	user_id, 
	preview, 
	content, 
	created_at, 
	updated_at
FROM blog
%s
LIMIT %v, %v
`

	// stmtSelectCount defines the SQL statement to
	// select the total number of blog entries according to the filters.
	stmtSelectCount = `
SELECT COUNT(*)
FROM blog
%s
`

	// stmtSelectByID defines the SQL statement to
	// select a blog entry by their ID.
	stmtSelectByID = `
SELECT 
	id,
	title, 
	user_id, 
	preview, 
	content, 
	created_at, 
	updated_at
FROM blog
WHERE id=?
`

	// stmtUpdate defines the SQL statement to
	// update a blog entry.
	stmtUpdate = `
UPDATE blog
SET %s
WHERE id=?
`

	// stmtDelete defines the SQL statement to
	// remove a blog entry.
	stmtDelete = `
DELETE FROM blog
WHERE id=?
`
)

// NewParams defines the parameters for the New method
// Similar to BlogEntry type but does not have ID or
// created/updated timestamps
type NewParams struct {
	Title   string `json:"title"`
	UserID  int    `json:"user_id"`
	Preview string `json:"preview"`
	Content string `json:"content"`
}

// New creates a new blog entry.
func (db *Database) New(params *NewParams) (*BlogEntry, error) {
	// Create a new blog entry
	blogEntry := &BlogEntry{
		Title:     params.Title,
		UserID:    params.UserID,
		Preview:   params.Preview,
		Content:   params.Content,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	// Create variable to hold the result
	var res sql.Result
	var err error

	// Execute the query
	if res, err = db.db.Exec(stmtInsert, blogEntry.Title, blogEntry.UserID, blogEntry.Preview, blogEntry.Content, blogEntry.CreatedAt, blogEntry.UpdatedAt); err != nil {
		return nil, err
	}

	// Get last insert ID
	id, err := res.LastInsertId()
	if err != nil {
		return nil, err
	}
	blogEntry.ID = int(id)

	return blogEntry, nil
}

// GetByID retrieves a blog entry by its ID
func (db *Database) GetByID(id int) (*BlogEntry, error) {
	// Create a new blog entry
	blogEntry := &BlogEntry{}

	// Execute the query
	err := db.db.QueryRow(stmtSelectByID, id).Scan(&blogEntry.ID, &blogEntry.Title, &blogEntry.UserID, &blogEntry.Preview, &blogEntry.Content, &blogEntry.CreatedAt, &blogEntry.UpdatedAt)
	switch {
	case err == sql.ErrNoRows:
		return nil, ErrBlogEntryNotFound
	case err != nil:
		return nil, err
	}

	return blogEntry, nil
}

// GetParams defines the parameters for the Get method.
type GetParams struct {
	ID        *int       `json:"id"`
	Title     *string    `json:"title"`
	UserID    *int       `json:"user_id"`
	Content   *string    `json:"content"`
	CreatedAt *time.Time `json:"created_at"`
	UpdatedAt *time.Time `json:"updated_at"`
	Offset    int        `json:"offset"`
	Limit     int        `json:"limit"`
}

// Get retrieves blog entries
func (db *Database) Get(params *GetParams) (*BlogEntries, error) {
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

	// Handle title field
	if params.CreatedAt != nil {
		if queryFields == "" {
			queryFields = "WHERE title LIKE'%?%'"
		} else {
			queryFields += " AND title LIKE'%?%'"
		}

		queryValues = append(queryValues, *params.Title)
	}

	// Handle user_id field
	if params.CreatedAt != nil {
		if queryFields == "" {
			queryFields = "WHERE user_id=?"
		} else {
			queryFields += " AND user_id=?"
		}

		queryValues = append(queryValues, *params.UserID)
	}

	// Handle content field
	if params.CreatedAt != nil {
		if queryFields == "" {
			queryFields = "WHERE content LIKE'%?%'"
		} else {
			queryFields += " AND content LIKE'%?%'"
		}

		queryValues = append(queryValues, *params.Content)
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

	query := fmt.Sprintf(stmtSelect, queryFields, params.Offset, params.Limit)

	// Create new blogEntries
	blogEntries := &BlogEntries{
		BlogEntries: []*BlogEntry{},
	}

	// Execute the query and receive rows of blog entries
	blogRows, err := db.db.Query(query, queryValues...)
	if err != nil {
		return nil, err
	}
	defer blogRows.Close()

	for blogRows.Next() {
		// Create a new blogEntry
		blogEntry := &BlogEntry{}

		// Place data from blog row into blog entry
		if err := blogRows.Scan(&blogEntry.ID, &blogEntry.Title, &blogEntry.UserID, &blogEntry.Preview, &blogEntry.Content, &blogEntry.CreatedAt, &blogEntry.UpdatedAt); err != nil {
			return nil, err
		}

		// Add blogEntry to result blogEntrys slice
		blogEntries.BlogEntries = append(blogEntries.BlogEntries, blogEntry)
	}
	if err = blogRows.Err(); err != nil {
		return nil, err
	}

	// Build the total count query
	queryCount := fmt.Sprintf(stmtSelectCount, queryFields)

	// Get total count
	var total int
	if err = db.db.QueryRow(queryCount, queryValues...).Scan(&total); err != nil {
		return nil, err
	}
	blogEntries.Total = total

	return blogEntries, nil
}

// UpdateParams defines the parameters for the Update method.
type UpdateParams struct {
	Title     *string    `json:"title"`
	UserID    *int       `json:"user_id"`
	Preview   *string    `json:"preview"`
	Content   *string    `json:"content"`
	UpdatedAt *time.Time `json:"updated_at"`
}

// Update updates a blog entry
func (db *Database) Update(id int, params *UpdateParams) (*BlogEntry, error) {
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
		return nil, err
	}

	// Since the GetByID method is straight forward,
	// we can use this method to retrieve the updated
	// blog entry. Anything more complicated should use the
	// original statement constants.
	return db.GetByID(id)
}

// Delete deletes a blog entry
func (db *Database) Delete(id int) error {
	// Execute the query
	_, err := db.db.Exec(stmtDelete, id)
	if err != nil {
		return err
	} else {
		return nil
	}
}
