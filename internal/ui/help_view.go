package ui

import tea "charm.land/bubbletea/v2"

func renderHelpScreen(m Model) tea.View {
	content := `tfpretty Help

NAVIGATION
	up/down      move between resources
	enter        view selected resource details
	  d            inspect the selected resource changes

GENERAL
	?            show this help screen
	esc          return to the plan
	q, ctrl+c    quit
`

	return tea.NewView(content)
}
