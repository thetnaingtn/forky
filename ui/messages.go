package ui

import (
	"github.com/thetnaingtn/forky"
)

type getReposListMsg struct{}
type gotReposListMsg struct {
	repos []*forky.RepositoryWithDetails
}

type mergeSelectedRepos struct{}
type mergedSelectedRepos struct{}
