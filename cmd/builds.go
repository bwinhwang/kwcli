package cmd

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
)

type MyTime struct {
	time.Time
}

type Build struct {
	ID     int      `json:"id"`
	Name   string   `json:"name"`
	Date   MyTime   `json:"date"`
	KeepIt bool     `json:"keepit"`
	Tags   []string `json:"tags"`
}

func (i *MyTime) UnmarshalJSON(data []byte) error {

	var timestamp int64
	if err := json.Unmarshal(data, &timestamp); err != nil {
		return err // Handle any initial unmarshaling errors
	}

	sec := timestamp / 1000          // Milliseconds to seconds
	nsec := (timestamp % 1000) * 1e6 // Remaining nanoseconds
	i.Time = time.Unix(sec, nsec)

	return nil
}

var buildCmd = &cobra.Command{
	Use:   "builds",
	Short: "Retrieve the list of builds for a project.",
	Run: func(cmd *cobra.Command, args []string) {
		// Create a map to store the parameters
		paramMap := make(map[string]interface{})
		cmd.Flags().VisitAll(func(flag *pflag.Flag) {
			if definedFlags[flag.Name] { // Filter here!
				paramMap[flag.Name] = flag.Value.String()
			}
		})
		paramMap["action"] = "builds"

		client := getKWClientInstance()
		lines, err := client.Execute(paramMap)

		if err != nil {
			fmt.Println("Error parsing build:", err)
			return
		}

		var builds []Build
		for _, line := range lines {
			var build Build
			if len(strings.TrimSpace(line)) == 0 {
				continue // Skip to the next line if is empty
			}
			err := json.Unmarshal([]byte(line), &build)
			if err != nil {
				// ... handle error ...
				fmt.Println("Error parsing build:", err)
				fmt.Println(line)
				continue // Skip to the next line if there's a parsing error
			}
			builds = append(builds, build)
		}

		// Marshall the entire data as an array of objects
		jsonData, err := json.Marshal(builds)
		if err != nil {
			// ... handle error ...
			fmt.Println("Error marshalling JSON:", err)
		}

		// Write to file
		err = os.WriteFile(outputFile, jsonData, 0644)
		if err != nil {
			// ... handle error ...
			fmt.Println("Error writing file:", err)
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
