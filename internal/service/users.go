package service

import (
	"context"

	"github.com/nsym-m/go-graphql/internal/db"
	"github.com/nsym-m/go-graphql/internal/graph/model"
	"github.com/volatiletech/sqlboiler/boil"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
)

type userService struct {
	exec boil.ContextExecutor
}

func (u userService) GetUserByName(ctx context.Context, name string) (*model.User, error) {

	user, err := db.Users(
		qm.Select(db.UserTableColumns.ID, db.UserTableColumns.Name),
		db.UserWhere.Name.EQ(name),
	).One(ctx, u.exec)
	if err != nil {
		return nil, err
	}

	return u.convertUser(user), nil
}

func (u userService) ListUserByIDs(ctx context.Context, ids []string) ([]*model.User, error) {

	users, err := db.Users(
		qm.Select(db.UserTableColumns.ID, db.UserTableColumns.Name),
		db.UserWhere.ID.IN(ids),
	).All(ctx, u.exec)
	if err != nil {
		return nil, err
	}

	return u.convertUsers(users), nil
}

func (u userService) convertUser(user *db.User) *model.User {
	return &model.User{
		ID:   user.ID,
		Name: user.Name,
	}
}

func (u userService) convertUsers(users []*db.User) []*model.User {
	res := make([]*model.User, len(users))
	for i, user := range users {
		res[i] = u.convertUser(user)
	}
	return res
}
