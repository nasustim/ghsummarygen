package graph

import (
	"bytes"
	"math"
	"os"

	"github.com/nasustim/ghsummarygen/internal/domain/model"
	"github.com/nasustim/ghsummarygen/internal/domain/repository"
	"gonum.org/v1/plot"
	"gonum.org/v1/plot/plotter"
	"gonum.org/v1/plot/plotutil"
	"gonum.org/v1/plot/vg"
	"gonum.org/v1/plot/vg/draw"
	"gonum.org/v1/plot/vg/vgsvg"
)

type graph struct{}

func NewGraphClient() repository.GraphClient { return &graph{} }

func (g *graph) RenderContributionGraphEachYears(data []model.Contribution, outputFile string) error {
	p := plot.New()
	p.X.Label.Text = "Year"
	p.Y.Label.Text = "Contributions"

	commits := make(plotter.XYs, 0, len(data))
	issues := make(plotter.XYs, 0, len(data))
	prs := make(plotter.XYs, 0, len(data))
	reviews := make(plotter.XYs, 0, len(data))
	for _, v := range data {
		commits = append(commits, plotter.XY{
			X: float64(v.Year),
			Y: float64(v.CommitCount),
		})
		issues = append(issues, plotter.XY{
			X: float64(v.Year),
			Y: float64(v.IssueCount),
		})
		prs = append(prs, plotter.XY{
			X: float64(v.Year),
			Y: float64(v.PRCount),
		})
		reviews = append(reviews, plotter.XY{
			X: float64(v.Year),
			Y: float64(v.ReviewCount),
		})
	}
	err := plotutil.AddLinePoints(p,
		"commits", commits,
		"issues", issues,
		"PRs", prs,
		"reviews", reviews,
	)
	if err != nil {
		return err
	}

	size := 20 * vg.Centimeter
	canvas := vgsvg.New(size, size/vg.Length(math.Phi))
	p.Draw(draw.New(canvas))
	out := new(bytes.Buffer)
	_, err = canvas.WriteTo(out)
	if err != nil {
		return err
	}

	if err := os.WriteFile(outputFile, out.Bytes(), 0644); err != nil {
		return err
	}

	return nil
}
