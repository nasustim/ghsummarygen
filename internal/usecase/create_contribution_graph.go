package usecase

import (
	"context"

	"github.com/nasustim/ghsummarygen/internal/repository/github"
	"github.com/nasustim/ghsummarygen/internal/repository/graph"
)

type CreateContributionGraph interface {
	Execute(ctx context.Context, accessToken string, username string, output string) error
}

type createContributionGraph struct{}

func NewCreateContributionGraph() CreateContributionGraph {
	return &createContributionGraph{}
}

func (uc *createContributionGraph) Execute(ctx context.Context, accessToken string, username string, output string) error {
	gc := github.NewGitHubClient(accessToken)
	r, err := gc.GetContributionsEachYears(ctx, username)
	if err != nil {
		panic(err)
	}

	err = graph.RenderContributionGraphEachYears(r, output)
	if err != nil {
		panic(err)
	}
	return nil
}
