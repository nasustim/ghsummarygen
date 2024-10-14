package github

import (
	"context"
	"time"

	"github.com/nasustim/ghsummarygen/internal/domain/model"
	"github.com/nasustim/ghsummarygen/internal/domain/repository"
	"github.com/pkg/errors"
	"github.com/shurcooL/githubv4"
	"golang.org/x/oauth2"
)

type gitHubClient struct {
	accessToken string
}

func NewGitHubClient(accessToken string) repository.GitHubClient {
	return &gitHubClient{
		accessToken: accessToken,
	}
}

const DATETIME_LAYOUT_ISO8601 string = "2006-01-02T15:04:05+09:00"

func (gc *gitHubClient) auth(ctx context.Context) *githubv4.Client {
	client := githubv4.NewClient(
		oauth2.NewClient(
			ctx,
			oauth2.StaticTokenSource(
				&oauth2.Token{AccessToken: gc.accessToken},
			),
		),
	)
	return client
}

func (gc *gitHubClient) GetYearAccountStarted(ctx context.Context, username string) (int, error) {
	client := gc.auth(ctx)

	var q struct {
		User struct {
			ContributionsCollection struct {
				ContributionYears []githubv4.Int
			}
		} `graphql:"user(login: $userName)"`
	}
	err := client.Query(
		ctx,
		&q,
		map[string]interface{}{
			"userName": githubv4.String(username),
		},
	)
	if err != nil {
		return 0, err
	}
	if len(q.User.ContributionsCollection.ContributionYears) == 0 {
		return 0, errors.New("not found")
	}

	startYear := q.User.ContributionsCollection.ContributionYears[0]
	for _, y := range q.User.ContributionsCollection.ContributionYears {
		if startYear > y {
			startYear = y
		}
	}
	return int(startYear), nil
}

func (gc *gitHubClient) GetContributions(ctx context.Context, userName string, startYear int, endYear int) ([]model.Contribution, error) {
	if startYear > endYear {
		return nil, errors.New("invalid arguments")
	}

	client := gc.auth(ctx)

	var q struct {
		User struct {
			ContributionsCollection struct {
				TotalCommitContributions            githubv4.Int
				TotalIssueContributions             githubv4.Int
				TotalPullRequestContributions       githubv4.Int
				TotalPullRequestReviewContributions githubv4.Int
			} `graphql:"contributionsCollection(from: $yearFrom, to: $yearTo)"`
		} `graphql:"user(login: $userName)"`
	}

	yearLength := endYear - startYear + 1
	r := make([]model.Contribution, 0, yearLength)
	for year := startYear; year <= endYear; year++ {
		v := map[string]interface{}{
			"userName": githubv4.String(userName),
			"yearFrom": githubv4.DateTime{time.Date(year, time.Month(1), 1, 0, 0, 0, 0, time.Local)},
			"yearTo":   githubv4.DateTime{time.Date(year, time.Month(12), 31, 23, 59, 59, 59, time.Local)},
		}
		err := client.Query(ctx, &q, v)
		if err != nil {
			return nil, errors.Wrap(err, "failed to Query")
		}

		r = append(r, model.Contribution{
			Year:                                year,
			TotalCommitContributions:            int(q.User.ContributionsCollection.TotalCommitContributions),
			TotalIssueContributions:             int(q.User.ContributionsCollection.TotalIssueContributions),
			TotalPullRequestContributions:       int(q.User.ContributionsCollection.TotalPullRequestContributions),
			TotalPullRequestReviewContributions: int(q.User.ContributionsCollection.TotalPullRequestReviewContributions),
		})
	}

	return r, nil
}
