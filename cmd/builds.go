package cmd

import (
	"binhong/kwcli/common"
	"fmt"

	"github.com/spf13/cobra"
)

type Build struct {
	ID     int           `json:"id"`
	Name   string        `json:"name"`
	Date   common.MyTime `json:"date"`
	KeepIt bool          `json:"keepit"`
	Tags   []string      `json:"tags"`
}

var buildCmd = &cobra.Command{
	Use:   "builds",
	Short: "Retrieve the list of builds for a project.",
	Run: func(cmd *cobra.Command, args []string) {

		var i Build
		results, err := fetchDataCommand(cmd, "builds", &i)

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
	buildCmd.Flags().StringP("project", "p", "", "project or stream name")
	buildCmd.MarkFlagRequired("project")

	// Record the flag in definedFlags
	definedFlags["project"] = true

	// Assuming you have a 'rootCmd' defined in your main
	rootCmd.AddCommand(buildCmd)
}
