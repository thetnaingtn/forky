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
	repo := i.repo

	if repo.Parent != "" {
		fork = fmt.Sprintf(" (fork from %s)", repo.Parent)
	}

	titleStr := repo.FullName + fork

	if i.synced {
		return iconSynced + " " + titleStr
	}

	if !i.synced && i.errMsg != "" {
		return iconSyncFailed + " " + titleStr
	}

	if i.selected {
		return iconSelected + " " + titleStr
	}

	return iconNotSelected + " " + titleStr
}

func (i item) Description() string {
	repo := i.repo
	upstream := fmt.Sprintf("%s:%s", repo.Parent, repo.DefaultBranch)
	base := fmt.Sprintf("%s:%s", repo.FullName, repo.DefaultBranch)
	var msg string

	if i.synced {
		msg = base + " " + "is up to date with" + " " + upstream
	}

	if !i.synced && i.errMsg != "" {
		reason := i.errMsg
		msg = base + " " + "fail to sync with" + " " + upstream + fmt.Sprintf("(%s)", reason)
	}

	if !i.synced {
		msg = fmt.Sprintf("%s is %d commit%s behind %s", base, repo.BehindBy, mayBePlural(repo.BehindBy), upstream)
	}

	return detailsStyle.Render(msg)
}

func (i item) FilterValue() string {
	return "  " + i.repo.FullName
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
