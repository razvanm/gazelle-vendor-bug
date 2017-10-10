load("@io_bazel_rules_go//go:def.bzl", "go_binary", "go_library")

go_library(
    name = "go_default_library",
    srcs = ["main.go"],
    importpath = "repo.4/cmd/cmd4",
    visibility = ["//visibility:private"],
    deps = [
        "//vendor/repo.1/pkg1:go_default_library",
        "//vendor/repo.2/pkg2:go_default_library",
        "//vendor/repo.3/pkg3:go_default_library",
        "//vendor/repo.4/pkg4:go_default_library",
    ],
)

go_binary(
    name = "cmd4",
    importpath = "repo.4/cmd/cmd4",
    library = ":go_default_library",
    visibility = ["//visibility:public"],
)
