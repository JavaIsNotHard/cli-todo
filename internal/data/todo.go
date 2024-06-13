package data

import (
	"context"
	"database/sql"
	"errors"
	"time"
)

type Todo struct {
	ID      int64
	Items   string
	Created time.Time
}

var (
	NotRecordError = errors.New("Record could not be found")
)

type TodoModel struct {
	DB *sql.DB
}

func (m TodoModel) GetAllTodo(id int64) ([]string, error) {
	query := `
		SELECT items FROM todo;
	`
	var item []Todo

	ctx, cancel := context.WithTimeout(context.TODO(), 3*time.Second)
	defer cancel()

	row, err := m.DB.QueryContext(ctx, query)

	for row.Next() {
		var todo Todo
		err := row.Scan(&todo.Items)
		if err != nil {
			return nil, err
		}
		item = append(item, todo)
	}

	err = row.Err()
	if err != nil {
		switch {
		case errors.Is(err, sql.ErrNoRows):
			return nil, NotRecordError
		default:
			return nil, err
		}
	}

	var result []string
	for _, entires := range item {
		result = append(result, entires.Items)
	}

	return result, nil
}

func (m TodoModel) InsertItem(todo *Todo) error {
	query := `
		INSERT INTO todo (items) VALUES (?)
	`
	args := []any{todo.Items}

	context, cancel := context.WithTimeout(context.TODO(), 3*time.Second)
	defer cancel()

	_, err := m.DB.ExecContext(context, query, args...)

	return err
}
