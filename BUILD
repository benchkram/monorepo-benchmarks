load("@bazel_gazelle//:def.bzl", "gazelle")

# gazelle:exclude repo-creator
# gazelle:prefix monorepo-benchmarks
gazelle(name = "gazelle")

gazelle(
    name = "gazelle-update-repos",
    args = [
        "-from_file=apps/go.mod",
        "-to_macro=deps.bzl%go_dependencies",
        "-prune",
        "-build_file_proto_mode=disable_global",
    ],
    command = "update-repos",
)
