package ui

import "github.com/charmbracelet/bubbles/key"

var (
	keySelectToggle       = key.NewBinding(key.WithKeys(" "), key.WithHelp("space", "toggle selected item"))
	keyMergedWithUpStream = key.NewBinding(key.WithKeys("m"), key.WithHelp("m", "merge selected item with upstream branch"))
	keyRefresh            = key.NewBinding(key.WithKeys("r"), key.WithHelp("r", "refresh the forks list"))
)
