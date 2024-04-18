package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// todo
// Define a struct for individual metrics (adjust as needed)
type Metric struct {
	Name string `json:"name"`
	// ... Other fields based on your Klocwork server's response ...
}

// (Potentially) Define a struct for aggregated statistics
type MetricStatistics struct {
	Min     float64 `json:"min"` // Adjust the type if needed
	Max     float64 `json:"max"`
	Total   float64 `json:"total"`
	Entries int     `json:"entries"`
}

var metricsCmd = &cobra.Command{
	Use:   "metrics",
	Short: "Retrieve the list of metrics.",
	Run: func(cmd *cobra.Command, args []string) {
		var i Metric
		results, err := fetchDataCommand(cmd, "metrics", &i)
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
	metricsCmd.Flags().StringP("project", "p", "", "Project or stream name...")
	metricsCmd.MarkFlagRequired("project")

	// Optional flags

	metricsCmd.Flags().StringP("query", "q", "", "Search query...")
	metricsCmd.Flags().StringP("view", "", "", "View name")
	metricsCmd.Flags().IntP("limit", "", 1000, "Search result limit")
	metricsCmd.Flags().BoolP("aggregate", "", false, "Retrieve aggregated statistics")
	metricsCmd.Flags().BoolP("exclude_system_files", "", false, "Omit system files")

	// Record flags in definedFlags
	definedFlags["project"] = true
	definedFlags["query"] = true
	definedFlags["view"] = true
	definedFlags["limit"] = true
	definedFlags["aggregate"] = true
	definedFlags["exclude_system_files"] = true

	rootCmd.AddCommand(metricsCmd)
}
