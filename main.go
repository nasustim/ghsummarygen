package main

import (
	"flag"
)

type Args struct {
	d bool
}

var args Args

func init() {
	flag.BoolVar(&args.d, "d", false, "debug option")
}

func main() {
	flag.Parse()

}
