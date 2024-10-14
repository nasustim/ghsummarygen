package repository

import (
	"context"

	"github.com/nasustim/ghsummarygen/internal/domain/model"
)

type GitHubClient interface {
	GetYearAccountStarted(ctx context.Context, username string) (int, error)
	GetContributions(ctx context.Context, userName string, startYear int, endYear int) ([]model.Contribution, error)
}
