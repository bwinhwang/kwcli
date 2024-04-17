package cmd

import (
	"github.com/spf13/cobra"
)

// ... (Your other imports)

var updateModuleCmd = &cobra.Command{
	Use:   "update_module",
	Short: "Update a module for a project.",
	Run: func(cmd *cobra.Command, args []string) {
		actionOrientedCommand(cmd, "update_module")
	},
}

func init() {
	// Required flags
	updateModuleCmd.Flags().StringP("project", "p", "", "project name")
	updateModuleCmd.MarkFlagRequired("project")
	updateModuleCmd.Flags().StringP("name", "", "", "module name")
	updateModuleCmd.MarkFlagRequired("name")

	// Other flags
	updateModuleCmd.Flags().StringP("new_name", "n", "", "new module name")
	updateModuleCmd.Flags().BoolP("allow_all", "a", false, "module access...")
	updateModuleCmd.Flags().StringP("allow_users", "", "", "grant access to users...")
	updateModuleCmd.Flags().StringP("allow_groups", "", "", "grant access to user groups...")
	updateModuleCmd.Flags().StringP("deny_users", "", "", "deny access to users...")
	updateModuleCmd.Flags().StringP("deny_groups", "", "", "deny access to user groups...")
	updateModuleCmd.Flags().StringP("paths", "", "", "list of comma separated path regexps")
	updateModuleCmd.Flags().StringP("tags", "", "", "list of comma separated tags...")

	// Record flags in definedFlags
	definedFlags["project"] = true
	definedFlags["name"] = true
	definedFlags["new_name"] = true
	definedFlags["allow_all"] = true
	definedFlags["allow_users"] = true
	definedFlags["allow_groups"] = true
	definedFlags["deny_users"] = true
	definedFlags["deny_groups"] = true
	definedFlags["paths"] = true
	definedFlags["tags"] = true
	// ... record other flags ...

	rootCmd.AddCommand(updateModuleCmd)
}
