package main

import (
	"context"
	"errors"
	"flag"

	"github.com/nasustim/ghsummarygen/pkg/github_client"
	"github.com/nasustim/ghsummarygen/pkg/graph"
)

type Args struct {
	d           bool
	userName    string
	accessToken string
	outputFile  string
}

var args Args

func init() {
	flag.BoolVar(&args.d, "d", false, "debug option")
	flag.StringVar(&args.userName, "user_name", "", "target user name")
	flag.StringVar(&args.accessToken, "access_token", "", "github pat")
	flag.StringVar(&args.outputFile, "out", "./graph.svg", "output file")
}

func main() {
	ctx := context.Background()
	flag.Parse()

	if args.accessToken == "" {
		panic(errors.New("access_token is required"))
	}
	if args.userName == "" {
		panic(errors.New("user_name is required"))
	}

	gc := github_client.NewGitHubClient(args.accessToken)
	r, err := gc.GetContributionsEachYears(ctx, args.userName)
	if err != nil {
		panic(err)
	}

	err = graph.RenderContributionGraphEachYears(r, args.outputFile)
	if err != nil {
		panic(err)
	}
}
