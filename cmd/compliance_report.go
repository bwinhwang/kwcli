package cmd

import (
	"github.com/spf13/cobra"
)

// ... (Your other imports)

var complianceReportCmd = &cobra.Command{
	Use:   "compliance_report",
	Short: "Generates compliance report of chosen type.",
	Run: func(cmd *cobra.Command, args []string) {
		actionOrientedCommand(cmd, "compliance_report")
	},
}

func init() {
	// Required flags
	complianceReportCmd.Flags().StringP("project", "p", "", "Project or stream full name (required)")
	complianceReportCmd.MarkFlagRequired("project")
	complianceReportCmd.Flags().StringP("taxonomy", "t", "", "Taxonomy name (required). Must be URL encoded")
	complianceReportCmd.MarkFlagRequired("taxonomy")

	// Optional flags
	complianceReportCmd.Flags().StringP("outputFile", "o", "", "Name of the report file to be generated")
	complianceReportCmd.Flags().StringP("fileType", "", "PDF", "Report file type. Options: PDF, DOCX (PDF by default)")
	complianceReportCmd.Flags().BoolP("summaryOnly", "", true, "Determines if report is summary only or full detailed (True by default)")
	complianceReportCmd.Flags().IntP("defectLimit", "", 100000, "Limits number of reported defects. Only works for full non-summary reports (100000 defects by default)")
	complianceReportCmd.Flags().StringP("outputFolder", "", "", "Name of the output folder where generated report file is stored (relative to projects 'compliance_reports' folder)")
	complianceReportCmd.Flags().StringP("view", "", "", "View used to generated report for (*default* is used if not specified)")
	complianceReportCmd.Flags().StringP("build", "", "", "Build id (last build used by default)")
	complianceReportCmd.Flags().StringP("reportFormat", "", "Generic", "Report format options: Generic, MISRA (Generic by default)")

	// Record flags in definedFlags
	// Record flags in definedFlags
	definedFlags["project"] = true
	definedFlags["taxonomy"] = true
	definedFlags["outputFile"] = true // Record optional flags
	definedFlags["fileType"] = true
	definedFlags["summaryOnly"] = true
	definedFlags["defectLimit"] = true
	definedFlags["outputFolder"] = true
	definedFlags["view"] = true
	definedFlags["build"] = true
	definedFlags["reportFormat"] = true

	rootCmd.AddCommand(complianceReportCmd)
}
