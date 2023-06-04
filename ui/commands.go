package ui

import (
	"context"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/google/go-github/v52/github"
	"github.com/thetnaingtn/forky"
)

func enqueuegetReposListCmd() tea.Msg {
	return getReposListMsg{}
}

func getReposCmd(client *github.Client) tea.Cmd {
	return func() tea.Msg {
		repos, _ := forky.GetForks(context.Background(), client)
		return gotReposListMsg{repos: repos}
	}
}
