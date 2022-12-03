package main

import (
	"flag"
	"fmt"
	"os"
	"os/exec"
	"strings"
	"time"

	"github.com/benchkram/errz"
)

func main() {

	iterations := flag.Int("iterations", 1, "number of iterations for this benchmark")
	cmd := flag.String("cmd", "", "cmd to execute")
	verbose := flag.Bool("v", false, "be chatty")
	before := flag.String("before", "", "cmd to before each iteration")
	flag.Parse()

	if *cmd == "" {
		fmt.Println("cmd not set")
		os.Exit(1)
	}

	beforeF := func() {}
	if *before != "" {
		ps := strings.Split(*before, " ")
		beforeF = func() {
			cmd := exec.Command(ps[0], ps[1:]...)
			if *verbose {
				cmd.Stdout = os.Stdout
				cmd.Stderr = os.Stderr
			}

			err := cmd.Run()
			errz.Fatal(err)
		}
	}

	parts := strings.Split(*cmd, " ")

	fmt.Printf("Benchmarking with %d iteration\n", *iterations)
	measurements := []time.Duration{}
	for i := 1; i < *iterations+1; i++ {
		beforeF()

		cmd := exec.Command(parts[0], parts[1:]...)
		if *verbose {
			cmd.Stdout = os.Stdout
			cmd.Stderr = os.Stderr
		}

		started := time.Now()

		err := cmd.Run()
		errz.Fatal(err)

		d := time.Since(started)
		fmt.Printf("Iteration %d/%d iteration took %s\n", i, *iterations, d.String())
		measurements = append(measurements, d)
	}

	var sum time.Duration
	for _, measurement := range measurements {
		sum = sum + measurement
	}

	mean := sum / time.Duration(len(measurements))

	fmt.Printf("mean: %s\n", mean.String())
}
