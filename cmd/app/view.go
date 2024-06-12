package main

import (
	"fmt"

	"github.com/charmbracelet/lipgloss"
)

func (m model) View() string {
	var content string

	if m.addMode {
		content = fmt.Sprintf(
			"Enter todo item: %s\n\nPress Enter to save, Backspace to delete, and q to quit.\n",
			m.input,
		)
	} else if m.viewMode {
		if len(m.todos) == 0 {
			content = "No todos found.\n\nPress q to quit.\n"
		} else {
			content = "Todos:\n\n"
			for i, todo := range m.todos {
				content += fmt.Sprintf("%d. %s\n", i+1, todo)
			}
			content += "\nPress q to quit.\n"
		}
	} else {
		content = "What do you want to do?\n\n"
		for i, choice := range m.choices {
			cursor := " "
			if m.cursor == i {
				cursor = ">"
			}
			content += fmt.Sprintf("%s %s\n", cursor, choice)
		}
		content += "\nPress q to quit.\n"
	}

	// Use Lipgloss to style and center the content
	windowStyle := lipgloss.NewStyle().
		Width(m.width).
		Height(m.height).
		Align(lipgloss.Center, lipgloss.Center)

	return windowStyle.Render(content)
}
