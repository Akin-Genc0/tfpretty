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

var titleStyle = lipgloss.NewStyle().
	Foreground(lipgloss.Color("#E6E6E6")).
	Bold(true)

var mutedStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("#7C8A96"))
var sectionStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("#A7D8FF")).Bold(true)
var fieldNameStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("#B9FBC0")).Bold(true)
var valueStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("#F7F7F7"))
var footerStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("#8AB4F8")).Bold(true)
