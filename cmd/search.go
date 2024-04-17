package cmd

import (
	"binhong/kwcli/common"
	"fmt"

	"github.com/spf13/cobra"
)

type Issue struct {
	ID               int           `json:"id"`
	Status           string        `json:"status"`
	Severity         string        `json:"severity"`
	SeverityCode     int           `json:"severityCode"`
	SupportLevel     string        `json:"supportLevel"`
	SupportLevelCode int           `json:"supportLevelCode"`
	State            string        `json:"state"`
	Code             string        `json:"code"`
	Title            string        `json:"title"`
	Message          string        `json:"message"`
	File             string        `json:"file"`
	Method           string        `json:"method"`
	Owner            string        `json:"owner"`
	TaxonomyName     string        `json:"taxonomyName"`
	Date             common.MyTime `json:"dateOriginated"` // Use int64 for timestamps
	URL              string        `json:"url"`
	IssueIds         []int         `json:"issueIds"`
}

var searchCmd = &cobra.Command{
	Use:   "search",
	Short: "Retrieve the list of detected issues.",
	Run: func(cmd *cobra.Command, args []string) {

		var i Issue
		results, err := fetchDataCommand(cmd, "search", &i) // Placeholder
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
	searchCmd.Flags().StringP("project", "p", "", "project or stream name")
	searchCmd.MarkFlagRequired("project")

	// Optional flags
	searchCmd.Flags().StringP("query", "q", "", "search query...")
	searchCmd.Flags().StringP("view", "", "", "view name")
	searchCmd.Flags().IntP("limit", "", 1000, "search result limit")
	searchCmd.Flags().BoolP("summary", "", false, "include summary record to output stream")

	// Record flags in definedFlags
	definedFlags["project"] = true
	definedFlags["query"] = true
	definedFlags["view"] = true
	definedFlags["limit"] = true
	definedFlags["summary"] = true

	rootCmd.AddCommand(searchCmd)
}
