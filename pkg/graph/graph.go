package graph

import (
	"bytes"
	"io/ioutil"
	"math"

	"github.com/nasustim/github-all-summary/pkg/github_client"
	"gonum.org/v1/plot"
	"gonum.org/v1/plot/plotter"
	"gonum.org/v1/plot/plotutil"
	"gonum.org/v1/plot/vg"
	"gonum.org/v1/plot/vg/draw"
	"gonum.org/v1/plot/vg/vgsvg"
)

func RenderContributionGraphEachYears(data []github_client.Contributions, outputFile string) error {
	p := plot.New()
	p.X.Label.Text = "Year"
	p.Y.Label.Text = "Contribution count"

	ptcc := make(plotter.XYs, len(data))
	ptic := make(plotter.XYs, len(data))
	ptprc := make(plotter.XYs, len(data))
	ptprrc := make(plotter.XYs, len(data))
	for i, v := range data {
		ptcc[i].X = float64(v.Year)
		ptcc[i].Y = float64(v.TotalCommitContributions)

		ptic[i].X = float64(v.Year)
		ptic[i].Y = float64(v.TotalIssueContributions)

		ptprc[i].X = float64(v.Year)
		ptprc[i].Y = float64(v.TotalPullRequestContributions)

		ptprrc[i].X = float64(v.Year)
		ptprrc[i].Y = float64(v.TotalPullRequestReviewContributions)

	}
	err := plotutil.AddLinePoints(p,
		"TotalCommitContributions", ptcc,
		"TotalIssueContributions", ptic,
		"TotalPullRequestContributions", ptprc,
		"TotalPullRequestReviewContributions", ptprrc,
	)
	if err != nil {
		panic(err)
	}

	size := 20 * vg.Centimeter
	canvas := vgsvg.New(size, size/vg.Length(math.Phi))
	p.Draw(draw.New(canvas))
	out := new(bytes.Buffer)
	_, err = canvas.WriteTo(out)
	if err != nil {
		return err
	}

	if err := ioutil.WriteFile(outputFile, out.Bytes(), 0644); err != nil {
		return err
	}

	return nil
}
