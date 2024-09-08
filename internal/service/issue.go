package service

import (
	"context"
	"log"

	"github.com/nsym-m/go-graphql/internal/db"
	"github.com/nsym-m/go-graphql/internal/graph/model"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
)

type issueService struct {
	exec boil.ContextExecutor
}

func (i issueService) GetIssueByRepoAndNumber(ctx context.Context, repoID string, number int) (*model.Issue, error) {

	repo, err := db.Issues(
		qm.Select(
			db.IssueColumns.ID,
			db.IssueColumns.URL,
			db.IssueColumns.Title,
			db.IssueColumns.Closed,
			db.IssueColumns.Number,
			db.IssueColumns.Repository,
		),
		db.IssueWhere.Repository.EQ(repoID),
		db.IssueWhere.Number.EQ(int64(number)),
	).One(ctx, i.exec)
	if err != nil {
		return nil, err
	}
	return i.convertRepository(repo), nil
}

func (i issueService) convertRepository(issue *db.Issue) *model.Issue {
	issueURL, err := model.UnmarshalURI(issue.URL)
	if err != nil {
		log.Println("invalid URI", issue.URL)
	}

	return &model.Issue{
		ID:         issue.ID,
		URL:        issueURL,
		Title:      issue.Title,
		Closed:     (issue.Closed == 1),
		Number:     int(issue.Number),
		Repository: &model.Repository{ID: issue.Repository},
	}
}
