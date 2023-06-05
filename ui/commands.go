package ui

import (
	"context"
	"log"
	"strings"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/google/go-github/v52/github"
	"github.com/thetnaingtn/forky"
)

func enqueuegetReposListCmd() tea.Msg {
	return getReposListMsg{}
}

func requestMergeReposCmd() tea.Msg {
	return mergeSelectedRepos{}
}

func getReposCmd(client *github.Client) tea.Cmd {
	return func() tea.Msg {
		repos, _ := forky.GetForks(context.Background(), client)
		return gotReposListMsg{repos: repos}
	}
}

func mergeReposCmd(client *github.Client, repos []*forky.RepositoryWithDetails) tea.Cmd {
	return func() tea.Msg {
		var names []string
		for _, repo := range repos {
			names = append(names, repo.Name)
		}
		log.Printf("mergeReposCmd: %s", strings.Join(names, ", "))

		if err := forky.SyncBranchWithUpstreamRepo(client, repos); err != nil {
		}

		return mergedSelectedRepos{}
	}
}
