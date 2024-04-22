package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// Define a struct for issue details (adjust as needed)
type IssueDetails struct {
	ID           string `json:"id"`
	Code         string `json:"code"`
	Name         string `json:"name"`
	Location     string `json:"location"`
	Build        string `json:"build"`
	Severity     string `json:"severity"`
	SupportLevel string `json:"supportLevel"`
	Owner        string `json:"owner"`
	State        string `json:"state"`
	Status       string `json:"status"`
}

// (Potentially) define a struct for XSync information
type XSync struct {
	// Include fields based on your API's xsyncInfo
}

var issueDetailsCmd = &cobra.Command{
	Use:   "issue_details",
	Short: "Get details for the given issue id.",
	Run: func(cmd *cobra.Command, args []string) {
		var issueDetails IssueDetails
		var issueDetailss []interface{}
		err := fetchDataCommand2(cmd, "issue_details", &issueDetailss, &issueDetails)
		if err != nil {
			fmt.Println(err)
			return
		}

		if len(issueDetailss) == 1 {
			err = writeJSONToFile(issueDetailss[0], outputFile)
			if err != nil {
				fmt.Println(err)
			}
		} else {
			fmt.Println("no such issue")
			return
		}

	},
}

func init() {
	// Required flags
	issueDetailsCmd.Flags().StringP("project", "p", "", "Name of the project or stream")
	issueDetailsCmd.MarkFlagRequired("project")
	issueDetailsCmd.Flags().IntP("id", "i", 0, "The ID to search")
	issueDetailsCmd.MarkFlagRequired("id")

	// Optional flags
	issueDetailsCmd.Flags().BoolP("include_xsync", "", false, "Include xSyncInfo")

	// Record flags in definedFlags
	definedFlags["project"] = true
	definedFlags["id"] = true
	definedFlags["include_xsync"] = true

	rootCmd.AddCommand(issueDetailsCmd)
}
