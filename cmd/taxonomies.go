package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

type Taxonomy struct {
	Name     string `json:"name"`
	IsCustom bool   `json:"is_custom"`
}

var taxonomiesCmd = &cobra.Command{
	Use:   "taxonomies",
	Short: "Retrieve the list of taxonomy terms for a project.",
	Run: func(cmd *cobra.Command, args []string) {
		var i Taxonomy
		results, err := fetchDataCommand(cmd, "taxonomies", &i)

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
	taxonomiesCmd.Flags().StringP("project", "p", "", "project name")
	taxonomiesCmd.MarkFlagRequired("project")

	// Record the flag in definedFlags
	definedFlags["project"] = true

	// Assuming you have a 'rootCmd' defined in your main
	rootCmd.AddCommand(taxonomiesCmd)
}
