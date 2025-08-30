package repository

import (
	"context"
	"database/sql"
	"errors"

	"github.com/RyanHartadi06/clara-be/internal/entity"
)

type IProjectRepository interface {
	GetProjectByName(ctx context.Context, name string) (*entity.Project, error)
	InsertProject(ctx context.Context, project *entity.Project) error
}

type projectRepository struct {
	db *sql.DB
}

func (pr *projectRepository) GetProjectByName(ctx context.Context, name string) (*entity.Project, error) {
	row := pr.db.QueryRowContext(ctx, "SELECT id, name FROM project WHERE name = $1", name)
	if row.Err() != nil {
		return nil, row.Err()
	}

	var project entity.Project

	err := row.Scan(
		&project.Id,
		&project.Name,
	)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}
		return nil, err
	}

	return &project, nil
}

func (pr *projectRepository) InsertProject(ctx context.Context, project *entity.Project) error {
	_, err := pr.db.ExecContext(ctx, "INSERT INTO project (id, name, date_start, date_end, status, created_at, created_by, updated_at, updated_by) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9)",
		project.Id,
		project.Name,
		project.DateStart,
		project.DateEnd,
		project.Status,
		project.CreatedAt,
		project.CreatedBy,
		project.UpdatedAt,
		project.UpdatedBy,
	)

	if err != nil {
		return err
	}

	return nil

}

func NewProjectRepository(db *sql.DB) IProjectRepository {
	return &projectRepository{
		db: db,
	}
}
