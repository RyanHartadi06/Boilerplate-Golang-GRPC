package service

import (
	"context"
	"time"

	"github.com/RyanHartadi06/clara-be/internal/entity"
	"github.com/RyanHartadi06/clara-be/internal/repository"
	"github.com/RyanHartadi06/clara-be/internal/utils"
	projectpb "github.com/RyanHartadi06/clara-be/pb/project"
	"github.com/google/uuid"
)

type IProjectService interface {
	CreateProject(ctx context.Context, request *projectpb.CreateProjectRequest) (*projectpb.CreateProjectResponse, error)
}

type projectService struct {
	projectRepository repository.IProjectRepository
}

func (ps *projectService) CreateProject(ctx context.Context, request *projectpb.CreateProjectRequest) (*projectpb.CreateProjectResponse, error) {
	// Cek project ke database
	project, err := ps.projectRepository.GetProjectByName(ctx, request.Name)

	if err != nil {
		return nil, err
	}

	// Apabila project sudah terdaftar, kita error in

	if project != nil {
		return &projectpb.CreateProjectResponse{
			Base: utils.BadRequestResponse("Project already exists"),
		}, nil
	}

	dateStart, err := time.Parse("2006-01-02", request.DateStart)
	if err != nil {
		return nil, err
	}

	dateEnd, err := time.Parse("2006-01-02", request.DateEnd)
	if err != nil {
		return nil, err
	}

	newProject := entity.Project{
		Id:        uuid.NewString(),
		Name:      request.Name,
		DateStart: dateStart,
		DateEnd:   dateEnd,
		Status:    "TODO",
		CreatedAt: time.Now(),
		CreatedBy: "BACKOFFICE",
	}

	err = ps.projectRepository.InsertProject(ctx, &newProject)
	if err != nil {
		return nil, err
	}
	//insert ke db
	return &projectpb.CreateProjectResponse{
		Base: utils.SuccessResponse("Project created successfully"),
	}, nil
}

func NewProjectService(projectRepository repository.IProjectRepository) IProjectService {
	return &projectService{
		projectRepository: projectRepository,
	}
}
