# Monorepo-Benchmarks

Utilities to benchmark [Bob](https://bob.build) against Bazel on a monorepo.

## Prerequisites
* Golang
* Bob
* Bazel

## Run the Benchmark
Each benchmark does 200x the same incremental build
```
make benchmark-bob-10 
make benchmark-bob-50 
make benchmark-bob-100
make benchmark-bob-1000
make benchmark-bob-5000

make benchmark-bazel-10
make benchmark-bazel-50
make benchmark-bazel-100
make benchmark-bazel-1000
make benchmark-bazel-5000
```
The number at the end determines how many projects are created and build.


## Result

#projects | Bob 0.7.2 | Bazel 6.0
----------| ----------| ----------
go-projects-10      | 0,061s | 0,126s
go-projects-50      | 0,067s | 0,123s
go-projects-100     | 0,074s | 0,15s
go-projects-1000    | 0,184s | 0,279s
go-projects-5000    | 0,863s | 1,037s


System: Linux Ubuntu 20.04 | AMD Ryzen 7 PRO 4750U | 32GB of RAM | SSD.


17. Feb 2023
