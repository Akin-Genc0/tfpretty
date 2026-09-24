package ui

import (
	"encoding/json"
	"fmt"

	tea "charm.land/bubbletea/v2"
	"github.com/Akin-Genc0/tfpretty/internal/terraform"
)

func rawResource(resource terraform.ResourceChange) string {
	b, err := json.MarshalIndent(resource, "", "  ")
	if err != nil {
		return fmt.Sprintf("%#v", resource)
	}
	return string(b)
}

func renderDiffScreen(m Model) tea.View {
	resource := m.Plan.ResourceChanges[m.Cursor]

	content := fmt.Sprintf(
		"Resource: %s\n%s\n%s\n\n%s%s",
		resource.Name,
		"Raw Terraform Resource",
		"──────────────────────────────────────────────────────────────",
		rawResource(resource),
		renderFooter("esc back       ? help       q quit"),
	)

	return tea.NewView(content)
}
