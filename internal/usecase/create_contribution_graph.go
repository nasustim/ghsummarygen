package usecase

import (
	"context"
	"time"

	"github.com/nasustim/ghsummarygen/internal/infrastructure/github"
	"github.com/nasustim/ghsummarygen/internal/infrastructure/graph"
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

	startedAt, err := gc.GetYearAccountStarted(ctx, username)
	if err != nil {
		return err
	}

	r, err := gc.GetContributions(ctx, username, startedAt, time.Now().Year())
	if err != nil {
		return err
	}

	err = graph.NewGraphClient().RenderContributionByYears(r, output)
	if err != nil {
		return err
	}
	return nil
}
