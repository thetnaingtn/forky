package ui

import (
	"context"
	"log"

	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/google/go-github/v52/github"
	"github.com/thetnaingtn/forky"
)

func enqueuegetReposListCmd() tea.Msg {
	return getReposListMsg{}
}

func requestMergeReposCmd() tea.Msg {
	return mergeSelectedReposMsg{}
}

func getReposCmd(client *github.Client) tea.Cmd {
	return func() tea.Msg {
		repos, _ := forky.GetForks(context.Background(), client)
		return gotReposListMsg{repos: repos}
	}
}

func mergeReposCmd(client *github.Client, repos []list.Item) tea.Cmd {
	return func() tea.Msg {
		items := make([]list.Item, 0, len(repos))
		log.Println("mergeReposCmd")
		for _, repo := range repos {
			item := repo.(item)

			if item.selected {
				if err := forky.SyncBranchWithUpstreamRepo(client, item.repo); err != nil {
					item.synced = false
					item.errMsg = err.Error()
				} else {
					item.synced = true
				}
			}

			items = append(items, item)
		}

		return mergedSelectedReposMsg{items: items}
	}
}
