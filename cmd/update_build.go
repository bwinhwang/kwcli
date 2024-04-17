package cmd

import (
	"github.com/spf13/cobra"
)

// ... (Your other imports)

var updateBuildCmd = &cobra.Command{
	Use:   "update_build",
	Short: "Update a build.",
	Run: func(cmd *cobra.Command, args []string) {
		actionOrientedCommand(cmd, "update_build")
	},
}

func init() {
	// Required flags
	updateBuildCmd.Flags().StringP("project", "p", "", "project or stream name")
	updateBuildCmd.MarkFlagRequired("project")
	updateBuildCmd.Flags().StringP("name", "n", "", "build name")
	updateBuildCmd.MarkFlagRequired("name")

	// Optional flags
	updateBuildCmd.Flags().StringP("new_name", "", "", "new build name")
	updateBuildCmd.Flags().BoolP("keepit", "", false, "...")
	updateBuildCmd.Flags().StringP("tags", "", "", "list of comma separated tags...")

	// Record flags in definedFlags
	definedFlags["project"] = true
	definedFlags["name"] = true
	definedFlags["new_name"] = true
	definedFlags["keepit"] = true
	definedFlags["tags"] = true

	rootCmd.AddCommand(updateBuildCmd)
}
