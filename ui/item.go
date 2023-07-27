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

	if repo.ParentFullName != "" {
		fork = fmt.Sprintf(" (fork from %s)", repo.ParentFullName)
	}

	titleStr := repo.FullName + fork

	if i.synced {
		return iconSynced + " " + titleStr
	}

	if !i.synced && i.repo.Error != nil {
		return errorStyle.Render(iconSyncFailed + " " + titleStr)
	}

	if i.selected {
		return iconSelected + " " + titleStr
	}

	return iconNotSelected + " " + titleStr
}

func (i item) Description() string {
	repo := i.repo
	upstream := fmt.Sprintf("%s:%s", repo.ParentFullName, repo.DefaultBranch)
	base := fmt.Sprintf("%s:%s", repo.FullName, repo.DefaultBranch)

	if i.repo.Error != nil {
		reason := i.repo.Error.Error()
		msg := base + " " + "fail to sync with" + " " + upstream + fmt.Sprintf("(%s)", reason)
		return errorStyle.Copy().PaddingLeft(2).Render(msg)
	}

	if !i.synced {
		upstream = fmt.Sprintf("%s:%s", repo.Parent, repo.DefaultBranch)
		msg := fmt.Sprintf("%s is %d commit%s behind %s", base, repo.BehindBy, mayBePlural(repo.BehindBy), upstream)
		return detailsStyle.Render(msg)
	}

	if i.synced {
		msg := base + " " + "is up to date with" + " " + upstream
		return detailsStyle.Render(msg)
	}

	return ""
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
