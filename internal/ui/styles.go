package ui

import "charm.land/lipgloss/v2"

var selectedStyle = lipgloss.NewStyle().
	Background(lipgloss.Color("#333333")).
	Foreground(lipgloss.Color("#FFFFFF")).
	Bold(true)

var createStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("#78C850"))

var updateStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("#F0C75E"))

var deleteStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("#EF6A67"))

var replaceStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("#AA7CFF"))
