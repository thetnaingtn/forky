package ui

import "github.com/charmbracelet/lipgloss"

var (
	errorColor = lipgloss.AdaptiveColor{
		Light: "#e94560",
		Dark:  "#f05945",
	}

	listStyle      = lipgloss.NewStyle().Margin(2)
	listTitleStyle = lipgloss.NewStyle().Bold(true).Foreground(lipgloss.Color("#E36CEE"))

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
