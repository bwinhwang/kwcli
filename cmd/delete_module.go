package cmd

import (
	"github.com/spf13/cobra"
)

// ... (Your other imports)

var deleteModuleCmd = &cobra.Command{
	Use:   "delete_module",
	Short: "Delete a module.",
	Run: func(cmd *cobra.Command, args []string) {
		actionOrientedCommand(cmd, "delete_module")
	},
}

func init() {
	// Required flags
	deleteModuleCmd.Flags().StringP("project", "p", "", "project name")
	deleteModuleCmd.MarkFlagRequired("project")
	deleteModuleCmd.Flags().StringP("name", "n", "", "module name")
	deleteModuleCmd.MarkFlagRequired("name")

	// Record flags in definedFlags
	definedFlags["project"] = true
	definedFlags["name"] = true

	rootCmd.AddCommand(deleteModuleCmd)
}
