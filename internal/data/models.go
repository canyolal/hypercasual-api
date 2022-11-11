package data

import (
	"database/sql"
	"errors"
)

var (
	ErrRecordNotFound = errors.New("record not found")
	ErrEditConflict   = errors.New("edit conflict")
)

// Hold all data models
type Models struct {
	Publisher PublisherModel
	Game      GameModel
}

// Initialize all data models for application
func NewModels(db *sql.DB) Models {
	return Models{
		Publisher: PublisherModel{db},
		Game:      GameModel{db},
	}
}
