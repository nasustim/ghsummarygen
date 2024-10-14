package repository

import "github.com/nasustim/ghsummarygen/internal/domain/model"

type GraphClient interface {
	RenderContributionGraphEachYears(data []model.Contribution, outputFile string) error
}
