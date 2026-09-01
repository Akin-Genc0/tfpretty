package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "tfpretty",
	Short: "A beautiful Terraform plan viewer",
	Long:  "tfpretty makes Terraform plans easier to read directly from your terminal.",
	Run: func(cmd *cobra.Command, args []string) {
		test()
	},
}

func test() {

	fmt.Println("testing")
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintf(os.Stderr, "Oops. An error while executing tfpretty '%s'\n", err)
		os.Exit(1)
	}
}
