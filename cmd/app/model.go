package main

import (
	"chat/internal/data"
)

type model struct {
	cursor    int
	choices   []string
	todos     []string
	input     string
	addMode   bool
	viewMode  bool
	width     int
	height    int
	todoModel *data.TodoModel
}
