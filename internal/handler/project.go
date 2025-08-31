package handler

import (
	"context"

	"github.com/RyanHartadi06/clara-be/internal/service"
	"github.com/RyanHartadi06/clara-be/internal/utils"
	"github.com/RyanHartadi06/clara-be/pb/project"
)

type projectHandler struct {
	project.UnimplementedProjectServiceServer

	projectService service.IProjectService
}

func (ph *projectHandler) CreateProject(ctx context.Context, request *project.CreateProjectRequest) (*project.CreateProjectResponse, error) {
	validationErrors, err := utils.CheckValidation(request)
	if err != nil {
		return nil, err
	}

	if validationErrors != nil {
		return &project.CreateProjectResponse{
			Base: utils.ValidationErrorResponse(validationErrors),
		}, nil
	}

	//Process Create Project
	res, err := ph.projectService.CreateProject(ctx, request)
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (ph *projectHandler) DetailProject(ctx context.Context, request *project.DetailProjectRequest) (*project.DetailProjectResponse, error) {
	validationErrors, err := utils.CheckValidation(request)
	if err != nil {
		return nil, err
	}

	if validationErrors != nil {
		return &project.DetailProjectResponse{
			Base: utils.ValidationErrorResponse(validationErrors),
		}, nil
	}

	//Process Create Project
	res, err := ph.projectService.DetailProject(ctx, request)
	if err != nil {
		return nil, err
	}

	return res, nil
}

func NewProjectHandler(projectService service.IProjectService) *projectHandler {
	return &projectHandler{
		projectService: projectService,
	}
}
