package ui

import (
	tea "charm.land/bubbletea/v2"
	"github.com/Akin-Genc0/tfpretty/internal/terraform"
)

func renderFooter(hints ...string) string {
	content := "↑↓ navigate       enter details       ? help       q quit"

	if len(hints) > 0 {
		content = hints[0]
	}

	return "\n───────────────────────────────────────────────────────────────\n" + content + "\n"
}

func renderActionSymbol(action terraform.Action) string {
	switch action {
	case terraform.ActionCreate:
		return "+"
	case terraform.ActionDelete:
		return "-"
	case terraform.ActionReplace:
		return "+/-"
	default:
		return "~"
	}
}

func (m Model) View() tea.View {
	switch m.Screen {
	case DetailScreen:
		return renderDetailScreen(m)
	case Help:
		return renderHelpScreen(m)
	case PlanScreen:
		return renderPlanScreen(m)
	case Diff:
		return renderDiffScreen(m)
	default:
		return renderPlanScreen(m)
	}
}
