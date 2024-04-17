package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// Project struct (adjust field types if needed)
type Project struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	Creator     string `json:"creator"`
	Description string `json:"description"`
}

var projectsCmd = &cobra.Command{
	Use:   "projects",
	Short: "Retrieve list of projects (and optionally streams)",
	Run: func(cmd *cobra.Command, args []string) {
		// Create a map to store the parameters
		var i Project
		results, err := fetchDataCommand(cmd, "projects", &i)

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
	projectsCmd.Flags().BoolP("include_streams", "s", false, "Retrieve streams as well")
	definedFlags["include_streams"] = true

	rootCmd.AddCommand(projectsCmd)
}
