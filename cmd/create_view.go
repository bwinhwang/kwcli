package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
)

var createViewCmd = &cobra.Command{
	Use:   "create_view",
	Short: "Create a view for a project.",
	Run: func(cmd *cobra.Command, args []string) {
		// 1. Get all flag values
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
	createViewCmd.Flags().StringP("project", "p", "", "project name")
	createViewCmd.MarkFlagRequired("project")

	createViewCmd.Flags().StringP("name", "n", "", "view name")
	createViewCmd.MarkFlagRequired("name")

	createViewCmd.Flags().String("query", "", "search query for the view")
	createViewCmd.MarkFlagRequired("query")

	// Optional flags
	createViewCmd.Flags().StringSlice("tags", []string{}, "list of comma separated tags (for example, 'c,security')")
	createViewCmd.Flags().Bool("visible_to_all", false, "whether the view is visible to all users with access to this project (true|false)")

	// Record flags in definedFlags
	definedFlags["project"] = true
	definedFlags["name"] = true
	definedFlags["query"] = true
	definedFlags["tags"] = true
	definedFlags["visible_to_all"] = true

	// Assuming you have a 'rootCmd' defined in your main
	rootCmd.AddCommand(createViewCmd)
}
