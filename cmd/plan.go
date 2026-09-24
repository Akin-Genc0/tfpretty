package cmd

import (
	"fmt"
	"time"

	tea "charm.land/bubbletea/v2"
	"github.com/Akin-Genc0/tfpretty/internal/terraform"
	"github.com/Akin-Genc0/tfpretty/internal/ui"
	"github.com/spf13/cobra"
)

// spin shows a simple loading animation until stop is closed.
func spin(stop <-chan struct{}) {
	frames := []string{"/", "-", "\\", "|"}
	for i := 0; ; i++ {
		select {
		case <-stop:
			fmt.Print("\r  \r")
			return
		default:
			fmt.Printf("\r%s running terraform plan...", frames[i%len(frames)])
			time.Sleep(120 * time.Millisecond)
		}
	}
}

var planCmd = &cobra.Command{
	Use:   "plan",
	Short: "Display a formatted Terraform plan",
	Long:  "Runs Terraform plan and displays a clearer, formatted result.",
	RunE: func(cmd *cobra.Command, args []string) error {
		stop := make(chan struct{})
		go spin(stop)
		output, err := terraform.RunPlan()
		close(stop)
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
