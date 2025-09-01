package entity

import (
	"database/sql"
	"time"
)

// enum
type ProjectStatus string

const (
	ProjectStatusTodo       ProjectStatus = "todo"
	ProjectStatusInProgress ProjectStatus = "in_progress"
	ProjectStatusCompleted  ProjectStatus = "completed"
)

type Project struct {
	Id        string
	Name      string
	DateStart time.Time
	DateEnd   time.Time
	Status    string
	CreatedAt time.Time
	CreatedBy string
	UpdatedAt sql.NullTime
	UpdatedBy sql.NullString
}
