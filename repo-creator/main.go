package main

import (
	"fmt"
	"os"
	"repo-creator/template"

	"github.com/benchkram/errz"
)

func main() {
	fmt.Println("Creating monorepo")

	err := os.MkdirAll("test", 0755)
	errz.Fatal(err)

	err = template.CreateGoProject("./test", 1)
	errz.Fatal(err)
}
