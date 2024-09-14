package services

import (
	"context"

	"github.com/nsym-m/go-graphql/internal/db"
	"github.com/nsym-m/go-graphql/internal/graph/model"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
)

type ownerService struct {
	exec boil.ContextExecutor
}

func (o ownerService) GetUserByID(ctx context.Context, id string) (*model.User, error) {

	repo, err := db.Users(
		qm.Select(
			db.UserColumns.ID,
		),
		db.UserWhere.ID.EQ(id),
	).One(ctx, o.exec)
	if err != nil {
		return nil, err
	}
	return o.convertRepository(repo), nil
}

func (o ownerService) convertRepository(user *db.User) *model.User {
	return &model.User{
		ID: user.ID,
	}
}
