package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// ... (Your other imports)
type View struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	Query    string `json:"query"`
	Creator  string `json:"creator"`
	IsPublic bool   `json:"is_public"`
}

var viewsCmd = &cobra.Command{
	Use:   "views",
	Short: "Retrieve list of views.",
	Run: func(cmd *cobra.Command, args []string) {
		// Create a map to store the parameters
		var i View
		results, err := fetchDataCommand(cmd, "views", &i)

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
	viewsCmd.Flags().StringP("project", "p", "", "project name")
	viewsCmd.MarkFlagRequired("project")

	// Record the flag in definedFlags
	definedFlags["project"] = true

	// Assuming you have a 'rootCmd' defined in your main
	rootCmd.AddCommand(viewsCmd)
}
