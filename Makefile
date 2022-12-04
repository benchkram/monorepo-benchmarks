MAKEFLAGS := --no-print-directory --always-make

bazel-from-scratch:
	@bazel clean --expunge
	@bazel run --experimental_convenience_symlinks=ignore //:gazelle


gazelle:
	bazel run --experimental_convenience_symlinks=ignore //:gazelle

gazelle-fix:
	bazel run --experimental_convenience_symlinks=ignore //:gazelle -- fix

gazelle-update:
	bazel run --experimental_convenience_symlinks=ignore //:gazelle -- update

gazelle-update-repos:
	bazel run --experimental_convenience_symlinks=ignore  //:gazelle -- update-repos
	# Quoting the gazelle docs:
	# "After running update-repos, you might want to run bazel run //:gazelle again, as the update-repos command can affect the output of a normal run of Gazelle."
	make gazelle-update

bazel-build:
	time bazel build --experimental_convenience_symlinks=ignore  //apps/go-p1:go-p1



build-repo-creator:
	go build -o rc ./repo-creator
generate-repo: build-repo-creator clean
	./rc --dir=apps --projects=100
clean:
	rm -rf apps/*


build-benchmark:
	go build -o benchmark ./benchmark
	mv ./benchmark/benchmark bm

### ------- benchmark system preparation ----------

# At the end do a `time bob build` manually
bob-prepare-initial: generate-repo
	nix-collect-garbage -d
	bob clean system
	go clean --modcache

# At the end do a `bazel build --experimental_convenience_symlinks=ignore //...`
bazel-prepare-initial: generate-repo
	bazel clean --expunge
	bazel run --experimental_convenience_symlinks=ignore //:gazelle
	bazel clean --expunge

# At the end do a `time bob build` manually
bob-prepare-incremental: generate-repo
	bob build

# At the end do a `bazel build --experimental_convenience_symlinks=ignore //...` manually
bazel-prepare-incremental: generate-repo
	bazel run --experimental_convenience_symlinks=ignore //:gazelle
	bazel build --experimental_convenience_symlinks=ignore //...

# benchmarks
prepare: build-benchmark build-repo-creator clean
projects10: 
	./rc --dir=apps --projects=10
projects50: 
	./rc --dir=apps --projects=50
projects100:
	./rc --dir=apps --projects=100
projects1000:
	./rc --dir=apps --projects=1000
projects5000:
	./rc --dir=apps --projects=5000

benchmark-bob-10: prepare projects10
	bob build
	./bm -iterations=200 -cmd "bob build"
benchmark-bob-50: prepare projects50
	bob build
	./bm -iterations=200 -cmd "bob build"
benchmark-bob-100: prepare projects100
	bob build
	./bm -iterations=200 -cmd "bob build"
benchmark-bob-1000: prepare projects1000
	bob build
	./bm -iterations=200 -cmd "bob build"
benchmark-bob-5000: prepare projects5000
	bob build
	./bm -iterations=200 -v -cmd "bob build"
	
benchmark-bazel-10: prepare projects10
	bazel run --experimental_convenience_symlinks=ignore //:gazelle
	bazel build --experimental_convenience_symlinks=ignore //...
	./bm -iterations=200 -cmd "bazel build --experimental_convenience_symlinks=ignore //..."
benchmark-bazel-50: prepare projects50
	bazel run --experimental_convenience_symlinks=ignore //:gazelle
	bazel build --experimental_convenience_symlinks=ignore //...
	./bm -iterations=200 -cmd "bazel build --experimental_convenience_symlinks=ignore //..."
benchmark-bazel-100: prepare projects100
	bazel run --experimental_convenience_symlinks=ignore //:gazelle
	bazel build --experimental_convenience_symlinks=ignore //...
	./bm -iterations=200 -cmd "bazel build --experimental_convenience_symlinks=ignore //..."
benchmark-bazel-1000: prepare projects1000
	bazel run --experimental_convenience_symlinks=ignore //:gazelle
	bazel build --experimental_convenience_symlinks=ignore //...
	./bm -iterations=200 -cmd "bazel build --experimental_convenience_symlinks=ignore //..."
benchmark-bazel-5000: prepare projects5000
	bazel run --experimental_convenience_symlinks=ignore //:gazelle
	bazel build --experimental_convenience_symlinks=ignore //...
	./bm -iterations=200 -cmd "bazel build --experimental_convenience_symlinks=ignore //..."