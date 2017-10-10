load("@io_bazel_rules_go//go:def.bzl", "gazelle")

gazelle(
    name = "gazelle",
    prefix = ".",
    args = [
        "-build_file_name=BUILD",
    ],
    external = "vendored",
)
