package cmd

import (
	"github.com/spf13/cobra"
)

// ... (Your other imports)

var updateViewCmd = &cobra.Command{
	Use:   "update_view",
	Short: "Update a view.",
	Run: func(cmd *cobra.Command, args []string) {
		actionOrientedCommand(cmd, "update_view")
	},
}

func init() {
	// Required flags
	updateViewCmd.Flags().StringP("project", "p", "", "project name")
	updateViewCmd.MarkFlagRequired("project")
	updateViewCmd.Flags().StringP("name", "n", "", "current view name")
	updateViewCmd.MarkFlagRequired("name")

	// Optional flags
	updateViewCmd.Flags().StringP("new_name", "", "", "new view name")
	updateViewCmd.Flags().StringP("query", "", "", "updates the search query")
	updateViewCmd.Flags().StringP("tags", "", "", "list of ...")
	updateViewCmd.Flags().BoolP("is_public", "", false, "whether the view is visible...")

	// Record flags in definedFlags
	definedFlags["project"] = true
	definedFlags["name"] = true
	definedFlags["new_name"] = true
	definedFlags["query"] = true
	definedFlags["tags"] = true
	definedFlags["is_public"] = true

	rootCmd.AddCommand(updateViewCmd)
}
