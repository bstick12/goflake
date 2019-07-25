load("@io_bazel_rules_go//go:def.bzl", "go_library", "go_test")
load("@bazel_gazelle//:def.bzl", "gazelle")

# gazelle:prefix github.com/bstick12/goflake
gazelle(name = "gazelle")

go_library(
    name = "go_default_library",
    srcs = ["goflake.go"],
    importpath = "github.com/bstick12/goflake",
    visibility = ["//visibility:public"],
)

go_test(
    name = "go_default_test",
    srcs = [
        "goflake_suite_test.go",
        "goflake_test.go",
    ],
    embed = [":go_default_library"],
    deps = [
        "@com_github_onsi_ginkgo//:go_default_library",
        "@com_github_onsi_ginkgo//reporters:go_default_library",
        "@com_github_onsi_gomega//:go_default_library",
    ],
)
