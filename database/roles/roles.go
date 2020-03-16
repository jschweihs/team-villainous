package roles

import (
	"database/sql"
	"fmt"
	"reflect"
)

// Database defines the roles database
type Database struct {
	db *sql.DB
}

// New creates a new roles database
func New(db *sql.DB) *Database {
	return &Database{
		db: db,
	}
}

// Role defines a role
type Role struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

// Roles defines a set of roles
type Roles struct {
	Roles []*Role `json:"roles"`
	Total int     `json:"total"`
}

const (
	// stmtInsert defines the SQL statement to
	// insert a new role into the database
	stmtInsert = `
INSERT INTO roles (
	name
)
VALUES (
	?
)
`

	// stmtSelect defines the SQL statement to
	// select a set of roles.
	stmtSelect = `
SELECT 
	id,
	name
FROM roles
%s
LIMIT %v, %v
`

	// stmtSelectCount defines the SQL statement to
	// select the total number of roles according to the filters.
	stmtSelectCount = `
SELECT COUNT(*)
FROM roles
%s
`

	// stmtSelectByID defines the SQL statement to
	// select a role by their ID.
	stmtSelectByID = `
SELECT 
	id,
	name
FROM roles
WHERE id=?
`

	// stmtSelectByName defines the SQL statement to
	// select a role by their name.
	stmtSelectByName = `
SELECT 
	id,
	name
FROM roles
WHERE name=?
`

	// stmtUpdate defines the SQL statement to
	// update a role.
	stmtUpdate = `
UPDATE roles
SET %s
WHERE id=?
`

	// stmtDelete defines the SQL statement to
	// remove a role.
	stmtDelete = `
DELETE FROM roles
WHERE id=?
`
)

// NewParams defines the parameters for the New method
// Similar to Role type but does not have ID
type NewParams struct {
	Name string `json:"name"`
}

// New creates a new role.
func (db *Database) New(params *NewParams) (*Role, error) {
	// Create a new role
	role := &Role{
		Name: params.Name,
	}

	// Create variable to hold the result
	var res sql.Result
	var err error

	// Execute the query
	if res, err = db.db.Exec(stmtInsert, role.Name); err != nil {
		return nil, err
	}

	// Get last insert ID
	id, err := res.LastInsertId()
	if err != nil {
		return nil, err
	}
	role.ID = int(id)

	return role, nil
}

// GetByID retrieves a role by their ID
func (db *Database) GetByID(id int) (*Role, error) {
	// Create a new role
	role := &Role{}

	// Execute the query
	err := db.db.QueryRow(stmtSelectByID, id).Scan(&role.ID, &role.Name)
	switch {
	case err == sql.ErrNoRows:
		return nil, ErrRoleNotFound
	case err != nil:
		return nil, err
	}

	return role, nil
}

// GetByName retrieves a role by their ID
func (db *Database) GetByName(name string) (*Role, error) {
	// Create a new role
	role := &Role{}

	// Execute the query
	err := db.db.QueryRow(stmtSelectByName, name).Scan(&role.ID, &role.Name)
	switch {
	case err == sql.ErrNoRows:
		return nil, ErrRoleNotFound
	case err != nil:
		return nil, err
	}

	return role, nil
}

// GetParams defines the parameters for the Get method.
type GetParams struct {
	ID     *int    `json:"id"`
	Name   *string `json:"name"`
	Offset int     `json:"offset"`
	Limit  int     `json:"limit"`
}

// GetAll retrieves all roles
func (db *Database) Get(params *GetParams) (*Roles, error) {
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

	// Handle name field
	if params.Name != nil {
		if queryFields == "" {
			queryFields = "WHERE name=?"
		} else {
			queryFields += " AND name=?"
		}

		queryValues = append(queryValues, *params.Name)
	}

	query := fmt.Sprintf(stmtSelect, queryFields, params.Offset, params.Limit)

	// Create new roles
	roles := &Roles{
		Roles: []*Role{},
	}

	// Execute the query and receive rows of roles
	roleRows, err := db.db.Query(query, queryValues...)
	if err != nil {
		return nil, err
	}
	defer roleRows.Close()

	for roleRows.Next() {
		// Create a new role
		role := &Role{}

		// Place data from role row into role
		if err := roleRows.Scan(&role.ID, &role.Name); err != nil {
			return nil, err
		}

		// Add role to result roles slice
		roles.Roles = append(roles.Roles, role)
	}
	if err = roleRows.Err(); err != nil {
		return nil, err
	}

	// Build the total count query
	queryCount := fmt.Sprintf(stmtSelectCount, queryFields)

	// Get total count
	var total int
	if err = db.db.QueryRow(queryCount, queryValues...).Scan(&total); err != nil {
		return nil, err
	}
	roles.Total = total

	return roles, nil
}

// UpdateParams defines the parameters for the Update method.
type UpdateParams struct {
	Name *string `json:"name"`
}

// Update updates a role
func (db *Database) Update(id int, params *UpdateParams) (*Role, error) {
	// Create variables to hold the query fields
	// being updated and their new values
	var queryFields string
	var queryValues []interface{}

	// Handle name field
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
	// role. Anything more complicated should use the
	// original statement constants.
	return db.GetByID(id)
}

// Delete removes a role
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
