package cmd

import (
	"github.com/spf13/cobra"
)

// ... (Your other imports)

var updateCiBuildCmd = &cobra.Command{
	Use:   "update_ci_build",
	Short: "Update a CI build.",
	Run: func(cmd *cobra.Command, args []string) {
		actionOrientedCommand(cmd, "update_ci_build")
	},
}

func init() {
	// Required flags
	updateCiBuildCmd.Flags().StringP("project", "p", "", "project or stream name")
	updateCiBuildCmd.MarkFlagRequired("project")
	updateCiBuildCmd.Flags().StringP("name", "n", "", "CI build name")
	updateCiBuildCmd.MarkFlagRequired("name")

	// Optional flags
	updateCiBuildCmd.Flags().StringP("new_name", "", "", "new CI build name")
	updateCiBuildCmd.Flags().StringP("tags", "", "", "list of comma separated tags...")

	// Record flags in definedFlags
	definedFlags["project"] = true
	definedFlags["name"] = true
	definedFlags["new_name"] = true
	definedFlags["tags"] = true

	rootCmd.AddCommand(updateCiBuildCmd)
}
