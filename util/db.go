package util

import (
	"database/sql"

	"github.com/charmbracelet/log"
	_ "modernc.org/sqlite"
)

type LinkTypes = int

const (
	LinkTaskList LinkTypes = iota + 1
	LinkListMilestone
	LinkMilestoneProject
)

type ContainerType = int

const (
	List ContainerType = iota + 1
	Milestione
	Project
)

var db *sql.DB

func SetupDB(l *log.Logger) {
	var err error

	db, err = sql.Open("sqlite", "goaler.db")
	if err != nil {
		l.Fatal(err)
	}

	query := `
CREATE TABLE IF NOT EXISTS task (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    user_id TEXT NOT NULL,
    title TEXT,
    description TEXT,
    status TEXT DEFAULT 'pending',
    reminder TEXT,
    deadline TEXT,
    quest INTEGER DEFAULT 0
);

CREATE TABLE IF NOT EXISTS container (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    user_id TEXT NOT NULL,
    title TEXT,
    description TEXT,
    type INTEGER
);

CREATE TABLE IF NOT EXISTS liks (
    container_id INTEGER,
    task_id INTEGER,
    link_type INTEGER,
    PRIMARY KEY (container_id, task_id),
    FOREIGN KEY (container_id) REFERENCES container(id) ON DELETE CASCADE,
    FOREIGN KEY (task_id) REFERENCES task(id) ON DELETE CASCADE
);
`
	_, err = db.Exec(query)
	if err != nil {
		l.Fatal("Error creando la tabla:", err)
	}
}

func CreateTask(task *Task) error {
	query := `INSERT INTO task (user_id, title, description, reminder, deadline)
	VALUES (?, ?, ?, ?, ?)`
	_, err := db.Exec(
		query,
		task.UserID,
		task.Title,
		task.Description,
		task.Reminder,
		task.Deadline,
	)
	return err
}

func CreateContainer(c *Container) error {
	query := `INSERT INTO list (user_id, title, description, type)
	VALUES (?, ?, ?, ?)`
	_, err := db.Exec(
		query,
		c.UserID,
		c.Title,
		c.Description,
		c.Type,
	)
	return err
}
