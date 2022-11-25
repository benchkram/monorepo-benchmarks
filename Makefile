MAKEFLAGS := --no-print-directory --always-make

bazel-from-scratch:
	@bazel clean --expunge
	@bazel run --experimental_convenience_symlinks=ignore //:gazelle

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
