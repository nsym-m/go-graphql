package services

import (
	"context"

	"github.com/nsym-m/go-graphql/internal/db"
	"github.com/nsym-m/go-graphql/internal/graph/model"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
)

type repositoryService struct {
	exec boil.ContextExecutor
}

func (r *repositoryService) GetRepoByID(ctx context.Context, id string) (*model.Repository, error) {

	repo, err := db.Repositories(
		qm.Select(db.RepositoryTableColumns.ID, db.RepositoryTableColumns.Name, db.RepositoryTableColumns.Owner),
		db.RepositoryWhere.ID.EQ(id),
	).One(ctx, r.exec)
	if err != nil {
		return nil, err
	}
	return r.convertRepository(repo), nil
}

func (r *repositoryService) GetRepoByFullName(ctx context.Context, name, owner string) (*model.Repository, error) {

	repo, err := db.Repositories(
		qm.Select(db.RepositoryTableColumns.ID, db.RepositoryTableColumns.Name, db.RepositoryTableColumns.Owner),
		db.RepositoryWhere.Name.EQ(name),
		db.RepositoryWhere.Owner.EQ(owner),
	).One(ctx, r.exec)
	if err != nil {
		return nil, err
	}
	return r.convertRepository(repo), nil
}

func (r *repositoryService) convertRepository(repo *db.Repository) *model.Repository {
	return &model.Repository{
		ID:    repo.ID,
		Name:  repo.Name,
		Owner: &model.User{ID: repo.Owner},
	}
}
