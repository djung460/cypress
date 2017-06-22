package models

import "errors"

// ErrNotFound whenever the entry is not found in the db
var ErrNotFound = errors.New("not found")

// DB is the interface to the database.
// The idea here is to have multiple different DB implementations.
// This codebase provides an in-memory database and a sqlite3 db
type DB interface {
	// Upsert sets or creates the value for the existing username and article.
	// returns true and nil
	// if the article was created on this call, false and nil if the article was not created
	// but still successfully updated, and false and the appropriate error otherwise
	CreateNugget(nugget Nugget) (bool, error)

	// GetNuggetsByUser gets all articles associated with a username and places
	// it inside Model
	// returns the value and nil if it was found, nil and ErrNotFound if it
	// was not found and another appropriate error otherwise
	GetNuggetsByUser(id int) ([]Nugget, error)

	CreateCategory(category Category) (bool, error)

	GetCategories() ([]Category, error)
}
