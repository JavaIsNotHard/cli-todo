package main

import (
	"database/sql"
	"log"
	"os"

	"chat/internal/data"
	tea "github.com/charmbracelet/bubbletea"
)

func openDB() *sql.DB {
	db, err := sql.Open("sqlite3", "./users.db")
	if err != nil {
		log.Fatal(err)
	}

	return db
}

func main() {
	db := openDB()
	defer db.Close()

	var todoModel data.NewModel(db)

	p := tea.NewProgram(initialModel(), tea.WithAltScreen())

	if _, err := p.Run(); err != nil {
		log.Fatalf("Error: %v", err)
		os.Exit(1)
	}
}
