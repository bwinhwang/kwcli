package cmd

import (
	"binhong/kwcli/common"
	"encoding/json"
	"fmt"
	"strings"

	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
)

type SummaryLine struct {
	Summary SearchSummary `json:"summary"`
}

type SearchSummary struct {
	Query    string   `json:"query"`
	Project  string   `json:"project"`
	View     string   `json:"view"`
	Limit    int      `json:"limit"`
	Total    int      `json:"total"`
	Warnings []string `json:"warnings"` // Or a more descriptive type if warnings have structure
}

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

		paramMap := make(map[string]interface{})
		cmd.Flags().VisitAll(func(flag *pflag.Flag) {
			if definedFlags[flag.Name] {
				paramMap[flag.Name] = flag.Value.String()
			}
		})
		paramMap["action"] = "search"

		client := getKWClientInstance()
		lines, err := client.Execute(paramMap)
		if err != nil {
			fmt.Printf("error parsing JSON: %s ", err)
			return
		}

		var results []interface{}
		var lastline string
		isSummary, _ := cmd.Flags().GetBool("summary")
		//fmt.Println(lines[len(lines)-2])
		// Iterate through results and unmarshal
		for _, line := range lines {
			if len(strings.TrimSpace(line)) == 0 {
				continue
			}
			var result Issue
			err := json.Unmarshal([]byte(line), &result)
			if err != nil {
				fmt.Printf("error parsing JSON: %s - line: %s", err, line)
			}
			if isSummary {
				lastline = line
			}
			// Append the unmarshalled data (You'll need to adjust how the data is stored)
			results = append(results, result)
		}
		if isSummary {
			var summaryline SummaryLine
			// remove the last one
			results = results[:len(results)-1]
			//fmt.Println(lastline)
			err := json.Unmarshal([]byte(lastline), &summaryline)
			if err == nil {
				summary := summaryline.Summary
				writeJSONToFile(summary, "search_summary.json")
				writeJSONToFile(summary, "-")
			}

		}

		writeJSONToFile(results, outputFile)

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
