package repository

import (
	"context"

	"github.com/nasustim/ghsummarygen/internal/domain/model"
)

type GitHubClient interface {
	GetContributionsEachYears(ctx context.Context, userName string) ([]model.Contribution, error)
}
