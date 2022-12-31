package client

import (
	"context"

	"github.com/shurcooL/githubv4"
	"golang.org/x/oauth2"
)

type gitHubClient struct {
	AccessToken string
}

type GitHubClient interface {
	Get(ctx context.Context) bool
}

func NewGitHubClient(accessToken string) GitHubClient {
	return &gitHubClient{
		AccessToken: accessToken,
	}
}

func (gc *gitHubClient) Get(ctx context.Context, user string) bool {
	ts := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: gc.AccessToken},
	)
	httpClient := oauth2.NewClient(ctx, ts)
	client := githubv4.NewClient(httpClient)

	var query struct {
		User struct {
			ContributionsCollection struct {
				ContributionCalendar struct {
					TotalContributions githubv4.Int
					Weeks              struct {
						ContributionDays struct {
							ContributionCount githubv4.Int
							Date              githubv4.Date
						}
					}
				}
			}
		} `graphql:"user(login: $user)"`
	}
	variables := map[string]any{
		"user": githubv4.String(user),
	}
	err := client.Query(ctx, &query, variables)

	return false
}
