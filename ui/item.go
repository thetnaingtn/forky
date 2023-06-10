package ui

import (
	"fmt"

	"github.com/charmbracelet/bubbles/list"
	"github.com/thetnaingtn/forky"
)

type item struct {
	repo     *forky.RepositoryWithDetails
	selected bool
	synced   bool
	errMsg   string
}

func (i item) Title() string {
	var fork string
	if i.repo.Parent != "" {
		fork = fmt.Sprintf(" (fork from %s)", i.repo.Parent)
	}

	if !i.synced && i.errMsg != "" {
		return iconSyncFailed + " " + i.repo.FullName + " " + "is faied to sync with" + " " + i.repo.Parent
	}

	if i.synced {
		return iconSynced + " " + i.repo.FullName + " " + "is up to date with" + " " + i.repo.Parent
	}

	if i.selected {
		return iconSelected + " " + i.repo.FullName + fork
	}

	return iconNotSelected + " " + i.repo.FullName + fork
}

func (i item) Description() string {
	var details []string
	repo := i.repo
	if !i.synced {
		details = append(details, fmt.Sprintf("%d commit%s behind", repo.BehindBy, mayBePlural(repo.BehindBy)))
	}

	if !i.synced && i.errMsg != "" {
		details = append([]string{}, i.errMsg)
	}

	return detailsStyle.Render(details...)
}

func (i item) FilterValue() string {
	return "  " + i.repo.FullName
}

func splitBySelection(items []list.Item) ([]*forky.RepositoryWithDetails, []*forky.RepositoryWithDetails) {
	var selected, unselected []*forky.RepositoryWithDetails

	for _, it := range items {
		item := it.(item)
		if item.selected {
			selected = append(selected, item.repo)
		} else {
			unselected = append(unselected, item.repo)
		}
	}

	return selected, unselected
}

func reposToItems(repos []*forky.RepositoryWithDetails) []list.Item {
	items := make([]list.Item, 0, len(repos))
	for _, repo := range repos {
		items = append(items, item{repo: repo})
	}

	return items
}

func mayBePlural(behindBy int) string {
	if behindBy > 1 {
		return "s"
	}

	return ""
}
