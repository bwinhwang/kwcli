package cmd

import (
	"github.com/spf13/cobra"
)

// ... (Your other imports)

var deleteCiBuildCmd = &cobra.Command{
	Use:   "delete_ci_build",
	Short: "Delete a CI build.",
	Run: func(cmd *cobra.Command, args []string) {
		actionOrientedCommand(cmd, "delete_ci_build")
	},
}

func init() {
	// Required flags
	deleteCiBuildCmd.Flags().StringP("project", "p", "", "project or stream name")
	deleteCiBuildCmd.MarkFlagRequired("project")
	deleteCiBuildCmd.Flags().StringP("name", "n", "", "CI build name")
	deleteCiBuildCmd.MarkFlagRequired("name")

	// Record flags in definedFlags
	definedFlags["project"] = true
	definedFlags["name"] = true

	rootCmd.AddCommand(deleteCiBuildCmd)
}
