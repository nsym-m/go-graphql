package services

import (
	"context"

	"github.com/nsym-m/go-graphql/internal/graph/model"
	"github.com/volatiletech/sqlboiler/v4/boil"
)

type Services interface {
	UserService
	RepositoryService
	IssueService
	OwnerService
	ProjectService
	PullRequestService
}

type services struct {
	*userService
	*repositoryService
	*issueService
	*ownerService
	*projectV2Service
	*pullRequestService
}

type UserService interface {
	GetUserByName(ctx context.Context, name string) (*model.User, error)
	ListUserByIDs(ctx context.Context, ids []string) ([]*model.User, error)
}

type RepositoryService interface {
	GetRepoByFullName(ctx context.Context, name, owner string) (*model.Repository, error)
	GetRepoByID(ctx context.Context, id string) (*model.Repository, error)
}

type IssueService interface {
	GetIssueByID(ctx context.Context, id string) (*model.Issue, error)
	GetIssueByRepoAndNumber(ctx context.Context, repoID string, number int) (*model.Issue, error)
	ListIssueInRepository(ctx context.Context, repoID string, after *string, before *string, first *int, last *int) (*model.IssueConnection, error)
}

type ProjectService interface {
	GetProjectByID(ctx context.Context, id string) (*model.ProjectV2, error)
}

type OwnerService interface {
	GetUserByID(ctx context.Context, id string) (*model.User, error)
}

type PullRequestService interface {
	GetPullRequestByID(ctx context.Context, id string) (*model.PullRequest, error)
}

func New(exec boil.ContextExecutor) Services {
	return &services{
		userService:        &userService{exec: exec},
		repositoryService:  &repositoryService{exec: exec},
		issueService:       &issueService{exec: exec},
		ownerService:       &ownerService{exec: exec},
		projectV2Service:   &projectV2Service{exec: exec},
		pullRequestService: &pullRequestService{exec: exec},
	}
}
