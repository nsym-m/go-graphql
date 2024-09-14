package services

import (
	"context"

	"github.com/nsym-m/go-graphql/internal/db"
	"github.com/nsym-m/go-graphql/internal/graph/model"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
)

type projectV2Service struct {
	exec boil.ContextExecutor
}

func (p projectV2Service) Create(ctx context.Context, contentID, projectID string) (*model.ProjectV2, error) {

	repo, err := db.Repositories(
		qm.Select(db.RepositoryTableColumns.ID, db.RepositoryTableColumns.Name),
		db.RepositoryWhere.Name.EQ(name),
		db.RepositoryWhere.Owner.EQ(owner),
	).One(ctx, p.exec)
	if err != nil {
		return nil, err
	}
	return p.convertRepository(repo), nil
}

func (p projectV2Service) convertRepository(repo *db.Repository) *model.ProjectV2 {
	return &model.ProjectV2{
		ID:   repo.ID,
		Name: repo.Name,
	}
}
