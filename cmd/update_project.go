package cmd

import (
	"github.com/spf13/cobra"
)

// ... (Your other imports)

var updateProjectCmd = &cobra.Command{
	Use:   "update_project",
	Short: "Update a project.",
	Run: func(cmd *cobra.Command, args []string) {
		actionOrientedCommand(cmd, "update_project")
	},
}

func init() {
	// Required flags
	updateProjectCmd.Flags().StringP("name", "n", "", "project or stream name")
	updateProjectCmd.MarkFlagRequired("name")

	// Optional flags
	updateProjectCmd.Flags().StringP("new_name", "", "", "new name")
	updateProjectCmd.Flags().StringP("description", "", "", "new description")
	updateProjectCmd.Flags().StringP("tags", "", "", "list of comma separated tags...")
	updateProjectCmd.Flags().BoolP("auto_delete_builds", "", false, "whether the builds in the project should automatically be deleted (true|false) (projects only)")
	updateProjectCmd.Flags().IntP("auto_delete_threshold", "", 20, "the number of builds to keep in the project if auto_delete_builds is true (default: 20) (projects only)")

	// Record flags in definedFlags
	definedFlags["name"] = true
	definedFlags["new_name"] = true
	definedFlags["description"] = true
	definedFlags["tags"] = true
	definedFlags["auto_delete_builds"] = true
	definedFlags["auto_delete_threshold"] = true
	// ... record other flags similarly

	rootCmd.AddCommand(updateProjectCmd)
}
