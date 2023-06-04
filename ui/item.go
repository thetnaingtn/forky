package ui

import (
	"fmt"

	"github.com/charmbracelet/bubbles/list"
	"github.com/thetnaingtn/forky"
)

type Item struct {
	repo     *forky.RepositoryWithDetails
	selected bool
}

func (i Item) Title() string {
	return fmt.Sprintf("%s (fork from %s)", i.repo.FullName, i.repo.Parent)
}

func (i Item) Description() string {
	var details []string
	repo := i.repo
	details = append(details, fmt.Sprintf("%d commit%s behind", repo.BehindBy, mayBePlural(repo.BehindBy)))

	return detailsStyle.Render(details...)
}

func (i Item) FilterValue() string {
	return "  " + i.repo.FullName
}

func reposToItems(repos []*forky.RepositoryWithDetails) []list.Item {
	items := make([]list.Item, 0, len(repos))
	for _, repo := range repos {
		items = append(items, Item{repo: repo})
	}

	return items
}

func mayBePlural(behindBy int) string {
	if behindBy > 1 {
		return "s"
	}

	return ""
}
