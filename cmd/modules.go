package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// ... (Your other imports)

type Module struct {
	Name     string   `json:"name"`
	AllowAll bool     `json:"allowAll"`
	Paths    []string `json:"paths"`
}

var modulesCmd = &cobra.Command{
	Use:   "modules",
	Short: "Retrieve the list of modules for a project.",
	Run: func(cmd *cobra.Command, args []string) {
		var i Module
		results, err := fetchDataCommand(cmd, "modules", &i)

		if err != nil {
			fmt.Println(err)
			return
		}

		err = writeJSONToFile(results, outputFile)
		if err != nil {
			fmt.Println(err)
		}
	},
}

func init() {
	// Required flags
	modulesCmd.Flags().StringP("project", "p", "", "project name")
	modulesCmd.MarkFlagRequired("project")

	// Record the flag in definedFlags
	definedFlags["project"] = true

	rootCmd.AddCommand(modulesCmd)
}
