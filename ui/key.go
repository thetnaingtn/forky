package ui

import "github.com/charmbracelet/bubbles/key"

var (
	keySelectAll         = key.NewBinding(key.WithKeys("a"), key.WithHelp("a", "select all repositories"))
	keySelectNone        = key.NewBinding(key.WithKeys("n"), key.WithHelp("n", "unselect all repositories"))
	keySelectToggle      = key.NewBinding(key.WithKeys(" "), key.WithHelp("space", "toggle repository"))
	keyMergeWithUpstream = key.NewBinding(key.WithKeys("m"), key.WithHelp("m", "merge with upstream repository's branch"))
	keyRefresh           = key.NewBinding(key.WithKeys("r"), key.WithHelp("r", "refresh the list"))
)
