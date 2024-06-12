package main

type model struct {
	cursor   int
	choices  []string
	todos    []string
	input    string
	addMode  bool
	viewMode bool
	width    int
	height   int
}

func initialModel() model {
	return model{
		choices:  []string{"Add Todo", "View Todos"},
		todos:    []string{},
		input:    "",
		addMode:  false,
		viewMode: false,
	}
}
