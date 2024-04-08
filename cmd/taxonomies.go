package cmd

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"

	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
)

type Taxonomy struct {
	Name     string `json:"name"`
	IsCustom bool   `json:"is_custom"`
}

var taxonomiesCmd = &cobra.Command{
	Use:   "taxonomies",
	Short: "Retrieve the list of taxonomy terms for a project.",
	Run: func(cmd *cobra.Command, args []string) {
		// Create a map to store the parameters
		paramMap := make(map[string]interface{})
		cmd.Flags().VisitAll(func(flag *pflag.Flag) {
			if definedFlags[flag.Name] { // Filter here!
				paramMap[flag.Name] = flag.Value.String()
			}
		})
		paramMap["action"] = "taxonomies"

		client := getKWClientInstance()
		lines, err := client.Execute(paramMap)
		//fmt.Println(lines)
		if err != nil {
			fmt.Println(err)
			return
		}

		var taxonomies []Taxonomy
		for _, line := range lines {
			var taxonomy Taxonomy
			if len(strings.TrimSpace(line)) == 0 {
				continue // Skip to the next line if is empty
			}
			err := json.Unmarshal([]byte(line), &taxonomy)
			if err != nil {
				// ... handle error ...
				fmt.Println("Error parsing taxonomy:", err)
				fmt.Println(line)
				continue // Skip to the next line if there's a parsing error
			}
			taxonomies = append(taxonomies, taxonomy)
		}

		// Marshall the entire data as an array of objects
		jsonData, err := json.Marshal(taxonomies)
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
	taxonomiesCmd.Flags().StringP("project", "p", "", "project name")
	taxonomiesCmd.MarkFlagRequired("project")

	// Record the flag in definedFlags
	definedFlags["project"] = true

	// Assuming you have a 'rootCmd' defined in your main
	rootCmd.AddCommand(taxonomiesCmd)
}
