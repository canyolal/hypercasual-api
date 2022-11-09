package data

import "database/sql"

type Models struct {
	Publisher PublisherModel
}

func NewModels(db *sql.DB) Models {
	return Models{
		Publisher: PublisherModel{db},
	}
}
