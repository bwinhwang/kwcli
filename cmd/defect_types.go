package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// Define a struct for defect types (adjust fields as needed)
type DefectType struct {
	Code         string `json:"code"`
	Name         string `json:"name"`
	Enabled      bool   `json:"enabled"`
	Severity     int    `json:"severity"`
	SupportLevel int    `json:"supportLevel"`
}

var defectTypesCmd = &cobra.Command{
	Use:   "defect_types",
	Short: "Retrieve the list of defect types.",
	Run: func(cmd *cobra.Command, args []string) {
		var defectTypes []interface{}
		var defectType DefectType
		err := fetchDataCommand2(cmd, "defect_types", &defectTypes, &defectType)
		if err != nil {
			fmt.Println(err)
			return
		}

		err = writeJSONToFile(defectTypes, outputFile)
		if err != nil {
			fmt.Println(err)
		}

	},
}

func init() {
	// Required flags
	defectTypesCmd.Flags().StringP("project", "p", "", "Project name")
	defectTypesCmd.MarkFlagRequired("project")

	// Optional flags
	defectTypesCmd.Flags().StringP("taxonomy", "t", "", "Filter by taxonomy(\"C and C++\", \"Java\", \"MISRA C++ 2008 certified\", \"MISRA C 2023 (C11) all\")")

	// Record flags in definedFlags
	definedFlags["project"] = true
	definedFlags["taxonomy"] = true

	rootCmd.AddCommand(defectTypesCmd)
}
