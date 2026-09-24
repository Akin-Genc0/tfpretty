package ui

import (
	"fmt"

	tea "charm.land/bubbletea/v2"
)

func renderHelpScreen(m Model) tea.View {
	content := fmt.Sprintf(
		"%s\n\n%s\n%s\n\n%s\n%s\n\n%s\n%s",
		titleStyle.Render("tfpretty Help"),
		sectionStyle.Render("NAVIGATION"),
		"  ↑/↓     move between resources\n  enter   view selected resource details\n  r       inspect the raw Terraform resource\n  esc     return to the plan\n  ?       show this help screen\n  q       quit",
		sectionStyle.Render("SCREEN SUMMARY"),
		"  plan screen   shows all resource changes and action counts\n  detail screen shows the selected resource fields and values\n  raw screen    shows the full Terraform resource JSON for inspection",
		sectionStyle.Render("GENERAL"),
		"  ?        show this help screen\n  esc      go back to the previous screen\n  q        quit the app",
	)

	return tea.NewView(content)
}
