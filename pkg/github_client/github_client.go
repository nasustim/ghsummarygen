package github_client

import "context"

type gitHubClient struct {
	AccessToken string
}

type GitHubClient interface {
	GetContributionsEachYears(ctx context.Context, userName string) ([]Contributions, error)
}

func NewGitHubClient(accessToken string) GitHubClient {
	return &gitHubClient{
		AccessToken: accessToken,
	}
}
