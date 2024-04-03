package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
)

var createModuleCmd = &cobra.Command{
	Use:   "create_module",
	Short: "Create a module for a project",
	Run: func(cmd *cobra.Command, args []string) {
		// Create a map to store the parameters
		paramMap := make(map[string]interface{})
		cmd.Flags().VisitAll(func(flag *pflag.Flag) {
			if definedFlags[flag.Name] { // Filter here!
				paramMap[flag.Name] = flag.Value.String()
			}
		})

		// Access and use your map
		for name, value := range paramMap {
			fmt.Printf("Parameter: %s, Value: %s\n", name, value)
		}
	},
}

func init() {
	// Required flags
	createModuleCmd.Flags().StringP("project", "p", "", "project name")
	createModuleCmd.MarkFlagRequired("project")

	createModuleCmd.Flags().StringP("name", "n", "", "module name")
	createModuleCmd.MarkFlagRequired("name")

	createModuleCmd.Flags().StringSliceP("paths", "", []string{}, "list of comma separated path regexps")
	createModuleCmd.MarkFlagRequired("paths")

	// Optional flags
	createModuleCmd.Flags().Bool("allow_all", false, "module access (‘true’ to allow access for everyone by default)")
	createModuleCmd.Flags().StringSlice("allow_users", []string{}, "grant access to users (list of comma separated user names)")
	createModuleCmd.Flags().StringSlice("allow_groups", []string{}, "grant access to user groups (list of comma separated group names)")
	createModuleCmd.Flags().StringSlice("deny_users", []string{}, "deny access to users (list of comma separated user names)")
	createModuleCmd.Flags().StringSlice("deny_groups", []string{}, "deny access to user groups (list of comma separated group names)")
	createModuleCmd.Flags().StringSlice("tags", []string{}, "list of comma separated tags (for example, 'c,security')")

	// Record ALL flags in definedFlags
	definedFlags["project"] = true
	definedFlags["name"] = true
	definedFlags["paths"] = true
	definedFlags["allow_all"] = true
	definedFlags["allow_users"] = true
	definedFlags["allow_groups"] = true
	definedFlags["deny_users"] = true
	definedFlags["deny_groups"] = true
	definedFlags["tags"] = true

	rootCmd.AddCommand(createModuleCmd)
}
