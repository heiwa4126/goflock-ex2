package main

import (
	"fmt"
	"log"
	"os"

	"github.com/heiwa4126/goflock-ex2/ex5"
)

var (
	// Version = $(git tag --sort=-v:refname | grep '^v' | head -1 | sed 's/^v//')
	Version string = "9.9.9"
	// Revision = $(git rev-parse --short HEAD)
	Revision string = "zzzzzzz"
)

func help() {
	// help or version
	fmt.Printf("goflock-ex1 %s (%s)\n", Version, Revision)
	os.Exit(2)
}

func main() {
	var err error
	var cnt uint64

	if len(os.Args) < 2 {
		help()
	}

	switch os.Args[1] {
	case "init":
		cnt, err = ex5.InitCounter()
	case "inc":
		cnt, err = ex5.IncCounter10000(false)
	case "flockinc":
		cnt, err = ex5.IncCounter10000(true)
	case "files":
		cntf, lockf := ex5.Files()
		fmt.Printf("%s %s\n", cntf, lockf)
		return
	default:
		help()
	}

	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%d\n", cnt)
}
