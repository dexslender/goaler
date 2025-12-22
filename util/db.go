// Package util
package util

import (
	"database/sql"

	"github.com/charmbracelet/log"
	_ "modernc.org/sqlite"
)

var db *sql.DB

func SetupDB(l *log.Logger) {
	var err error

	db, err = sql.Open("sqlite", "goaler.db")
	if err != nil {
		l.Fatal(err)
	}

	query := `
	CREATE TABLE IF NOT EXISTS goals (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		user_id TEXT,
		title TEXT,
		type TEXT,
		status TEXT DEFAULT 'pending'
	);

	CREATE TABLE IF NOT EXISTS task (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		user_id TEXT,
		title TEXT,
		description TEXT,
		status TEXT DEFAULT 'pending',
		reminder TEXT,
		deadline TEXT
	);
	
	CREATE TABLE IF NOT EXISTS todo_list (
		
	);

	CREATE TABLE IF NOT EXISTS milestone ();

	CREATE TABLE IF NOT EXISTS project ();

	CREATE TABLE IF NOT EXISTS habit ();
	`

	_, err = db.Exec(query)
	if err != nil {
		l.Fatal("Error creando la tabla:", err)
	}
}

func CreateGoal(userID, title, goalType string) error {
	query := `INSERT INTO goals (user_id, title, type) VALUES (?, ?, ?)`
	_, err := db.Exec(query, userID, title, goalType)
	return err
}
