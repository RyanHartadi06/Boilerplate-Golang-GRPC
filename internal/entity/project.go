package entity

import (
	"database/sql"
	"time"
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
