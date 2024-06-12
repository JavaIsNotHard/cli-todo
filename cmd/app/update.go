package main

import (
	tea "github.com/charmbracelet/bubbletea"
)

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {

	case tea.KeyMsg:
		switch msg.String() {

		case "ctrl+c", "q":
			return m, tea.Quit

		case "up", "k":
			if !m.addMode && !m.viewMode {
				if m.cursor > 0 {
					m.cursor--
				}
			}

		case "down", "j":
			if !m.addMode && !m.viewMode {
				if m.cursor < len(m.choices)-1 {
					m.cursor++
				}
			}

		case "enter":
			if m.addMode {
				m.todos = append(m.todos, m.input)
				m.addMode = false
				m.input = ""
			} else if m.viewMode {
				m.viewMode = false
			} else {
				switch m.cursor {
				case 0:
					m.addMode = true
				case 1:
					m.viewMode = true
				}
			}

		case "backspace":
			if m.addMode && len(m.input) > 0 {
				m.input = m.input[:len(m.input)-1]
			}

		default:
			if m.addMode {
				m.input += msg.String()
			}
		}

	case tea.WindowSizeMsg:
		m.width = msg.Width
		m.height = msg.Height
	}

	return m, nil
}
