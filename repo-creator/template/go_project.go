package template

import (
	"bytes"
	"fmt"
	"html/template"
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

	"monorepo-benchmarks/apps/go-p%d/pkg/pk0"
)

func main() {
	fmt.Println("Hello go-p%d")

	pk0.Print()
}
 
`, iteration, iteration))
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

var bobyamlT = `project: monorepo-benchmarks

nixpkgs: https://github.com/NixOS/nixpkgs/archive/nixos-22.11.tar.gz
dependencies: [
  go_1_19,
]

build:
  build:
    dependsOn:
    {{- range $val := Iterate .Projects }}
      - gop{{ $val }}
    {{- end }}
{{- range $val := Iterate .Projects }}
  gop{{ $val }}:
    input: go-p{{ $val }}/**/*.go
    cmd: cd go-p{{ $val }} && go build -o go-p{{ $val }}
    target: go-p{{ $val }}/go-p{{ $val }}
{{- end }}
`

type BobYamlTemplate struct {
	Projects uint
}

func createBobYaml(projects int) []byte {

	byT := BobYamlTemplate{Projects: uint(projects)}
	t := template.New("bobyaml")

	t = t.Funcs(template.FuncMap{
		"Iterate": func(count uint) []uint {
			var i uint
			var Items []uint
			for i = 0; i < count; i++ {
				Items = append(Items, i)
			}
			return Items
		},
	})

	t, err := t.Parse(bobyamlT)
	if err != nil {
		panic(err)
	}

	buf := bytes.NewBuffer([]byte{})
	err = t.Execute(buf, byT)
	if err != nil {
		panic(err)
	}

	return buf.Bytes()
}

func gomod() []byte {
	return []byte(`module monorepo-benchmarks

go 1.18

`)
}

func CreateEnvironment(basedir string, projects int) {
	wd, _ := os.Getwd()
	err := os.Chdir(basedir)
	errz.Fatal(err)
	defer os.Chdir(wd)

	// // go.mod
	// err = ioutil.WriteFile("go.mod", gomod(), 0644)
	// errz.Fatal(err)
	// bob.yaml
	err = ioutil.WriteFile("bob.yaml", createBobYaml(projects), 0644)
	errz.Fatal(err)
}

func CreateGoProject(basedir string, iteration int, packagesToGenerate int) error {

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
	for i := 0; i < packagesToGenerate; i++ {
		pkg := filepath.Join(pn, "pkg", pkgname(i))
		err = os.MkdirAll(pkg, 0755)
		errz.Fatal(err)
		err = ioutil.WriteFile(filepath.Join(pkg, pkgfile(i)), pkgx(i), 0644)
		errz.Fatal(err)
	}

	return nil
}
