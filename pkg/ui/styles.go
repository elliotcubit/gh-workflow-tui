package ui

import "github.com/charmbracelet/lipgloss"

var (
	appStyle = lipgloss.NewStyle().Padding(1, 2)

	titleStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color("#cad3f5")).
			Background(lipgloss.Color("#494d64")).
			Padding(0, 1)

	statusMessageStyle = lipgloss.NewStyle().
				Foreground(lipgloss.Color("#a6da95")).
				Render

	// The normal item state.
	itemStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color("#cad3f5"))

	// The selected item state.
	selectedItemStyle = lipgloss.NewStyle().
				Foreground(lipgloss.Color("#7dc4e4"))

	// The dimmed state, for when the filter input is initially activated.
	dimmedItemStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color("#8087a2")).
			Strikethrough(true)
)
