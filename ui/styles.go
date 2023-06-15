package ui

import "github.com/charmbracelet/lipgloss"

// colors
var (
	orange = lipgloss.Color("#FF9D00")

	dustyGray = lipgloss.Color("#a3b2b5")
	slateTeal = lipgloss.Color("#4f6965")

	mochaClay  = lipgloss.Color("#977d6a")
	desertMist = lipgloss.Color("#bfb89d")

	mysticJade = lipgloss.Color("#86a399")
	frostyMint = lipgloss.Color("#CCEBE3")
)

var (
	listStatusStyle = lipgloss.NewStyle().Bold(true).Foreground(orange)

	listTitleStyle = lipgloss.NewStyle().Foreground(
		lipgloss.AdaptiveColor{
			Light: string(slateTeal),
			Dark:  string(dustyGray),
		},
	).Bold(true)

	itemStyle = lipgloss.NewStyle().Foreground(
		lipgloss.AdaptiveColor{
			Light: string(mochaClay),
			Dark:  string(desertMist),
		},
	)

	itemSelectedStyle = lipgloss.NewStyle().Foreground(
		lipgloss.AdaptiveColor{
			Light: string(mysticJade),
			Dark:  string(frostyMint),
		},
	).Bold(true)
)

var (
	errorColor = lipgloss.AdaptiveColor{
		Light: "#e94560",
		Dark:  "#f05945",
	}

	listStyle = lipgloss.NewStyle().Margin(2)

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
