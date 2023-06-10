package ui

import "github.com/charmbracelet/bubbles/list"

func newDelegateList() list.DefaultDelegate {
	defaultItemStyles := list.NewDefaultItemStyles()

	defaultItemStyles.NormalTitle = itemColor
	defaultItemStyles.NormalDesc = itemColor

	defaultItemStyles.SelectedTitle = itemSelectedColor
	defaultItemStyles.SelectedDesc = itemSelectedColor

	list := list.NewDefaultDelegate()
	list.Styles = defaultItemStyles

	return list

}
