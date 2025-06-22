package ui

import (
	"github.com/charmbracelet/bubbles/list"
	"github.com/thetnaingtn/synrk"
)

type errorMsg struct {
	error
}

func (e errorMsg) Error() string {
	return e.error.Error()
}

type refreshReposListMsg struct{}

type getReposListMsg struct{}
type gotReposListMsg struct {
	repos []*synrk.RepositoryWithDetails
}

type mergeSelectedReposMsg struct{}
type mergedSelectedReposMsg struct {
	items []list.Item
}
