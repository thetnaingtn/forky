package ui

import (
	"context"
	"log"

	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
)

func enqueuegetReposListCmd() tea.Msg {
	return getReposListMsg{}
}

func refreshReposListCmd() tea.Msg {
	return refreshReposListMsg{}
}

func requestMergeReposCmd() tea.Msg {
	return mergeSelectedReposMsg{}
}

func (app *AppModel) getReposCmd(ctx context.Context) tea.Cmd {
	return func() tea.Msg {
		repos, err := app.synrk.GetForks(ctx)
		if err != nil {
			log.Println("getReposCmd: ", err)
			return errorMsg{error: err}
		}

		return gotReposListMsg{repos: repos}
	}
}

func (app *AppModel) mergeReposCmd(repos []list.Item) tea.Cmd {
	return func() tea.Msg {
		items := make([]list.Item, 0, len(repos))
		log.Println("mergeReposCmd")
		for _, repo := range repos {
			item := repo.(item)

			if item.selected {
				if err := app.synrk.SyncBranchWithUpstreamRepo(item.repo); err != nil {
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
