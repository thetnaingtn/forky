package ui

import (
	"github.com/thetnaingtn/forky"
)

type getReposListMsg struct{}
type gotReposListMsg struct {
	repos []*forky.RepositoryWithDetails
}

type mergeSelectedReposMsg struct{}
type mergedSelectedReposMsg struct{}
