package main

import (
	"flag"
	"os"

	"github.com/budougumi0617/hc"
)

var (
	printHeaderFlag bool
)

func main() {
	flag.BoolVar(&printHeaderFlag, "print", false, "print result header")
	flag.BoolVar(&printHeaderFlag, "p", false, "print result header")
	flag.Parse()

	cli := &hc.Client{}
	os.Exit(cli.Execute())

}
