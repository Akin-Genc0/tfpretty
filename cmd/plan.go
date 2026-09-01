package cmd

import (
	"github.com/spf13/cobra"
)

var planCmd = &cobra.Command{
	Use:   "plan",
	Short: "Display a formatted Terraform plan",
	Long:  "Runs Terraform plan and displays a clearer, formatted result.",
	RunE: func(cmd *cobra.Command, args []string) error {

		return nil
	},
}

func init() {
	rootCmd.AddCommand(planCmd)
}
