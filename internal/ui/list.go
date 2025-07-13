package ui

import (
	"github.com/charmbracelet/bubbles/key"
	"github.com/charmbracelet/bubbles/list"
	"github.com/charmbracelet/bubbles/spinner"
)

func newList() list.Model {
	defaultItemStyles := list.NewDefaultItemStyles()

	defaultItemStyles.NormalTitle = itemStyle.PaddingLeft(2)
	defaultItemStyles.SelectedTitle = itemSelectedStyle.PaddingLeft(1)

	defaultItemStyles.NormalDesc = itemStyle.Copy().Faint(true)
	defaultItemStyles.SelectedDesc = itemSelectedStyle

	delegateList := list.NewDefaultDelegate()
	delegateList.Styles = defaultItemStyles

	list := list.New([]list.Item{}, delegateList, 0, 0)
	list.SetSpinner(spinner.MiniDot)

	list.Styles.Title = listTitleStyle
	list.Title = "Synrk - Synchronize your forks"

	list.AdditionalShortHelpKeys = func() []key.Binding {
		return []key.Binding{
			keySelectToggle,
			keyMergeWithUpstream,
		}
	}

	list.AdditionalFullHelpKeys = func() []key.Binding {
		return []key.Binding{keyRefresh, keySelectAll, keySelectNone}
	}

	return list
}
