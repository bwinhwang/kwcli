package cmd

import (
	"github.com/spf13/cobra"
)

var complianceDeleteCmd = &cobra.Command{
	Use:   "compliance_delete",
	Short: "Delete compliance report file or folder.",
	Run: func(cmd *cobra.Command, args []string) {
		actionOrientedCommand(cmd, "compliance_delete")
	},
}

func init() {
	// Required flags
	complianceDeleteCmd.Flags().StringP("project", "p", "", "Project name (required)")
	complianceDeleteCmd.MarkFlagRequired("project")
	complianceDeleteCmd.Flags().StringP("path", "", "", "Relative path to the report file or folder (returned by the compliance report generate method) (required)")
	complianceDeleteCmd.MarkFlagRequired("path")

	// Optional flags
	complianceDeleteCmd.Flags().StringP("force", "f", "false", "Set to true if you want to delete a compliance report folder and all of its contents (default: false)")

	// Record flags in definedFlags
	definedFlags["project"] = true
	definedFlags["path"] = true
	definedFlags["force"] = true

	rootCmd.AddCommand(complianceDeleteCmd)
}
