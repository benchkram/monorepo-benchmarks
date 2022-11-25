package template

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"

	"github.com/benchkram/errz"
)

func projectName(iteration int) string {
	return fmt.Sprintf("go-p%d", iteration)
}

func maingo(iteration int) []byte {
	return []byte(fmt.Sprintf(`package main

import (
	"fmt"

	"monorepo-benchmarks/apps/go-p%d/pkg/pk%d"
)

func main() {
	fmt.Println("Hello go-p%d")

	pk%d.Print()
}
 
`, iteration, iteration, iteration, iteration))
}

func pkgname(iteration int) string {
	return fmt.Sprintf("pk%d", iteration)
}
func pkgfile(iteration int) string {
	return fmt.Sprintf("pk%d.go", iteration)
}
func pkgx(interation int) []byte {
	return []byte(fmt.Sprintf(`package pk%d

import "fmt"

func Print() {
	fmt.Println("Helloo pk%d")
}

`, interation, interation))
}

func CreateGoProject(basedir string, iteration int) error {
	wd, _ := os.Getwd()
	err := os.Chdir(basedir)
	errz.Fatal(err)
	defer os.Chdir(wd)

	pn := projectName(iteration)

	err = os.Mkdir(pn, 0755)
	errz.Fatal(err)

	// main.go
	err = ioutil.WriteFile(filepath.Join(pn, "main.go"), maingo(iteration), 0644)
	errz.Fatal(err)

	// pkg/pkx/pkx.go
	pkgiteration := 1
	pkg := filepath.Join(pn, "pkg", pkgname(pkgiteration))
	err = os.MkdirAll(pkg, 0755)
	errz.Fatal(err)
	err = ioutil.WriteFile(filepath.Join(pkg, pkgfile(pkgiteration)), pkgx(pkgiteration), 0644)
	errz.Fatal(err)

	return nil
}
