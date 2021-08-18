package ci

import (
	"os"

	"github.com/akitasoftware/akita-libs/github"
	"github.com/akitasoftware/akita-libs/tags"
)

// https://circleci.com/docs/2.0/env-vars/#built-in-environment-variables
func circleCIInfo() (*github.PullRequest, map[tags.Key]string) {
	pr, err := github.ParsePullRequestURL(os.Getenv("CIRCLE_PULL_REQUEST"))
	if err == nil {
		pr.Branch = os.Getenv("CIRCLE_BRANCH")
		pr.Commit = os.Getenv("CIRCLE_SHA1")
	} else {
		debugln("Unable to determine GitHub PR from environment. Is `CIRCLE_PULL_REQUEST` set correctly?")
	}

	tags := map[tags.Key]string{
		tags.XAkitaCI:               CircleCI.String(),
		tags.XAkitaGitRepoURL:       os.Getenv("CIRCLE_REPOSITORY_URL"),
		tags.XAkitaGitBranch:        os.Getenv("CIRCLE_BRANCH"),
		tags.XAkitaGitCommit:        os.Getenv("CIRCLE_SHA1"),
		tags.XAkitaGitHubPRURL:      os.Getenv("CIRCLE_PULL_REQUEST"),
		tags.XAkitaCircleCIBuildURL: os.Getenv("CIRCLE_BUILD_URL"),
	}
	debugDumpEnv([]string{
		"CIRCLE_CI_REPOSITORY_URL",
		"CIRCLE_BRANCH",
		"CIRCLE_SHA1",
		"CIRCLE_PULL_REQUEST",
		"CIRCLE_BUILD_URL",
	})
	return pr, tags
}
