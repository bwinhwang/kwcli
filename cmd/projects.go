package cmd

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"

	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
)

// Project struct (adjust field types if needed)
type Project struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	Creator     string `json:"creator"`
	Description string `json:"description"`
}

var action = "projects"
var projectsCmd = &cobra.Command{
	Use:   action,
	Short: "Retrieve list of projects (and optionally streams)",
	Run: func(cmd *cobra.Command, args []string) {
		// Create a map to store the parameters
		paramMap := make(map[string]interface{})
		cmd.Flags().VisitAll(func(flag *pflag.Flag) {
			if definedFlags[flag.Name] { // Filter here!
				paramMap[flag.Name] = flag.Value.String()
			}
		})
		paramMap["action"] = action

		client := getKWClientInstance()
		lines, err := client.Echo(paramMap)

		if err != nil {
			fmt.Println("Error parsing project:", err)
			return
		}

		var projects []Project
		for _, line := range lines {
			var project Project
			if len(strings.TrimSpace(line)) == 0 {
				continue // Skip to the next line if is empty
			}
			err := json.Unmarshal([]byte(line), &project)
			if err != nil {
				// ... handle error ...
				fmt.Println("Error parsing project:", err)
				fmt.Println(line)
				continue // Skip to the next line if there's a parsing error
			}
			projects = append(projects, project)
		}

		// Marshall the entire data as an array of objects
		jsonData, err := json.Marshal(projects)
		if err != nil {
			// ... handle error ...
			fmt.Println("Error marshalling JSON:", err)
		}

		// Write to file
		err = os.WriteFile("projects.json", jsonData, 0644)
		if err != nil {
			// ... handle error ...
			fmt.Println("Error writing file:", err)
		}

	},
}

func init() {
	projectsCmd.Flags().BoolP("include_streams", "s", false, "Retrieve streams as well")
	definedFlags["include_streams"] = true

	rootCmd.AddCommand(projectsCmd)
}
