package ui

import (
	"github.com/charmbracelet/lipgloss"
)

var (
	errorColor = lipgloss.AdaptiveColor{
		Light: "#e94560",
		Dark:  "#f05945",
	}

	itemColor         = lipgloss.NewStyle().Bold(true).Foreground(lipgloss.AdaptiveColor{Light: "#9EFFFF", Dark: "#2AFFDF"})
	itemSelectedColor = lipgloss.NewStyle().Bold(true).Foreground(lipgloss.Color("#9EFFFF"))

	listStyle    = lipgloss.NewStyle().Margin(2)
	detailsStyle = lipgloss.NewStyle().PaddingLeft(2)

	errorStyle = lipgloss.NewStyle().Foreground(errorColor)
)

const (
	iconSelected    = "●"
	iconNotSelected = "○"
	iconSynced      = "✓"
	iconSyncFailed  = "⨯"
	separator       = " • "
)
