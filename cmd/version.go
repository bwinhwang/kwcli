package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// Define a struct for the version information (adjust if needed)
type KlocworkVersion struct {
	MajorVersion string `json:"majorVersion"`
	MinorVersion string `json:"minorVersion"`
}

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Retrieve Validate server version.",
	Run: func(cmd *cobra.Command, args []string) {
		var versions []interface{}
		var version KlocworkVersion
		err := fetchDataCommand2(cmd, "version", &versions, &version)
		if err != nil {
			fmt.Println(err)
			return
		}
		value, ok := versions[0].(KlocworkVersion)
		if ok {
			fmt.Printf("MajorVersion: %s, MinorVersion %s", value.MajorVersion, value.MinorVersion)
		} else {
			fmt.Println("Fail to get version")
		}
	},
}

func init() {
	// No flags necessary for a simple version check
	rootCmd.AddCommand(versionCmd)
}
