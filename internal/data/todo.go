package data

import (
	"context"
	"database/sql"
	"encoding/json"
	"errors"
	"time"
)

type Todo struct {
	ID      int64     `json:"id"`
	Title   string    `json:"title"`
	Items   []string  `json:"items"`
	Created time.Time `json:"created_at"`
}

var (
	NotRecordError = errors.New("Record could not be found")
)

type TodoModel struct {
	DB *sql.DB
}

func (m TodoModel) GetSingleItem(id int64) (*Todo, error) {
	query := `
		SELECT title, items FROM todo;
	`

	var item Todo

	ctx, cancel := context.WithTimeout(context.TODO(), 3*time.Second)
	defer cancel()

	err := m.DB.QueryRowContext(ctx, query, id).Scan(
		&item.Title,
		&item.Items,
	)

	if err != nil {
		switch {
		case errors.Is(err, sql.ErrNoRows):
			return nil, NotRecordError
		default:
			return nil, err
		}
	}

	return &item, nil
}

func (m TodoModel) InsertItem(todo *Todo) error {
	query := `
		INSERT INTO todo (title, items) VALUES (?, ?)
		RETURNING title, items;
	`
	jsonData, err := json.Marshal(todo.Items)

	if err != nil {
		return err
	}

	args := []any{todo.Title, jsonData}

	context, cancel := context.WithTimeout(context.TODO(), 3*time.Second)
	defer cancel()

	_, err = m.DB.ExecContext(context, query, args...)

	return err
}
