package services

import (
	"context"
	"log"

	"github.com/nsym-m/go-graphql/internal/db"
	"github.com/nsym-m/go-graphql/internal/graph/model"
	"github.com/volatiletech/sqlboiler/v4/boil"
)

type projectV2Service struct {
	exec boil.ContextExecutor
}

func (p *projectV2Service) GetProjectByID(ctx context.Context, id string) (*model.ProjectV2, error) {

	project, err := db.FindProject(ctx, p.exec, id,
		db.ProjectColumns.ID,
		db.ProjectColumns.Title,
		db.ProjectColumns.URL,
		db.ProjectColumns.Owner,
	)
	if err != nil {
		return nil, err
	}
	return p.convertProjectV2(project), nil
}

func (p *projectV2Service) Create(ctx context.Context, contentID, projectID string) (*model.ProjectV2, error) {

	return nil, nil
}

func (p *projectV2Service) convertProjectV2(project *db.Project) *model.ProjectV2 {
	projectURL, err := model.UnmarshalURI(project.URL)
	if err != nil {
		log.Println("invalid URI", project.URL)
	}

	return &model.ProjectV2{
		ID:    project.ID,
		Title: project.Title,
		URL:   projectURL,
		Owner: &model.User{ID: project.Owner},
	}
}
