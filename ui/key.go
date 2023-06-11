package ui

import "github.com/charmbracelet/bubbles/key"

var (
	keySelectToggle      = key.NewBinding(key.WithKeys(" "), key.WithHelp("space", "toggle repository"))
	keyMergeWithUpstream = key.NewBinding(key.WithKeys("m"), key.WithHelp("m", "merge with upstream repository's branch"))
	keyRefresh           = key.NewBinding(key.WithKeys("r"), key.WithHelp("r", "refresh the list"))
)
