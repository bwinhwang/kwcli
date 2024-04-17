package cmd

import (
	"github.com/spf13/cobra"
)

// ... (Your other imports)

var deleteBuildCmd = &cobra.Command{
	Use:   "delete_build",
	Short: "Delete a build.",
	Run: func(cmd *cobra.Command, args []string) {
		actionOrientedCommand(cmd, "delete_build")
	},
}

func init() {
	// Required flags
	deleteBuildCmd.Flags().StringP("project", "p", "", "project or stream name")
	deleteBuildCmd.MarkFlagRequired("project")
	deleteBuildCmd.Flags().StringP("name", "n", "", "build name")
	deleteBuildCmd.MarkFlagRequired("name")

	// Record flags in definedFlags
	definedFlags["project"] = true
	definedFlags["name"] = true

	rootCmd.AddCommand(deleteBuildCmd)
}
