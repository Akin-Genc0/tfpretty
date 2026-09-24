package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:     "tfpretty",
	Short:   "A beautiful Terraform plan viewer",
	Long:    "tfpretty makes Terraform plans easier to read directly from your terminal.",
	Version: "0.1.0",
	RunE: func(cmd *cobra.Command, args []string) error {
		return cmd.Help()
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, "Oops. An error while executing tfpretty:", err)
		os.Exit(1)
	}
}
