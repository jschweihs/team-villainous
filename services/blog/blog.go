package blog

import (
	"vil/database"
	dbblog "vil/database/blog"
	dbusers "vil/database/users"
	"vil/services/errors"
)

// Service defines the blog service
type Service struct {
	db *database.Database
}

// New returns a new blog service
func New(db *database.Database) *Service {
	return &Service{
		db: db,
	}
}

// BlogEntry defines a blog entry
type BlogEntry dbblog.BlogEntry

// BlogEntries defines a set of blog entries
type BlogEntries struct {
	BlogEntries []*BlogEntry `json:"blog_entries"`
	Total       int          `json:"total"`
}

// NewParams defines the parameters for the New method.
type NewParams dbblog.NewParams

// New creates a new blog entry
func (s *Service) New(params *NewParams) (*BlogEntry, error) {
	// Create a new ParamErrors
	pes := errors.NewParamErrors()

	// Check title
	if params.Title == "" {
		pes.Add(errors.NewParamError("title", ErrTitleEmpty))
	}

	// Check user id is not empty and is a real user
	if params.UserID < 1 {
		pes.Add(errors.NewParamError("user_id", ErrUserIDEmpty))
	} else {
		_, err := s.db.Users.GetByID(params.UserID)
		if err != nil && err != dbusers.ErrUserNotFound {
			return nil, err
		}
	}

	// Check content
	if params.Content == "" {
		pes.Add(errors.NewParamError("content", ErrContentEmpty))
	}

	// Default preview text if none is provided
	if params.Preview == "" {
		params.Preview = params.Content[0:512]
		if len(params.Content) > 512 {
			params.Preview += "..."
		}
	}

	// Create this blog entry in the database
	dbb, err := s.db.Blog.New(&dbblog.NewParams{
		Title:   params.Title,
		UserID:  params.UserID,
		Preview: params.Preview,
		Content: params.Content,
	})
	if err != nil {
		return nil, err
	}

	// Create a new blog entry
	blogEntry := &BlogEntry{
		ID:        dbb.ID,
		Title:     dbb.Title,
		UserID:    dbb.UserID,
		Preview:   dbb.Preview,
		Content:   dbb.Content,
		CreatedAt: dbb.CreatedAt,
		UpdatedAt: dbb.UpdatedAt,
	}

	return blogEntry, nil
}

// GetParams defines the parameters for the Get method
type GetParams dbblog.GetParams

func (s *Service) Get(params *GetParams) (*BlogEntries, error) {
	// Try to pull the blog entries from the database
	dbbs, err := s.db.Blog.Get(&dbblog.GetParams{
		ID:        params.ID,
		Title:     params.Title,
		UserID:    params.UserID,
		Content:   params.Content,
		CreatedAt: params.CreatedAt,
		UpdatedAt: params.UpdatedAt,
		Offset:    params.Offset,
		Limit:     params.Limit,
	})
	if err != nil {
		return nil, err
	}

	// Create a new BlogEntries
	blogEntries := &BlogEntries{
		BlogEntries: []*BlogEntry{},
		Total:       dbbs.Total,
	}

	// Move blog entry rows into blogEntries
	for _, dbb := range dbbs.BlogEntries {
		blogEntry := &BlogEntry{
			ID:        dbb.ID,
			Title:     dbb.Title,
			UserID:    dbb.UserID,
			Preview:   dbb.Preview,
			Content:   dbb.Content,
			CreatedAt: dbb.CreatedAt,
			UpdatedAt: dbb.UpdatedAt,
		}

		blogEntries.BlogEntries = append(blogEntries.BlogEntries, blogEntry)

	}

	return blogEntries, nil
}

// GetByID retrieves a blog entry by its ID
func (s *Service) GetByID(id int) (*BlogEntry, error) {
	// Get blog entry from database
	be, err := s.db.Blog.GetByID(id)
	if err != nil {
		return nil, err
	}

	// Create a new blog entry
	blogEntry := &BlogEntry{
		ID:        be.ID,
		Title:     be.Title,
		UserID:    be.UserID,
		Preview:   be.Preview,
		Content:   be.Content,
		CreatedAt: be.CreatedAt,
		UpdatedAt: be.UpdatedAt,
	}

	return blogEntry, nil
}

// Update params defines the parameters for the update methods
type UpdateParams dbblog.UpdateParams

// UpdateByID updates a blog entry
func (s *Service) UpdateByID(id int, params *UpdateParams) (*BlogEntry, error) {
	// Try and pull this blog entry from the database
	dbb, err := s.db.Blog.GetByID(id)
	if err == dbblog.ErrBlogEntryNotFound {
		return nil, ErrBlogEntryNotFound
	} else if err != nil {
		return nil, err
	}

	dbb, err = s.db.Blog.Update(id, &dbblog.UpdateParams{
		Title:   params.Title,
		UserID:  params.UserID,
		Preview: params.Preview,
		Content: params.Content,
	})
	if err != nil {
		return nil, err
	}

	blogEntry := &BlogEntry{
		ID:        dbb.ID,
		Title:     dbb.Title,
		UserID:    dbb.UserID,
		Preview:   dbb.Preview,
		Content:   dbb.Content,
		CreatedAt: dbb.CreatedAt,
		UpdatedAt: dbb.UpdatedAt,
	}

	return blogEntry, nil
}

// DeleteByID deletes a blog entry
func (s *Service) DeleteByID(id int) error {
	// Try and pull this blog entry from the database
	_, err := s.db.Blog.GetByID(id)
	if err == dbblog.ErrBlogEntryNotFound {
		return ErrBlogEntryNotFound
	} else if err != nil {
		return err
	}

	err = s.db.Blog.Delete(id)
	if err != nil {
		return err
	}

	return nil
}
