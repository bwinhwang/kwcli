package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
)

// ... (Your other imports)

var deleteProjectCmd = &cobra.Command{
	Use:   "delete_project",
	Short: "Delete a project.",
	Run: func(cmd *cobra.Command, args []string) {
		paramMap := make(map[string]interface{})
		cmd.Flags().VisitAll(func(flag *pflag.Flag) {
			if definedFlags[flag.Name] { // Filter here!
				paramMap[flag.Name] = flag.Value.String()
			}
		})
		paramMap["action"] = "delete_project"

		client := getKWClientInstance()
		_, err := client.Execute(paramMap)
		if err != nil {
			fmt.Println(err)
		}
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
