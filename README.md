Simple example that attempts to use gazelle with multiple `src/`
repositories.

How to check that the go flow is fine:

    $ export GOPATH=$PWD
    $ go install repo.4/cmd/cmd4
    $ ./bin/cmd4
    pkg1
    pkg2
    pkg3
    pkg4
    $

The broken bazel flow:

    $ cd src
    $ bazel run //:gazelle

The issue is some of the `deps` from `repo.4/cmd/cmd4/BUILD`
incorrectly point to `//vendor/`.

The resulting `go_default_library` target:

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

The expected `go_default_library` target:

    go_library(
        name = "go_default_library",
        srcs = ["main.go"],
        importpath = "repo.4/cmd/cmd4",
        visibility = ["//visibility:private"],
        deps = [
            "//vendor/repo.1/pkg1:go_default_library",
            "//vendor/repo.2/pkg2:go_default_library",
            "//repo.3/pkg3:go_default_library",
            "//repo.4/pkg4:go_default_library",
        ],
    )
