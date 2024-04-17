package cmd

import (
	"github.com/spf13/cobra"
)

// ... (Your other imports)

var deleteProjectCmd = &cobra.Command{
	Use:   "delete_project",
	Short: "Delete a project.",
	Run: func(cmd *cobra.Command, args []string) {
		actionOrientedCommand(cmd, "delete_project")
	},
}

func init() {
	// Required flags
	deleteProjectCmd.Flags().StringP("name", "n", "", "project or stream name")
	deleteProjectCmd.MarkFlagRequired("name")

	// Optional flags
	deleteProjectCmd.Flags().BoolP("force", "f", false, "set to true to force delete the project with it's streams (default : false)")

	// Record flags in definedFlags
	definedFlags["name"] = true
	definedFlags["force"] = true

	// Assuming you have a 'rootCmd' defined in your main
	rootCmd.AddCommand(deleteProjectCmd)
}
