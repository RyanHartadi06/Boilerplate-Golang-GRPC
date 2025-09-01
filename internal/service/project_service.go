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
	DetailProject(ctx context.Context, request *projectpb.DetailProjectRequest) (*projectpb.DetailProjectResponse, error)
	DeleteProject(context.Context, *projectpb.DeleteProjectRequest) (*projectpb.DeleteProjectResponse, error)
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
		Status:    string(entity.ProjectStatusTodo),
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

func (ps *projectService) DetailProject(ctx context.Context, request *projectpb.DetailProjectRequest) (*projectpb.DetailProjectResponse, error) {
	// Cek project ke database
	project, err := ps.projectRepository.GetProjectById(ctx, request.Id)
	if err != nil {
		return nil, err
	}

	if project == nil {
		return &projectpb.DetailProjectResponse{
			Base: utils.NotFoundResponse("Not Found"),
		}, nil
	}

	return &projectpb.DetailProjectResponse{
		Base:      utils.SuccessResponse("Project found"),
		Id:        project.Id,
		Name:      project.Name,
		DateStart: project.DateStart.Format("2006-01-02"),
		DateEnd:   project.DateEnd.Format("2006-01-02"),
		Status:    project.Status,
		CreatedAt: project.CreatedAt.Format("2006-01-02 15:04:05"),
	}, nil
}
func (ps *projectService) DeleteProject(ctx context.Context, request *projectpb.DeleteProjectRequest) (*projectpb.DeleteProjectResponse, error) {
	// Cek project ke database
	project, err := ps.projectRepository.GetProjectById(ctx, request.Id)
	if err != nil {
		return nil, err
	}

	if project == nil {
		return &projectpb.DeleteProjectResponse{
			Base: utils.NotFoundResponse("Not Found"),
		}, nil
	}

	error := ps.projectRepository.DeleteProjectId(ctx, request.Id)
	if error != nil {
		return nil, error
	}

	return &projectpb.DeleteProjectResponse{
		Base: utils.SuccessResponse("Project deleted successfully"),
	}, nil
}

func NewProjectService(projectRepository repository.IProjectRepository) IProjectService {
	return &projectService{
		projectRepository: projectRepository,
	}
}
