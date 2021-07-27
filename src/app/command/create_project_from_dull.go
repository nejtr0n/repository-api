package command

import (
	"context"

	"github.com/nejtr0n/repository-api/src/domain"
)

type CreateProjectFromDull struct {
	Dull struct {
		ProviderType domain.ProviderType `json:"provider_type" validate:"required"`
		ProjectId    int                 `json:"project_id" validate:"required"`
	} `json:"dull"`
	New struct {
		ProviderType domain.ProviderType `json:"provider_type" validate:"required"`
		Name         string              `json:"name" validate:"required"`
		GroupId      int                 `json:"group_id"`
	} `json:"new"`
}

func NewCreateProjectFromDullHandler(service domain.Service) CreateProjectFromDullHandler {
	return CreateProjectFromDullHandler{service: service}
}

type CreateProjectFromDullHandler struct {
	service domain.Service
}

func (c CreateProjectFromDullHandler) Handle(ctx context.Context, cmd CreateProjectFromDull) error {
	dull, err := c.service.GetProject(cmd.Dull.ProviderType, cmd.Dull.ProjectId)
	if err != nil {
		return err
	}

	target, err := c.service.CreateProject(cmd.New.ProviderType, cmd.New.Name, cmd.New.GroupId)
	if err != nil {
		return err
	}

	err = c.service.CloneProject(dull, target)
	if err != nil {
		return err
	}

	return nil
}
