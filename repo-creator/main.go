package main

import (
	"flag"
	"fmt"
	"os"
	"repo-creator/template"

	"github.com/benchkram/errz"
)

func main() {

	projects := flag.Int("projects", 42, "number of projects created")
	dir := flag.String("dir", "", "directory to create projects")
	flag.Parse()

	if *dir == "" {
		fmt.Println("failed to provide --dir")
		os.Exit(1)
	}

	fmt.Println("Creating monorepo")

	err := os.MkdirAll(*dir, 0755)
	errz.Fatal(err)

	template.CreateEnvironment(*dir, *projects)
	for i := 0; i < *projects; i++ {
		err = template.CreateGoProject(*dir, i, 1)
		errz.Fatal(err)
	}

}
