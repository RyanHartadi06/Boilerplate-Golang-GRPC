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
	GetProjectById(ctx context.Context, id string) (*entity.Project, error)
}

type projectRepository struct {
	db *sql.DB
}

func (repo *projectRepository) GetProjectByName(ctx context.Context, name string) (*entity.Project, error) {
	row := repo.db.QueryRowContext(ctx, "SELECT id, name FROM project WHERE name = $1", name)
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

func (repo *projectRepository) InsertProject(ctx context.Context, project *entity.Project) error {
	_, err := repo.db.ExecContext(ctx, "INSERT INTO project (id, name, date_start, date_end, status, created_at, created_by) VALUES ($1, $2, $3, $4, $5, $6, $7)",
		project.Id,
		project.Name,
		project.DateStart,
		project.DateEnd,
		project.Status,
		project.CreatedAt,
		project.CreatedBy,
	)

	if err != nil {
		return err
	}

	return nil

}

func (repo *projectRepository) GetProjectById(ctx context.Context, id string) (*entity.Project, error) {
	var productEntity entity.Project

	row := repo.db.QueryRowContext(ctx, "SELECT id, name, date_start, date_end, status, created_at FROM project WHERE id = $1", id)

	if row.Err() != nil {
		return nil, row.Err()
	}

	err := row.Scan(
		&productEntity.Id,
		&productEntity.Name,
		&productEntity.DateStart,
		&productEntity.DateEnd,
		&productEntity.Status,
		&productEntity.CreatedAt,
	)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}
		return nil, err
	}
	return &productEntity, nil
}

func NewProjectRepository(db *sql.DB) IProjectRepository {
	return &projectRepository{
		db: db,
	}
}
