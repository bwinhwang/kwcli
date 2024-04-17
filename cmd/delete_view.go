package cmd

import (
	"github.com/spf13/cobra"
)

// ... (Your other imports)

var deleteViewCmd = &cobra.Command{
	Use:   "delete_view",
	Short: "Delete a view.",
	Run: func(cmd *cobra.Command, args []string) {
		actionOrientedCommand(cmd, "delete_view")
	},
}

func init() {
	// Required flags
	deleteViewCmd.Flags().StringP("project", "p", "", "project name")
	deleteViewCmd.MarkFlagRequired("project")
	deleteViewCmd.Flags().StringP("name", "n", "", "view name")
	deleteViewCmd.MarkFlagRequired("name")

	// Record flags in definedFlags
	definedFlags["project"] = true
	definedFlags["name"] = true

	rootCmd.AddCommand(deleteViewCmd)
}
