package data

import (
	"database/sql"
)

type Models struct {
	Todo TodoModel
}

func NewModel(db *sql.DB) Models {
	return Models{
		Todo: TodoModel{DB: db},
	}
}
