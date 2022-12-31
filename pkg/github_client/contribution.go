package github_client

import (
	"context"
	"time"

	"github.com/pkg/errors"
	"github.com/shurcooL/githubv4"
	"golang.org/x/oauth2"
)

const DATETIME_LAYOUT_ISO8601 string = "2006-01-02T15:04:05+09:00"

type Contributions struct {
	TotalCommitContributions            int
	TotalIssueContributions             int
	TotalPullRequestContributions       int
	TotalPullRequestReviewContributions int
}

func (gc *gitHubClient) GetContributionsEachYears(ctx context.Context, userName string) (map[int]Contributions, error) {
	ts := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: gc.AccessToken},
	)
	httpClient := oauth2.NewClient(ctx, ts)
	client := githubv4.NewClient(httpClient)

	var contributedYearsQuery struct {
		User struct {
			ContributionsCollection struct {
				ContributionYears []githubv4.Int
			}
		} `graphql:"user(login: $userName)"`
	}
	err := client.Query(
		ctx,
		&contributedYearsQuery,
		map[string]interface{}{
			"userName": githubv4.String(userName),
		},
	)
	if err != nil {
		return map[int]Contributions{}, err
	}
	contributionYearsList := contributedYearsQuery.User.ContributionsCollection.ContributionYears

	// Note: contributionのない年を埋める処理を書きたい

	type ContributionsByYearQuery struct {
		User struct {
			ContributionsCollection struct {
				TotalCommitContributions            githubv4.Int
				TotalIssueContributions             githubv4.Int
				TotalPullRequestContributions       githubv4.Int
				TotalPullRequestReviewContributions githubv4.Int
			} `graphql:"contributionsCollection(from: $yearFrom, to: $yearTo)"`
		} `graphql:"user(login: $userName)"`
	}

	r := make(map[int]Contributions, len(contributionYearsList))
	for _, v := range contributionYearsList {
		year := int(v)

		var q ContributionsByYearQuery
		variables := map[string]interface{}{
			"userName": githubv4.String(userName),
			"yearFrom": githubv4.DateTime{time.Date(year, time.Month(1), 1, 0, 0, 0, 0, time.Local)},
			"yearTo":   githubv4.DateTime{time.Date(year, time.Month(12), 31, 23, 59, 59, 59, time.Local)},
		}
		err := client.Query(ctx, &q, variables)
		if err != nil {
			return nil, errors.Wrap(err, "failed to Query")
		}

		r[year] = Contributions{
			TotalCommitContributions:            int(q.User.ContributionsCollection.TotalCommitContributions),
			TotalIssueContributions:             int(q.User.ContributionsCollection.TotalIssueContributions),
			TotalPullRequestContributions:       int(q.User.ContributionsCollection.TotalPullRequestContributions),
			TotalPullRequestReviewContributions: int(q.User.ContributionsCollection.TotalPullRequestReviewContributions),
		}
	}

	return r, nil
}
