package data

import (
	"database/sql"
)

func NewModel(db *sql.DB) *TodoModel {
	return &TodoModel{
		DB: db,
	}
}
