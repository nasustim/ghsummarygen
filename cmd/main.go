package main

import (
	"context"
	"errors"
	"flag"

	"github.com/nasustim/ghsummarygen/internal/usecase"
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

	err := usecase.NewCreateContributionGraph().Execute(ctx, args.accessToken, args.userName, args.outputFile)
	if err != nil {
		panic(err)
	}
}
