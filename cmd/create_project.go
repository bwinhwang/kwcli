package cmd

import (
	"github.com/spf13/cobra"
)

var createProjectCmd = &cobra.Command{
	Use:   "create_project",
	Short: "Create a project.",
	Run: func(cmd *cobra.Command, args []string) {
		actionOrientedCommand(cmd, "create_project")
	},
}

func init() {
	// Required flags
	createProjectCmd.Flags().StringP("name", "n", "", "project or stream name")
	createProjectCmd.MarkFlagRequired("name")

	// Record flags in definedFlags
	definedFlags["name"] = true

	rootCmd.AddCommand(createProjectCmd)
}
