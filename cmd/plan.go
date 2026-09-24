package cmd

import (
	tea "charm.land/bubbletea/v2"
	"github.com/Akin-Genc0/tfpretty/internal/terraform"
	"github.com/Akin-Genc0/tfpretty/internal/ui"
	"github.com/spf13/cobra"
)

var planCmd = &cobra.Command{
	Use:   "plan",
	Short: "Display a formatted Terraform plan",
	Long:  "Runs Terraform plan and displays a clearer, formatted result.",
	RunE: func(cmd *cobra.Command, args []string) error {
		output, err := terraform.RunPlan()
		if err != nil {
			return err
		}

		plan, err := terraform.ParsePlan(output)
		if err != nil {
			return err
		}

		program := tea.NewProgram(ui.Model{
			Plan:   plan,
			Screen: ui.PlanScreen,
			Cursor: 0,
		})

		_, err = program.Run()
		return err
	},
}

func init() {
	rootCmd.AddCommand(planCmd)
}
